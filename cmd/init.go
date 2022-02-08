package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/a-poor/taks/lib"
	"github.com/urfave/cli/v2"
)

const initDesc = `Initializes a new task database. If one already exists
the user will be prompted to overwrite it.`

// cliInit corresponds to the CLI's init command.
// It initializes the taks database and prepares it for use.
func cliInit(c *cli.Context) error {
	// Temp...
	fmt.Println("Initializing the taks database...")

	// Get the config directory location
	home := "" // TODO: Get the home directory from a user-suppled argument.
	cfg, err := lib.GetConfigDir(home)
	if err != nil {
		return fmt.Errorf("unable to get user's home directory: %w", err)
	}

	// Does the config directory already exist?
	_, err = os.Stat(cfg)
	exists := os.IsExist(err)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("unable to stat config directory: %w", err)
	}

	// If the path exists, already, ask the user if they want to overwrite it.
	var overwrite bool
	if exists {
		fmt.Printf("The config directory '%s' already exists.\n", cfg)
		survey.AskOne(
			&survey.Confirm{
				Message: "Do you want to overwrite it?",
				Default: false,
			},
			&overwrite,
		)
	}

	// If the path exists and the user doesn't want to overwrite it, exit.
	if exists && !overwrite {
		fmt.Println("Exiting...")
		return nil
	}

	// If the path exists and the user wants to overwrite it, delete it.
	if exists && overwrite {
		fmt.Printf("Deleting the existing config directory %q...\n", cfg)
		err = os.RemoveAll(cfg)
		if err != nil {
			return fmt.Errorf("failed to delete the existing config directory: %w", err)
		}
	}

	// Now, the config directory doesn't exist, so create it.
	fmt.Printf("Creating the config directory %q...\n", cfg)
	db, err := lib.OpenDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to create the taks database: %w", err)
	}

	// Now, create the metadata and store it in the DB
	meta := lib.NewMetadata()
	err = db.PutMetadata(meta)
	if err != nil {
		return fmt.Errorf("failed to store the metadata: %w", err)
	}
	fmt.Println("Done.")

	return nil
}
