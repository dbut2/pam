package main

import (
	"log"
	"os"

	"github.com/dbut2/pam/pkg/cli"
)

func main() {
	app := cli.Pam()

	markdown, err := app.ToMarkdown()
	if err != nil {
		log.Fatalf("error generating docs: %s\n", err.Error())
	}

	err = os.WriteFile("docs/docs.md", []byte(markdown), 0777)
	if err != nil {
		log.Fatalf("error generating docs: %s\n", err.Error())
	}
}
