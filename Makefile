.PHONY: pre-commit

pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...

test:
	go test -race -v ./... ${TEST_FLAGS}
	