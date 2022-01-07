cli:
	rm -f pam
	go build -o pam cmd/cli/main.go

server:
	rm -f server
	go build -o server cmd/server/main.go
