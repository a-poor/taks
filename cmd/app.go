package cmd

import (
	"github.com/urfave/cli/v2"
)

// NewApp creates the taks CLI application to be run.
func NewApp() *cli.App {
	return &cli.App{
		Name:  "taks",
		Usage: "A CLI for managing your tasks.",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize the taks database.",
				Action: cliInit,
			},
		},
		Action: cliRoot,
	}
}
