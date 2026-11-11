# Project conventions

Shared by the Go projects under `~/repos/sjg`.

## Build output

- Binaries build to `bin/` (`bin/<name>`; projects with multiple binaries use
  `bin/<binary>`).
- `bin/` is git-ignored. `make clean` removes it.
- `make link` installs or refreshes the `~/bin/<name>` symlink.

## Make targets

| Target | What it does |
| --- | --- |
| `make init` | Point git at `.githooks` (run once per checkout/worktree). |
| `make build` | Run `staticcheck`, `lint-check`, and `go test -race`, then compile cleanly. |
| `make checks` | The check suite alone (`staticcheck`, `lint-check`, `test`). |
| `make staticcheck` | Run staticcheck. |
| `make lint` | Apply golangci-lint fixes in place. |
| `make lint-check` | Run golangci-lint without rewriting files (CI uses this). |
| `make fmt` | Apply goimports via `golangci-lint fmt`. |
| `make test` | Run `go test -race ./...`. |
| `make verify` | Check the built binary against git HEAD and tracked source mtimes. |
| `make modules` | `go mod tidy`. |
| `make clean` | Remove build output. |

`make build` and `make verify` refuse to run until `make init` has configured
the hooks.

## Git hooks

`.githooks/pre-commit` runs staticcheck, lint (with fixes), and tests. If lint
rewrites tracked files, the commit is aborted so the fixes can be reviewed and
staged.

## Linting and formatting

`.golangci.yml` is a golangci-lint v2 config shared by every project:

- standard linters (`errcheck`, `govet`, `ineffassign`, `staticcheck`, `unused`)
- `goimports` as the formatter

A project may layer documented exceptions on top (for example roster's extra
linters or tasks' `forbidigo` `os.*` policy).

## CI

`.github/workflows/ci.yml` runs on pushes and pull requests to `main`:

- latest stable Go
- pinned golangci-lint and staticcheck versions
- `make build`, then `make verify`

Docs-only changes are skipped, and superseded runs are cancelled.
