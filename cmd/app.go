package cmd

import (
	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

// NewApp creates the taks CLI application to be run.
func NewApp() *cli.App {
	return &cli.App{
		Name:    "taks",
		Usage:   "A CLI for managing your tasks.",
		Version: version,
		Authors: []*cli.Author{{
			Name:  "Austin Poor",
			Email: "code@austinpoor.com",
		}},
		Copyright:   "Copyright (c) 2022 Austin Poor",
		Description: rootDesc,
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "Initialize the taks database.",
				Description: initDesc,
				Action:      cliInit,
			},
			{
				Name:        "new",
				Usage:       "Create a new task.",
				Description: newDesc,
				Action:      cliNew,
			},
			{
				Name:        "list",
				Usage:       "Show tasks from the user's task-list.",
				Description: listDesc,
				Action:      cliList,
			},
		},
		Action: cliRoot,
	}
}
