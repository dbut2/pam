package main

import (
	"github.com/dbut2/pam/internal/docs"
	"github.com/dbut2/pam/pkg/cli"
	"log"
	"os"
)

func main() {
	cmd := cli.Pam()
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	docs.Gen()
}
