package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"app"
	"fmt"
)

func main() {

	cliApp := cli.NewApp()

	err := app.Init(cliApp)
	if err != nil {
		fmt.Fprintf(os.Stdout, "ERROR: %v\n", err)
		os.Exit(1)
	}

	err = cliApp.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stdout, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
