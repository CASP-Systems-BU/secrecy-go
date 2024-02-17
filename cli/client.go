package cli

import (
	"fmt"

	"github.com/CASP-Systems-BU/secrecy-go/api"
	"github.com/urfave/cli"
)

// run_client is a function that sends requests to the service.
func (c *Cli) run_client(ctx *cli.Context) error {
	fmt.Println("Sending requests to the service...")

	client := api.NewClient(ctx.GlobalString("host"))
	response, err := client.PostPing("2020-01-01T00:00:00Z")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(response)

	return nil
}
