package service

import (
    "bytes"
    "io"
    "math"
    "net/http"
    "strings"
    "time"

    "webhook-orchestrator-go/internal/audit"
    "webhook-orchestrator-go/internal/config"

    log "github.com/sirupsen/logrus"
)

type RequestPayload struct {
    URL         string   `json:"url"`
    Method      string   `json:"method"`
    Auth        string   `json:"auth"`
    Headers     []string `json:"headers"`
    Body        string   `json:"body"`
    CallbackURL string   `json:"callback_url"`
}

func CallThirdParty(reqData RequestPayload) ([]byte, error) {
    client := &http.Client{Timeout: time.Duration(config.AppConfig.ThirdParty.TimeoutSecs) * time.Second}

    req, err := http.NewRequest(reqData.Method, reqData.URL, bytes.NewBuffer([]byte(reqData.Body)))
    if err != nil {
        return nil, err
    }

    for _, h := range reqData.Headers {
        if parts := strings.SplitN(h, ": ", 2); len(parts) == 2 {
            req.Header.Set(parts[0], parts[1])
        }
    }
    if reqData.Auth != "" {
        if parts := strings.SplitN(reqData.Auth, ": ", 2); len(parts) == 2 {
            req.Header.Set(parts[0], parts[1])
        }
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}

func SendCallback(url string, payload []byte) {
    retries := config.AppConfig.Callback.MaxRetries
    baseDelay := time.Duration(config.AppConfig.Callback.BaseDelaySecs) * time.Second

    for i := 0; i < retries; i++ {
        client := &http.Client{Timeout: 5 * time.Second}
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
        req.Header.Set("Content-Type", "application/json")

        resp, err := client.Do(req)
        if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
            log.WithField("url", url).Info("Callback successful")
            resp.Body.Close()
            audit.InsertCallback(url, resp.StatusCode)
            return
        }
        if resp != nil {
            resp.Body.Close()
        }
        log.WithField("attempt", i+1).Warn("Callback failed, retrying...")
        time.Sleep(time.Duration(math.Pow(2, float64(i))) * baseDelay)
    }
    log.WithField("url", url).Error("Callback failed after retries")
}
