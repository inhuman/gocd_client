package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
	"commands"
	"gocd"
)

func main() {

	gocd.Init("http://localhost:8153", "", "")

	app := cli.NewApp()
	app.Name = "gocd client"
	app.Usage = "managing pipelines, etc"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Commands = commands.Get()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
