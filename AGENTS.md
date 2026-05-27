# AGENTS.md

## Project overview

This repository ships a GitHub Action that wraps `helm upgrade --install`.

- `action.yml` defines the published action interface.
- `invoke-binary.js` is a small JavaScript shim that selects the vendored binary for the current platform.
- `main.go` and `config.go` contain the action logic.
- `main-linux-amd64-<version>` is a committed release artifact built from this repository.

## Setup commands

- Enter the reproducible dev shell with `direnv allow .` and re-enter the directory, or run `nix develop`.
- Run unit tests with `go test ./...` or `make test`.
- Rebuild the vendored Linux binary with `make binaries`.
- Validate the Nix configuration with `nix flake check`.

## Code and release conventions

- Keep `invoke-binary.js` as a thin launcher; put action behavior in Go.
- When updating the vendored binary, keep the version string aligned between `Makefile` and `invoke-binary.js`.
- Preserve pinned SHAs for third-party GitHub Actions unless there is a deliberate upgrade.

## Testing instructions

- Run `go test ./...` after Go or JavaScript shim changes.
- The integration path lives in `.github/workflows/ci.yml` and exercises the action against `kind`, `kubectl`, and `helm`.
- If you change Nix or workflow files, run `nix flake check` and `direnv exec . go test ./...`.

## Security notes

- Treat committed binaries as release artifacts, not hand-edited files.
- Prefer rebuilding vendored binaries from the pinned Nix environment and verifying the result in CI before merging PRs that change them.
