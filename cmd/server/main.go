package main

import (
	"os"

	"github.com/dbut2/pam/internal/server"
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
