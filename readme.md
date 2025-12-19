# Semicolon URL Shortener Service

![Go Version](https://img.shields.io/badge/go-1.24-00ADD8?style=flat&logo=go)
![Fiber Framework](https://img.shields.io/badge/fiber-v2-black?style=flat&logo=go)
![Redis](https://img.shields.io/badge/redis-v7-DC382D?style=flat&logo=redis)
![Docker](https://img.shields.io/badge/docker-ready-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/license-MIT-green)

High-performance, microservice-ready URL Shortener built with **Golang (Fiber)** and **Redis**.
Architected using **Clean Architecture** principles to ensure scalability, maintainability, and testability.

---

## 🚀 Key Features

- **⚡ Blazing Fast:** Powered by [Go Fiber](https://gofiber.io/) (fasthttp) and Redis in-memory storage.
- **🏗️ Clean Architecture:** Strict separation of concerns (Domain, Service, Infra, Handler).
- **🐳 Docker Ready:** Fully containerized with multi-stage builds and optimized caching.
- **📝 Structured Logging:** JSON logging with **Zerolog**, trace-id injection, and stack traces (New Relic friendly).
- **⚙️ 12-Factor App:** Configuration management via `envconfig` (Environment Variables).
- **🛡️ Robust Error Handling:** Standardized error responses and graceful shutdowns.

---

## 🏗️ System Architecture

The service follows a unidirectional data flow pattern:

`HTTP Request` -> `Handler (Fiber)` -> `Service (Business Logic)` -> `Repository (Redis Interface)` -> `Redis Storage`

### Folder Structure (Monorepo Style)

We use a modular structure to separate the core logic from the specific service implementation, making code reusable for future services.

```text
.
├── repo/                   # Shared Kernel & Infrastructure
│   ├── common/             # Loggers, Config Loaders, Utils
│   ├── domain/             # Core Business Interfaces (Contracts)
│   └── infra/              # External Implementations (Redis, etc.)
│
└── url-shortener/          # The Microservice Implementation
    ├── Dockerfile          # Multi-stage build definition
    ├── docker-compose.yml  # Local development orchestration
    ├── main.go             # Entry point (Wiring)
    └── service/            # Application Logic
        ├── config/         # Service-specific config mapping
        ├── handlers/       # HTTP Transport layer
        └── routers/        # Fiber Route definitions
```

# 🛠️ Tech StackLanguage:

- Go: 1.24
- Web Framework: Fiber v2
- Storage/Cache: Redis v7
- Configuration: KelseyHightower Envconfig + Godotenv
- Logging: Zerolog
- Containerization: Docker & Docker Compose

# 🏁 Getting Started

## Prerequisites

- Docker & Docker Compose
- Make (Optional)

## Run with Docker (Recommended)

You don't need Go installed locally. Docker will handle everything including the Redis dependency.

1. Navigate to the service directory:

```bash
cd url-shortener
```

2. Start the services:

```bash# This will build the Go app and start Redis
docker compose up -d --build
```

3. Verify it's running:

```bash
curl http://localhost:8080/health
# Output: {"service":"GoShort-Service","status":"ok"}
```

# Configuration (Environment Variables)

The application is configured via Environment Variables. In Docker Compose, these are pre-configured.

| Variable       | Default Value   | Description           |
| -------------- | --------------- | --------------------- |
| APP_NAME       | GoShort-Service | Service Identifier    |
| APP_PORT       | 8080            | Http Port             |
| REDIS_ENABLED  | true            | Enables Redis Storage |
| REDIS_HOST     | localhost       | Redis Hostname        |
| REDIS_PORT     | 6379            | Redis Port            |
| REDIS_PASSWORD | ``              | Redis Password        |

# 🔌 API Documentation

## 1. Shorten a URLGenerates a unique short code for a given long URL.

- ``POST /api/v1/shorten``
- ``Content-Type``: ``application/json``

### Request:

```json
{
  "url": "https://www.google.com/search?q=golang+clean+architecture"
}
```

### Response (201 Created):

```JSON
{
    "code": "AbC12",
    "short_url": "http://localhost:8080/AbC12",
    "original_url": "[https://www.google.com/search?q=golang+clean+architecture](https://www.google.com/search?q=golang+clean+architecture)"
}
```

## 2. Redirect

Redirects the user to the original URL.

- Endpoint: `GET /:code (e.g., http://localhost:8080/AbC12)`
- Response:
  - 301 Moved Permanently -> Redirects to destination.
  - 404 Not Found -> If the code does not exist.

# 🧪 Development

If you want to run it locally (without Docker container for the app):

### Start Redis separately:

```Bash
docker run -d -p 6379:6379 redis:alpine
```

### Create a .env file in the url-shortener directory:

```
APP_PORT=8080
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_ENABLED=true
```

### Run the app:

```bash
cd url-shortener
go run main.go
```

# 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

<p align="center">Architected with ❤️ by <strong>Semicolon Indonesia</strong><br><em>Simple, Secure, Scalable.</em></p>
