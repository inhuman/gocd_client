package commands

import (
	"gopkg.in/urfave/cli.v1"
	"utils"
	"gocd"
	"fmt"
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
			Name:   "create",
			Usage:  "create a new pipeline",
			Action: pipelineSubCommandAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "Path to json file with data for create pipeline",
				},
				cli.StringFlag{
					Name:  "template",
					Usage: "Template name, uses with other required params",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Pipeline name, required, uses with --template",
				},
				cli.StringFlag{
					Name:  "group",
					Usage: "Group name, required, uses with --template",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Instance label, required, uses with --template",
				},
				cli.StringFlag{
					Name:  "lock-behavior",
					Usage: "Lock behavior, uses with --template. Default 'unlockWhenFinished'",
					Value: "unlockWhenFinished",
				},
				cli.StringSliceFlag{
					Name:  "material",
					Usage: "Material, required, uses with --template. Multi-flag",
				},
				cli.StringSliceFlag{
					Name:  "var",
					Usage: "Environment variable, uses with --template. Multi-flag. Format: --var 'USERNAME=admin'",
				},
				cli.StringSliceFlag{
					Name:  "var-secure",
					Usage: "Environment variable, uses with --template. Multi-flag. Format: --var-secure 'PASSWORD=1f3rrs9uhn63hd'",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "delete an existing pipeline",
			Action: pipelineSubCommandDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Pipeline name, required",
				},
			},
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
					Name:        "name",
					Usage:       "Required flag, pipeline name",
				},
			},
			Action: pipelineSubCommandStatus,
		},
	}

	return commands
}

func pipelineSubCommandAdd(c *cli.Context) error {

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

func pipelineSubCommandDelete(c *cli.Context) error {

	name := c.String("name")

	if name != "" {
		resp, err := gocd.DeletePipeline(name)
		if err != nil {
			return err
		}
		fmt.Println(resp.Message)

	} else {
		//TODO: fire error or something
	}

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
