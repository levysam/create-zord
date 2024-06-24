package main

import (
	"github.com/levysam/create-zord/internal"
	"github.com/levysam/create-zord/pkg"
)

func main() {
	cli := pkg.NewCli()
	cli.AddCommand(internal.CreateCliCommand())
	err := cli.Execute()
	if err != nil {
		panic(err)
	}
}
