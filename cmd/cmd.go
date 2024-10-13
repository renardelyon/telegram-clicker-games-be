package cmd

import (
	"telegram-clicker-game-be/application"
	"telegram-clicker-game-be/config"

	"github.com/urfave/cli"
)

func runCommand(cfg *config.Config) func(*cli.Context) error {
	return func(c *cli.Context) error {
		app, err := application.Setup(cfg, c)
		if err != nil {
			return err
		}
		return app.Run(cfg)
	}
}

func Cli(cfg *config.Config) *cli.App {
	clientApp := cli.NewApp()
	clientApp.Name = cfg.Application.ServiceName
	clientApp.Version = cfg.Application.ServiceVersion
	clientApp.Action = runCommand(cfg)
	// clientApp.Flags = []cli.Flag{
	// 	cli.BoolFlag{
	// 		Name:     "migrate-up",
	// 		Required: false,
	// 	},
	// 	cli.BoolFlag{
	// 		Name:     "migrate-down",
	// 		Required: false,
	// 	},
	// }
	return clientApp
}
