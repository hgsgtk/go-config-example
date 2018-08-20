tools:
	go get golang.org/x/tools/cmd/goimports

fmt:
	go fmt ./...

imports:
	goimports -l -w $(shell find . -name "*.go" | grep -v /vendor/)

tests:
	go test -v -cover ./...