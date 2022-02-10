package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/a-poor/taks/lib"
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
			initCmd,
			newCmd,
			listCmd,
			{
				Name:  "random",
				Usage: "Generates random tasks.",
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

					timeToPointer := func(t time.Time) *time.Time {
						return &t
					}

					// Create a task
					t := lib.NewTask("Get milk")
					t.Priority = lib.TaskPriorityLow
					t.Details = "Get 2% milk"
					if err = db.PutTask(t); err != nil {
						return err
					}

					t = lib.NewTask("Get bread")
					t.Priority = lib.TaskPriorityMedium
					t.CompletedAt = timeToPointer(time.Now())
					if err = db.PutTask(t); err != nil {
						return err
					}

					t = lib.NewTask("Get OJ")
					t.Details = "No Pulp!"
					if err = db.PutTask(t); err != nil {
						return err
					}

					t = lib.NewTask("Get coffee")
					if err = db.PutTask(t); err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "flush",
				Usage: "Deletes all tasks from the DB.",
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
					for _, t := range tasks {
						if err = db.DeleteTask(t.ID); err != nil {
							return err
						}
					}
					fmt.Printf("Deleted %d tasks.\n", len(tasks))
					return nil
				},
			},
		},
		Action: rootFunc,
	}
}
