package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// newApp creates the taks CLI application to be run.
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
		Action: func(c *cli.Context) error {
			fmt.Println("Running the root command...")
			return nil
		},
	}
}
