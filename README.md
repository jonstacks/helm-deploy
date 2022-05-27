# Helm Deploy Action

Rather than relying on the `helm` CLI, you can use the `helm-deploy` action to deploy a chart.

## Usage

```yaml

steps:
- uses: actions/checkout@v3
- uses: azure/setup-kubectl@v1
- uses: azure/setup-helm@v1
  with:
    version: v3.8.0
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
