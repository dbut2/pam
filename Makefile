.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

server:
	rm -f server
	go build -o server cmd/server/main.go