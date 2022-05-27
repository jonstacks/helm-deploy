VERSION := 65c7ebc607d90fd62527fe82a0659eba86061b3c

binaries:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main-linux-amd64-$(VERSION) 
