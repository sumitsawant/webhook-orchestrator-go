package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "webhook-orchestrator-go/internal/service"

    log "github.com/sirupsen/logrus"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var data service.RequestPayload
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    response, err := service.CallThirdParty(data)
    if err != nil {
        http.Error(w, fmt.Sprintf("3rd party call failed: %v", err), http.StatusBadGateway)
        return
    }

    go service.SendCallback(data.CallbackURL, response)

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
    log.WithField("duration_ms", time.Since(start).Milliseconds()).Info("Request handled")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"ok"}`))
}
