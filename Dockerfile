FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./vendor ./vendor

RUN go build -o server ./cmd/server/main.go

ARG PORT=8080

EXPOSE ${PORT}

CMD ["./server"]