package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

// run_configure is a function that sets the default secrecy configuration parameters.
func (c *Cli) run_configure(ctx *cli.Context) error {
	fmt.Println("Configuring defaults ...")

	return nil
}

// run_configure_add_party is a function that adds a party to the secrecy configuration.
func (c *Cli) run_configure_add_party(ctx *cli.Context) error {
	fmt.Println("Adding a party ...")

	return nil
}

// run_configure_add_cluster is a function that adds a cluster to the secrecy configuration.
func (c *Cli) run_configure_add_cluster(ctx *cli.Context) error {
	fmt.Println("Adding a cluster ...")

	return nil
}
