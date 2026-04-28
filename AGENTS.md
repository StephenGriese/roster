# AGENTS.md — roster

NHL roster web application backed by the public NHL API.

## Commands

```bash
go build ./...       # build all packages
go test ./...        # run all tests
```

## Package layout

- `cmd/server/` — Web server entrypoint
- `cmd/cmdline/` — CLI entrypoint
- `server/` — HTTP handlers, routes, components
- `nhle/` — NHL API client (`PlayerService`)
- `roster/` — Domain types

## HTTP client conventions

See the workspace-level `../AGENTS.md` for the full rules. Key points for this project:

- `nhle.PlayerService` stores its `*http.Client` as a struct field with a 1-minute timeout.
  Call `nhle.NewPlayerService()` **once** (in `NewServer`) and share it across all handlers.
  Do not call `nhle.NewPlayerService()` inside request handlers.
- Use `http.NewRequestWithContext` for all outbound requests (the existing `http.NewRequest`
  calls in `nhle/rest.go` should be migrated when that file is next touched).
- Base URLs are already named constants in `nhle/rest.go` — keep it that way.
