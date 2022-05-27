GIT_SHA := $(shell git rev-parse HEAD)

binaries:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main-linux-amd64-$(GIT_SHA) 
