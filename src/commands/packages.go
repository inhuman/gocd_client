package commands

import (
	"gopkg.in/urfave/cli.v1"
	"utils"
	"gocd"
	"fmt"
	"config"
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
					Name:  "repo",
					Usage: "Package repository name, where the package taken from. Command to see list of repos search in help",
					Value: config.Config.Repository.Rpm.Default,
				},
				cli.BoolFlag{
					Name:  "disable-auto-update",
					Usage: "Disable auto update package",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Package id, by default uses package name",
				},
				cli.StringSliceFlag{
					Name: "configuration",
					Usage: "Configuration key value field. Format --configuration 'PACKAGE_SPEC=package_name3' ",
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
