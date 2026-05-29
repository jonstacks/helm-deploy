# Helm Deploy Action

[![CI](https://github.com/jonstacks/helm-deploy/actions/workflows/ci.yml/badge.svg)](https://github.com/jonstacks/helm-deploy/actions/workflows/ci.yml)

Rather than relying on the `helm` CLI, you can use the `helm-deploy` action to deploy a chart.

## Usage

```yaml

steps:
- uses: actions/checkout@v4
- uses: azure/setup-kubectl@4
- uses: azure/setup-helm@v4
  with:
    version: v3.8.2
- uses: jonstacks/helm-deploy@v0
  with:
    # Required
    release-name: my-release
    chart: my-chart
    # Optional
    atomic: false
    cleaup-on-fail: false
    create-namespace: false
    debug: false
    dry-run: false
    kube-context: my-context
    namespace: my-namespace
    timeout: 10m
    wait: true
    sets: >
      some.set="hello"
      another.set="world"
```

## Development

This repo includes a pinned Nix development shell for local work and coding agents.

- Run `direnv allow .` once, then re-enter the repository, or start the shell with `nix develop`.
- Validate the Nix environment with `nix flake check`.
- Run tests with `go test ./...` or `make test`.
- Rebuild the vendored Linux action binary with `make binaries`.
