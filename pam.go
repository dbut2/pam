package main

import (
	"log"
	"os"

	"github.com/dbut2/pam/pkg/cli"
)

func main() {
	cmd := cli.Pam()
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
