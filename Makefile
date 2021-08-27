.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

server:
	rm -f server
	go build -o server cmd/server/main.go

.PHONY: docs
docs:
	go run docs/gen.go
