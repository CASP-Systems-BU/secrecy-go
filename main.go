package main

import (
	"github.com/CASP-Systems-BU/secrecy-go/cli"
	"github.com/CASP-Systems-BU/secrecy-go/config"
)

func init() {
	config.ReadEnvVariables()
}

func main() {
	cli := cli.NewCli()
	cli.Initialize()
	cli.Run()
}
