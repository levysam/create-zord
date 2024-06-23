package main

import (
	"create-zord/internal"
	"create-zord/pkg"
)

func main() {
	cli := pkg.NewCli()
	cli.AddCommand(internal.CreateCliCommand())
	err := cli.Execute()
	if err != nil {
		panic(err)
	}
}
