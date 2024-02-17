package cli

import (
	"fmt"

	"github.com/CASP-Systems-BU/secrecy-go/api"
	"github.com/urfave/cli"
)

// run_server is a function that starts the service that controls execution for an MPC party.
func (c *Cli) run_server(ctx *cli.Context) error {
	fmt.Println("Starting the service...")

	service := api.NewService()
	service.InitializeRoutes()
	service.Run()

	return nil
}
