VERSION := 5073eb98bdc3e8a15dfec6ddbf7d65a904388b38

binaries:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main-linux-amd64-$(VERSION) 

test:
	go test -v -race ./...
