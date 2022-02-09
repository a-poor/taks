package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const rootDesc = ``

// rootFunc runns the root CLI command
func rootFunc(c *cli.Context) error {
	fmt.Println("Running the root command...")
	return nil
}
