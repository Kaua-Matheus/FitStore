# FitStore

Personal project for practice. The goal is to build a full hand-made application for gym owners to manage and sell products (supplements, gear, etc.) — backend and frontend written from scratch, following clean architecture and general good development practices.

## Stack

- **Backend:** Go, Gin, GORM, PostgreSQL
- **Frontend:** TypeScript, React

## Architecture

The backend follows **clean architecture**: dependencies point inward, business rules don't know about HTTP or the database.

```
backend/
  cmd/
    main.go                # application entrypoint, wires everything together
  internal/
    config/                 # env-based configuration, loaded once at startup
    domain/                 # entities and repository interfaces — no framework tags, no I/O
    usecase/                 # business logic, depends only on domain interfaces
    repository/
      postgres/              # domain interfaces implemented with GORM/Postgres
    transport/
      http/
        router.go            # route registration
        middleware/           # auth, CORS, etc.
        handler/               # parses requests, calls usecases, writes responses
        dto/
          request/              # request payload shapes + validation tags
          response/              # response payload shapes
          mapper/                # conversions between domain <-> request/response
    pkg/                     # shared utilities (jwt, crypt, cookies)
```

**Why the layers are split this way:**
- `domain` holds the real-world entities and the interfaces the rest of the app depends on. It has no `json` or `gorm` tags — it doesn't know it's going to be persisted or serialized.
- `repository` implements those interfaces against Postgres/GORM. Persistence tags and table mappings live here, not in `domain`.
- `usecase` contains the actual business rules (validation, orchestration) and is the only layer both `transport` and `repository` are wired through — it can be tested without a running HTTP server or database.
- `transport/http` is the only layer that knows about Gin, HTTP status codes, and JSON wire formats. Request/response DTOs live here (not in `domain`) so the API contract can evolve independently of the domain model, and so sensitive fields (like password hashes) never accidentally leak into a response.

The frontend follows a feature-based structure (`features/<feature>`) rather than a type-based one, keeping components, hooks, and API calls for a given feature together.

## How to configure

To run the application, both backend and frontend need to be running.

1. Copy `.env.example` to `.env` inside `backend/` and fill in the values (database credentials, JWT secret, etc.).

### Run backend

Inside `backend/`:

```bash
go run cmd/main.go
```

### Run frontend

Inside `frontend/`:

```bash
npm run start:dev
```

For production (using IP):

```bash
npm run start:prod
```

## Status

Actively being rebuilt around clean architecture as a learning exercise — see the [issues](../../issues) for planned work and in-progress epics.