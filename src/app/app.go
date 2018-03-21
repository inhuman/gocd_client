package app

import (
	"gopkg.in/urfave/cli.v1"
	"commands"
	"gocd"
	"config"
)

//http://localhost:8153

func Init(cliApp *cli.App) error {

	cliApp.Name = "gocd client"
	cliApp.Usage = "managing pipelines, creating packages,  etc"
	cliApp.Version = "0.0.1"
	cliApp.EnableBashCompletion = true

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			EnvVar: "GOCD_HOST",
			Usage:  "Gocd server url",
		},
		cli.StringFlag{
			Name:   "user",
			EnvVar: "GOCD_USERNAME",
			Usage:  "Gocd user name",
		},
		cli.StringFlag{
			Name:   "password",
			EnvVar: "GOCD_PASSWORD",
			Usage:  "Gocd password",
		},
	}

	cliApp.Commands = commands.Get()
	cliApp.Before = beforeRun

	return nil
}

func beforeRun(c *cli.Context) error {

	host := c.String("host")
	username := c.String("username")
	password := c.String("password")
	gocd.SetClientConfig(host, username, password)

	err := config.Init()
	if err != nil {
		return err
	}

	return nil
}
