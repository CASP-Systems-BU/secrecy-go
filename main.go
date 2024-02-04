package main

import "github.com/CASP-Systems-BU/secrecy-go/cli"

func main() {
	cli := cli.NewCli()
	cli.Initialize()
	cli.Run()
}
