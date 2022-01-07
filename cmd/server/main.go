package main

import (
	_ "embed"
	"os"

	"github.com/dbut2/pam/cmd/config"
	"github.com/dbut2/pam/internal/app"
	"github.com/dbut2/pam/internal/server"
)

func main() {
	c := config.Find()

	a := app.NewApp(c.App)

	port := os.Getenv("PORT")
	if port != "" {
		c.Server.Address = ":" + port
	}
	s := server.NewServer(c.Server)

	err := s.Serve(a)
	if err != nil {
		panic(err.Error())
	}
}
