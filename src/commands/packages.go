package commands

import (
	"gopkg.in/urfave/cli.v1"
	"utils"
	"gocd"
	"fmt"
	"config"
	"github.com/inhuman/go-gocd"
	"github.com/hashicorp/go-multierror"
	"strings"
	"os"
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
				cli.StringFlag{
					Name:  "spec",
					Usage: "Package spec, e.g configuration PACKAGE_SPEC, required",
				},
				cli.StringSliceFlag{
					Name:  "configuration",
					Usage: "Configuration key value field. Format --configuration 'PACKAGE_SPEC=package-2.8-1.fc20.src'",
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


	var tmp multierror.Error
	multiError := &tmp

	name := c.String("name")
	utils.CheckStringParam("name", name, multiError)

	// If package id not set, use name
	id := c.String("id")
	if id == "" {
		id = name
	}

	repoName := c.String("repo")
	if repoName == "" {
		repoName = config.Config.Repository.Rpm.Default
	}
	repo, err := gocd.GetPackageRepoByName(repoName)
	if err != nil {
		return err
	}

	spec := c.String("spec")
	utils.CheckStringParam("spec", spec, multiError)

	if multiError.Errors != nil {
		return multiError.ErrorOrNil()
	}

	configurationSlice := c.StringSlice("configuration")

	pkg := go_gocd.Package{}
	pkg.Name = name
	pkg.Id = name
	pkg.AutoUpdate = !c.Bool("disable-auto-update")
	pkg.Configuration = hydrateConfigurationVars(configurationSlice)
	pkg.Configuration = addPackageSpec(spec, pkg.Configuration)
	pkg.PackageRepo.Id = repo.RepoId

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println("Prepared package")
		utils.PrettyPrintStruct(pkg)
	}

	_, resp, err := gocd.CreatePackage(pkg)
	if err != nil {
		multiError = multierror.Append(multiError, err)
	}

	if multiError.Errors != nil {
		return multiError.ErrorOrNil()
	}

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		utils.PrettyPrintStruct(resp)
	}

	fmt.Println("Package " + name + " created.")

	return nil
}

func addPackageSpec(spec string, configuration []go_gocd.Configuration) []go_gocd.Configuration {

	specStr := go_gocd.Configuration{
		Key: "PACKAGE_SPEC",
		Value: spec,
	}

	return append(configuration, specStr)
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

func hydrateConfigurationVars(varsSlice []string) []go_gocd.Configuration {
	var configuration []go_gocd.Configuration

	if len(varsSlice) > 0 {
		for _, envVar := range varsSlice {

			ar := strings.Split(envVar, "=")

			v := go_gocd.Configuration{}
			v.Key = ar[0]
			v.Value = ar[1]

			configuration = append(configuration, v)
		}
	}
	return configuration
}
