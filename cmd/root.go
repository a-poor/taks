package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// cliRoot runns the root CLI command
func cliRoot(c *cli.Context) error {
	fmt.Println("Running the root command...")
	return nil
}
