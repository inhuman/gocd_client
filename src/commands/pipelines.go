package commands

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"gocd"
	"utils"
)

func pipelinesCommand() cli.Command {

	return cli.Command{
		Name:        "pipelines",
		Aliases:     []string{"p"},
		Usage:       "options for pipelines",
		Subcommands: pipelinesSubCommands(),
	}
}

func pipelinesSubCommands() []cli.Command {

	commands := []cli.Command{
		{
			Name:   "add",
			Usage:  "add a new pipeline",
			Action: pipelineSubCommandAdd,
		},
		{
			Name:   "delete",
			Usage:  "delete an existing pipeline",
			Action: pipelineSubCommandDelete,
		},
		{
			Name:   "groups",
			Usage:  "shows groups of pipelines",
			Action: pipelineSubCommandGroups,
		},
		{
			Name:  "status",
			Usage: "shows status of pipeline",
			Flags: []cli.Flag{
				cli.StringFlag{
					"name",
					"required flag, sets name of pipeline",
					"",
					false,
					"",
					nil},
			},
			Action: pipelineSubCommandStatus,
		},
	}

	return commands
}

func pipelineSubCommandAdd(c *cli.Context) error {
	fmt.Println("new pipeline:", c.Args().First())
	return nil
}

func pipelineSubCommandDelete(c *cli.Context) error {
	fmt.Println("deleteting pipeline:", c.Args().First())
	return nil
}

func pipelineSubCommandGroups(c *cli.Context) error {

	pipelineGroups, err := gocd.GetPipelineGroups()

	if err != nil {
		return err
	}
	utils.PrettyPrintStruct(pipelineGroups)

	return nil
}

func pipelineSubCommandStatus(c *cli.Context) error {

	//TODO: think about what do if no name given

	name := c.String("name")
	if name != "" {
		status, err := gocd.GetPipelineStatus(name)

		if err != nil {
			return err
		}

		utils.PrettyPrintStruct(status)
	}

	return nil
}

//TODO: think about pretty format for pipeline groups
//TODO: implement create/delete group
//TODO: implement create/delete pipeline
