package main

import (
    "fmt"
    "net/http"

    "webhookorchestrator/internal/audit"
    "webhookorchestrator/internal/config"
    "webhookorchestrator/internal/handler"
    "webhookorchestrator/internal/logger"
    "webhookorchestrator/internal/middleware"
    "webhookorchestrator/internal/otelinit"

    log "github.com/sirupsen/logrus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    config.Load()
    logger.Init()
    audit.InitDB()
    otelinit.InitTracer("webhook-orchestrator")

    mux := http.NewServeMux()
    mux.HandleFunc("/", handler.HTTPHandler)
    mux.HandleFunc("/health", handler.HealthHandler)
    mux.Handle("/metrics", promhttp.Handler())

    wrapped := middleware.WithRequestID(mux)
    port := config.AppConfig.Server.Port

    if config.AppConfig.TLS.Enabled {
        log.Infof("🚀 HTTPS Server on :%d", port)
        log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", port),
            config.AppConfig.TLS.CertFile,
            config.AppConfig.TLS.KeyFile,
            wrapped))
    } else {
        log.Infof("🚀 HTTP Server on :%d", port)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), wrapped))
    }
}
