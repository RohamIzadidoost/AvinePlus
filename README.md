# Glass Cutting App

This repository contains a simple example application with a Go backend and a static frontend.

## Structure

```
backend/    Go API and business logic
frontend/   Static HTML/JS frontend
```

## Running Backend

1. Ensure you have a PostgreSQL instance running (default connection string uses `postgres:postgres@localhost:5432/glasscutting`).
2. From the `backend` directory run:

```bash
go run ./cmd/server
```

The server will start on `:8080`.

## Frontend

Open `frontend/public/index.html` in your browser or visit `http://localhost:8080` after running the backend.
