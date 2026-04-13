# The Go Engine 🚀

A high-performance background job processor built with Go, Redis, and Docker.

## Stack
- **Go** + Gin — Fast HTTP API
- **Redis** — Job queue and state management
- **Docker Compose** — One-command deployment

## Architecture
POST /jobs → API (202) → Redis Queue → Worker → done
## Endpoints
- `POST /jobs` — Submit a job
- `GET /jobs/:id` — Check job status
- `GET /health` — Healthcheck

## Run
```bash
docker compose up --build
```

## How it works
The API accepts jobs instantly and returns 202. A background worker picks up jobs from the Redis queue, processes them concurrently using goroutines, and updates the status in Redis.

## Benchmark
Tested with 50 concurrent POST /jobs requests:
- **50/50 successful** (202 Accepted)
- **0 failures** under concurrent load
- All jobs queued and processed via goroutines
- Tested locally with Docker on Windows
