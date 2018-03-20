package commands

import (
	"gopkg.in/urfave/cli.v1"
	"utils"
	"gocd"
	"fmt"
)

func packagesCommand() cli.Command {

	return cli.Command{
		Name:        "packages",
		Aliases:     []string{"pkg"},
		Usage:       "options for packages",
		Subcommands: packagesSubCommands(),
	}
}

func packagesSubCommands() []cli.Command {

	commands := []cli.Command{
		{
			Name:   "create",
			Usage:  "create a new package",
			Action: packageSubCommandAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Package name, required",
				},
				cli.StringFlag{
					Name:  "package-repo-id",
					Usage: "Package repository id, where the package taken from ",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "delete an existing package",
			Action: packageSubCommandDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Package name, required",
				},
			},
		},
	}

	return commands
}

func packageSubCommandAdd(c *cli.Context) error {

	// flags to create package
	// id []
	// name [req]
	// auto-update [] default true
	// package-repo-id [req]
	// configuration [multi]


	filePath := c.String("file")

	if filePath != "" {
		err := gocd.CreatePipelineFromFile(filePath)
		if err != nil {
			return err
		}
	}

	template := c.String("template")

	if template != "" {

		utils.DebugMessage("Template flag found")

		err := gocd.CreatePipelineFromTemplate(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func packageSubCommandDelete(c *cli.Context) error {

	name := c.String("name")

	if name != "" {
		resp, err := gocd.DeletePackage(name)
		if err != nil {
			return err
		}
		fmt.Println(resp.Message)
	}

	return nil
}
