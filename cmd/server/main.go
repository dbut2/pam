package main

import (
	"github.com/dbut2/go-pam/internal/server"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := server.Run(":"+port)
	if err != nil {
		panic(err.Error())
	}
}
