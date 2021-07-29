package cli

import (
	"github.com/urfave/cli"
)

func Pam() *cli.App {
	app := &cli.App{
		Name: "pam",
		Commands: []cli.Command{
			config(),
		},
	}

	return app
}

func config() cli.Command {
	cmd := cli.Command{
		Name:                   "config",
		Subcommands: []cli.Command{
			{
				Name: "add",
				Subcommands: []cli.Command{
					{
						Name: "note",
					},
				},
			},
		},
	}

	return cmd
}
