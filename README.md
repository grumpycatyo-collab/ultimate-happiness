# ultimate-happiness

A Go backend service skeleton focused on production-style API structure, HTTP middleware, JWT authentication, configuration, logging, debug endpoints, and graceful shutdown.

This repository is not meant to be presented as a finished product. It is a compact backend architecture project that shows how I structure a Go service, separate application/bootstrap concerns from business and foundation packages, and wire operational basics around an HTTP API.

## Purpose

The goal of this project is to demonstrate practical Go backend patterns in a small codebase:

- service startup and shutdown flow
- environment-based configuration with `ardanlabs/conf`
- structured logging with `zap`
- API and debug HTTP servers
- custom middleware chain for logging, errors, metrics, panic recovery, authentication, and authorization
- JWT authentication backed by RSA keys
- package separation between `app`, `business`, and `foundation`
- small tooling commands under `app/tooling`
- tests around authentication behavior

The project is intentionally smaller than a full production system. Its value is in the service layout, operational wiring, and backend structure rather than in business-domain complexity.

## Repository structure

```text
app/
  services/sales-api/      # Main HTTP API service
  tooling/                 # Small helper binaries
business/
  sys/auth/                # JWT auth primitives
  web/mid/                 # Middleware
foundation/
  keystore/                # RSA key loading for JWT signing/validation
  web/                     # Minimal HTTP app wrapper
zarf/                      # Runtime/support files
```

## Main service

The main API service lives in:

```text
app/services/sales-api/main.go
```

It is responsible for loading configuration, initializing logging, loading authentication keys, starting the debug server, starting the API server, and shutting down gracefully on process signals.

The HTTP routing layer lives in:

```text
app/services/sales-api/handlers/
```

It wires the API routes, debug routes, middleware, authentication, and authorization rules.

## Tech stack

```text
Go 1.21
net/http
ardanlabs/conf
httptreemux
JWT
zap
UUID
automaxprocs
```

## Running tests

```sh
go test ./...
```

## Running locally

The service expects RSA private keys in the configured key folder. By default, it looks at:

```text
zarf/keys/
```

The active key id defaults to:

```text
54bb2165-71e1-41a6-af3e-7da4a0e1e2c1
```

For local runs, provide a PEM key with that id as the filename, or override the auth configuration through environment variables.

```sh
go run ./app/services/sales-api
```

Default ports:

```text
API:   0.0.0.0:3000
Debug: 0.0.0.0:4000
```

Useful endpoints:

```text
GET /v1/test
GET /v1/testauth
GET /debug/liveness
GET /debug/readiness
GET /debug/vars
GET /debug/pprof/
```
