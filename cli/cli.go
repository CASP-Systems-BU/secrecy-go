package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

// Cli is a struct that holds the Command Line Interface application.
type Cli struct {
	App *cli.App
}

// NewCli creates a new instance of the Command Line Interface.
func NewCli() *Cli {
	return &Cli{
		App: cli.NewApp(),
	}
}

// Initialize Creates the CLI application and sets the commands and flags.
// The commands are:
// - configure: configures secrecy setup and computing cluster management.
// - local: manages the secrecy framework on a local machine and sends requests to it.
// - serve: starts the secrecy-go service.
// - client: sends API requests to the secrecy-go service.
// - execute: runs a secure multi-party computation across the specified computing parties.
func (c *Cli) Initialize() {
	c.App.Name = "secrecy-go"
	c.App.Usage = "A go package to manage the secrecy MPC framework across the computing parties."
	c.App.Version = "1.0.0"
	c.App.Commands = []cli.Command{
		{
			Name: "configure",
			// Aliases: []string{"l"},
			Usage:  "Configures the secrecy-go package defaults (~/.secrecy-go/config.env).",
			Action: c.run_configure,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "secrecy",
					Value: "~/.secrecy-go/secrecy",
					Usage: "Specifies the `PATH` to the installed secrecy source code.",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "party",
					Usage:  "Adds a computing party configuration.",
					Action: c.run_configure_add_party,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Value: "0",
							Usage: "The computing party `ID`.",
						},
						cli.StringFlag{
							Name:  "address",
							Value: "localhost",
							Usage: "The computing party `ADDRESS`.",
						},
						cli.StringFlag{
							Name:  "port",
							Value: "8080",
							Usage: "The computing party secrecy-go service `PORT`.",
						},
					},
				},
				{
					Name:   "cluster",
					Usage:  "Adds a cluster configuration.",
					Action: c.run_configure_add_cluster,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Value: "0",
							Usage: "The cluster `ID`.",
						},
						cli.StringSliceFlag{
							Name:  "parties",
							Usage: "The cluster `PARTIES`.",
						},
					},
				},
			},
		},
		{
			Name: "local",
			// Aliases: []string{"l"},
			Usage:  "Manages a local computing party and sends requests to it.",
			Action: c.run_local,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Value: "mpc.cpp",
					Usage: "The MPC `FILE` to be executed.",
				},
				cli.StringFlag{
					Name:  "cluster",
					Value: "0",
					Usage: "The cluster `ID`.",
				},
				cli.StringFlag{
					Name:  "threads",
					Value: "1",
					Usage: "The number of `THREADS` to be used.",
				},
				cli.StringFlag{
					Name:  "batch",
					Value: "1",
					Usage: "The `BATCH` size.",
				},
				cli.StringFlag{
					Name:  "num",
					Value: "1",
					Usage: "The number of computing `PARTIES`.",
				},
				cli.StringSliceFlag{
					Name:  "flags",
					Usage: "The compilation `FLAGS` to be used.",
				},
			},
		},
		{
			Name: "serve",
			// Aliases: []string{"s"},
			Usage:  "Starts the computing party service.",
			Action: c.run_server,
		},
		{
			Name: "client",
			// Aliases: []string{"c"},
			Usage:  "Sends API requests to the computing party service.",
			Action: c.run_client,
		},
		{
			Name: "execute",
			// Aliases: []string{"e"},
			Usage:  "Runs a secure multi-party computation across the specified computing parties.",
			Action: c.run_execute,
		},
	}

	c.App.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "http://localhost:8080",
			Usage: "The host address for the computing party service.",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8080",
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
}

// Run starts the CLI application.
func (c *Cli) Run() {
	c.App.RunAndExitOnError()
}
