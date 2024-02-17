package main

import (
	"github.com/CASP-Systems-BU/secrecy-go/cli"
	"github.com/CASP-Systems-BU/secrecy-go/config"
)

func init() {
	config.ReadEnvVariables()
}

// Creates a new CLI and runs it using the given command line arguments.
// Run go run main.go --help to see the available commands.
func main() {
	cli := cli.NewCli()
	cli.Initialize()
	cli.Run()
}
