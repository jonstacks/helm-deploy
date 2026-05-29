BIN_DIR := bin
LINUX_AMD64_BINARY := $(BIN_DIR)/helm-deploy-linux-amd64
LINUX_ARM64_BINARY := $(BIN_DIR)/helm-deploy-linux-arm64
LINUX_BINARIES := $(LINUX_AMD64_BINARY) $(LINUX_ARM64_BINARY)
BINARY_GOFLAGS := -trimpath -buildvcs=false
BINARY_LDFLAGS := -s -w -buildid=

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(LINUX_AMD64_BINARY): | $(BIN_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BINARY_GOFLAGS) -ldflags='$(BINARY_LDFLAGS)' -o $@ .

$(LINUX_ARM64_BINARY): | $(BIN_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(BINARY_GOFLAGS) -ldflags='$(BINARY_LDFLAGS)' -o $@ .

binaries: $(LINUX_BINARIES)

.PHONY: binaries verify test $(LINUX_BINARIES)

verify:
	@TMP_DIR=$$(mktemp -d) && \
	trap 'rm -rf "$$TMP_DIR"' EXIT && \
	for arch in amd64 arm64; do \
		CHECKED_IN="$(BIN_DIR)/helm-deploy-linux-$$arch"; \
		REBUILT="$$TMP_DIR/helm-deploy-linux-$$arch"; \
		CGO_ENABLED=0 GOOS=linux GOARCH=$$arch go build $(BINARY_GOFLAGS) -ldflags='$(BINARY_LDFLAGS)' -o "$$REBUILT" .; \
		if ! cmp -s "$$REBUILT" "$$CHECKED_IN"; then \
			echo "checked-in binary $$CHECKED_IN does not match a rebuild from source" >&2; \
			echo "run 'direnv exec . make binaries' to regenerate it" >&2; \
			echo "rebuilt sha256:" >&2; sha256sum "$$REBUILT" >&2; \
			echo "checked-in sha256:" >&2; sha256sum "$$CHECKED_IN" >&2; \
			exit 1; \
		fi; \
	done

test:
	go test -v -race ./...
