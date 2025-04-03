# Webhook Orchestrator (Go)

A production-grade webhook orchestration server built with Go.

## 🚀 Features

- Accepts dynamic JSON POST requests
- Calls 3rd-party APIs
- Sends async callback (with retry)
- SQLite-based audit logging
- Prometheus metrics
- TLS support
- OpenTelemetry tracing
- Request ID middleware

## 📦 Configuration

Edit `config.yaml`:

```yaml
server:
  port: 8080

callback:
  max_retries: 3
  base_delay_secs: 1

third_party:
  timeout_secs: 10

tls:
  enabled: false
  cert_file: cert.pem
  key_file: key.pem
