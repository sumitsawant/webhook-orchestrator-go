package main

import (
    "fmt"
    "net/http"

    "webhook-orchestrator-go/internal/audit"
    "webhook-orchestrator-go/internal/config"
    "webhook-orchestrator-go/internal/handler"
    "webhook-orchestrator-go/internal/logger"
    "webhook-orchestrator-go/internal/middleware"
    "webhook-orchestrator-go/internal/otelinit"

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
