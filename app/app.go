package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Start() *cli.App {
	app := &cli.App{
		Name:  "Health Checker",
		Usage: "A tool for checking the domain is down or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if ctx.String("port") == "" {
				port = "80"
			}
			status := check(ctx.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	return app
}
