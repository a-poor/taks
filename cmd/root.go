package cmd

import (
	"fmt"
	"os"

	"github.com/a-poor/taks/lib"
	"github.com/a-poor/taks/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

const rootDesc = ``

// rootFunc runns the root CLI command
func rootFunc(c *cli.Context) error {
	fmt.Println("Running the root command...")

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

	// Initialize the TUI model
	m, err := tui.NewModel(db)
	if err != nil {
		return err
	}

	// Run the TUI
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}

	return nil
}
