package cli

import (
	"fmt"

	"github.com/CASP-Systems-BU/secrecy-go/api"
	"github.com/urfave/cli"
)

type Cli struct {
	App *cli.App
}

func NewCli() *Cli {
	return &Cli{
		App: cli.NewApp(),
	}
}

func (c *Cli) Initialize() {
	c.App.Name = "secrecy-go"
	c.App.Usage = "A go package to manage the secrecy MPC framework across the computing parties."
	c.App.Version = "1.0.0"
	c.App.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Starts the computing party service.",
			Action:  c.run_server,
		},
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "Sends API requests to the computing party service.",
			Action:  c.run_client,
		},
		{
			Name:    "local",
			Aliases: []string{"l"},
			Usage:   "Runs a local computing party and sends requests to it.",
			Action:  c.run_local,
		},
	}

	c.App.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "http://localhost:8080",
			Usage: "The host of the computing party service.",
		},
	}

	c.App.Before = func(c *cli.Context) error {
		fmt.Println("Initializing...")
		return nil
	}

	c.App.After = func(c *cli.Context) error {
		fmt.Println("Done.")
		return nil
	}

	// c.App.Action = func(c *cli.Context) error {
	// 	fmt.Println("No command specified.")
	// 	return nil
	// }

}

func (c *Cli) run_server(ctx *cli.Context) error {
	fmt.Println("Starting the service...")

	service := api.NewService()
	service.InitializeRoutes()
	service.Run()

	return nil
}

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

func (c *Cli) run_local(ctx *cli.Context) error {
	fmt.Println("Starting the service...")

	service := api.NewService()
	service.InitializeRoutes()
	go service.Run()

	fmt.Println("Sending requests to the service...")

	client := api.NewClient("http://localhost:8080")
	response, err := client.PostPing("2020-01-01T00:00:00Z")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(response)

	return nil
}

func (c *Cli) Run() {
	c.App.RunAndExitOnError()
}
