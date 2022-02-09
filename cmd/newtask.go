package cmd

import "github.com/urfave/cli/v2"

const newDesc = ``

// newCmd corresponds to the CLI's new command.
// It creates a new task and adds it to the database.
var newCmd = &cli.Command{
	Name:        "new",
	Usage:       "Create a new task and add it to the user's task-list.",
	Description: newDesc,
	Action: func(c *cli.Context) error {
		return nil
	},
}
