package cmd

import "github.com/urfave/cli/v2"

const listDesc = ``

// listCmd corresponds to the CLI's list command.
// It lists the current tasks in the database.
var listCmd = &cli.Command{
	Name:        "list",
	Usage:       "Show tasks from the user's task-list.",
	Description: listDesc,
	Action: func(c *cli.Context) error {
		return nil
	},
}
