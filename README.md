# Pack Calculator

A simple REST API that calculates the optimal combination of packs to fulfill an order. Given a set of available pack sizes and a number of items to ship, it finds the combination that:

1. Sends the minimum extra items (can't break packs open!)
2. Uses the fewest number of packs

Uses dynamic programming (coin-change style) under the hood.

## Getting Started

### Run locally

```bash
go run cmd/main.go
```

Server starts at `http://localhost:8080`

### Run with Docker

```bash
docker-compose up --build
```

### Run tests

```bash
go test ./...
```

### Run tests in Docker

```bash
docker-compose run --rm test
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/calculate` | Calculate optimal packs for given items |
| GET | `/packs` | Get all available pack sizes |
| POST | `/packs/{size}` | Add a new pack size |
| DELETE | `/packs/{size}` | Remove a pack size |

### Example

```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"items": 501}'
```

## Project Structure

```
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   ├── app/
│   │   └── app.go           # App bootstrap & wiring
│   ├── model/
│   │   ├── domain/          # Domain models (Pack)
│   │   └── dto/             # Request/Response DTOs
│   ├── repository/
│   │   ├── pack_repository.go
│   │   └── interfaces/      # Repository interfaces
│   ├── server/
│   │   └── handler.go       # HTTP handlers & routing
│   └── usecase/
│       ├── calculator.go    # Core algorithm (DP)
│       ├── pack_creator.go
│       ├── pack_remover.go
│       ├── pack_retriever.go
│       └── interfaces/      # Usecase interfaces
├── web/
│   └── index.html           # Simple UI
├── Dockerfile
├── docker-compose.yml
└── render.yaml              # Render.com deployment config
```

## Tech Decisions

### No external frameworks

We intentionally avoided external HTTP frameworks like **Gin**, **Echo**, **Fiber**, etc. The project uses Go's standard library `net/http` with the new Go 1.22+ routing patterns (`POST /calculate`, `DELETE /packs/{size}`). Keeps dependencies minimal and the binary small.

### No validation library

No **go-playground/validator** or similar. Validation is done manually in the DTO layer. Simple and explicit.

### In-memory storage

Packs are stored in memory. No database, no Redis, just a slice behind a mutex.

## What Could Be Improved

- **Persistent storage** - Add a real database (PostgreSQL, SQLite) so packs survive restarts
- **Validation library** - go-playground/validator would be nice for complex validation rules
- **Structured logging** - Replace `log` with something like `slog` or `zerolog`
- **Configuration** - Use env vars or a config file for port, default packs, etc.
- **OpenAPI/Swagger** - Auto-generate API docs
- **Rate limiting** - Prevent abuse on public deployments
- **Caching** - Cache calculation results for repeated requests
- **Metrics** - Add Prometheus metrics for monitoring
- **Graceful shutdown** - Handle SIGTERM properly (partially done)
- **Better error responses** - Return structured JSON errors instead of plain text
