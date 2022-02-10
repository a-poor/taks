package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-poor/taks/lib"
	"github.com/urfave/cli/v2"
)

const listDesc = ``

// listCmd corresponds to the CLI's list command.
// It lists the current tasks in the database.
var listCmd = &cli.Command{
	Name:        "list",
	Usage:       "Show tasks from the user's task-list.",
	Description: listDesc,
	Action: func(c *cli.Context) error {
		// Get the path to the taks DB
		home := "" // TODO: Get the home directory from a user-suppled argument.
		cfg, err := lib.GetConfigDir(home)
		if err != nil {
			return fmt.Errorf("unable to get user's home directory: %w", err)
		}

		// Does it exist? If not, exit.
		_, err = os.Stat(cfg)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Error: Unable to access the take DB at %q", cfg), 1)
		}

		// Get the taks DB connection
		db, err := lib.OpenDB(cfg)
		if err != nil {
			return cli.Exit(fmt.Sprintf("Error: Unable to open the take DB at %q", cfg), 1)
		}

		// Validate the DB
		if err = db.Validate(); err != nil {
			return cli.Exit(fmt.Sprintf("Error: Unable to validate the take DB at %q", cfg), 1)
		}

		tasks, err := db.ListTasks()
		if err != nil {
			return err
		}

		fmt.Printf("Found %d tasks:\n", len(tasks))

		for i, task := range tasks {
			b, err := json.Marshal(task)
			if err != nil {
				return err
			}
			fmt.Printf("[%d] %s\n", i, string(b))
		}

		return nil
	},
}
