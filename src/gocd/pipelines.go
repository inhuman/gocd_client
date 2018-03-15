package gocd

import (
	"github.com/inhuman/go-gocd"
	"os"
	"encoding/json"
	"utils"
	"gopkg.in/urfave/cli.v1"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"fmt"
	"regexp"
	"strings"
)

func GetPipelineGroups() ([]*gocd.PipelineGroup, error) {
	return Client.GetPipelineGroups()
}

func GetPipelineStatus(name string) (*gocd.PipelineStatus, error) {
	return Client.GetPipelineStatus(name)
}

func CreatePipelineFromFile(filePath string) error {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(file)

	var createPipelineData gocd.PipelineConfig

	if err = jsonParser.Decode(&createPipelineData); err != nil {
		return err
	}

	resp, err := Client.CreatePipeline(createPipelineData)
	if err != nil {
		return err
	}

	fmt.Println(err)

	utils.PrettyPrintStruct(resp)

	return nil
}

func CreatePipelineFromTemplate(c *cli.Context) error {

	//TODO: fix this to norm
	var tmp multierror.Error
	multiError := &tmp

	name := c.String("name")
	checkStringParam("name", name, multiError)

	group := c.String("group")
	checkStringParam("group", group, multiError)

	label := c.String("label")
	checkStringParam("label", label, multiError)

	lockBehavior := c.String("lock-behavior")
	checkStringParam("lockBehavior", lockBehavior, multiError)

	template := c.String("template")
	checkStringParam("template", template, multiError)

	materials := c.StringSlice("material")
	checkStringSliceParam("material", materials, multiError)

	envVarsSlice := c.StringSlice("var")
	checkEnvVarParam(envVarsSlice, multiError)

	envSecureVarsSlice := c.StringSlice("var-secure")
	checkEnvVarParam(envSecureVarsSlice, multiError)

	envSecureVars := hydrateSecureEnvVars(envSecureVarsSlice)
	envVars := hydrateEnvVars(envVarsSlice)
	allEnvVars := append(envSecureVars, envVars...)

	var materialAggregator []gocd.Material

	for _, path := range materials {
		material, err := getMaterialFromFile(path)

		if err != nil {
			multiError = multierror.Append(multiError, err)
		}
		materialAggregator = append(materialAggregator, *material)
	}

	if multiError.Errors != nil {
		return multiError.ErrorOrNil()
	}

	var pipelineConfig gocd.PipelineConfig

	pipelineConfig.Group = group
	pipelineConfig.Pipeline.Name = name
	pipelineConfig.Pipeline.LabelTemplate = label
	pipelineConfig.Pipeline.Template = template
	pipelineConfig.Pipeline.LockBehavior = lockBehavior
	pipelineConfig.Pipeline.Materials = materialAggregator
	pipelineConfig.Pipeline.EnvironmentVariables = allEnvVars

	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		utils.PrettyPrintStruct(pipelineConfig)
	}

	resp, err := Client.CreatePipeline(pipelineConfig)
	if err != nil {
		return err
	}
	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		utils.PrettyPrintStruct(resp)
	}

	fmt.Println("Pipeline " + name + " created.")
	return nil
}

func DeletePipeline(name string) error {
	return Client.DeletePipeline(name)
}

func getMaterialFromFile(filePath string) (*gocd.Material, error) {

	var material gocd.Material

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(file)

	if err = jsonParser.Decode(&material); err != nil {
		return nil, err
	}

	return &material, nil
}

func hydrateEnvVars(envVarsSlice []string) []gocd.EnvironmentVariable {
	var environmentVariables []gocd.EnvironmentVariable

	if len(envVarsSlice) > 0 {
		for _, envVar := range envVarsSlice {
			ar := strings.Split(envVar, "=")
			v := gocd.EnvironmentVariable{
				Name:  ar[0],
				Value: ar[1],
			}
			environmentVariables = append(environmentVariables, v)
		}
	}
	return environmentVariables
}

func hydrateSecureEnvVars(envSecureVarsSlice []string) []gocd.EnvironmentVariable {

	var environmentVariables []gocd.EnvironmentVariable

	if len(envSecureVarsSlice) > 0 {
		for _, envVar := range envSecureVarsSlice {
			ar := strings.Split(envVar, "=")
			v := gocd.EnvironmentVariable{
				Name:           ar[0],
				EncryptedValue: ar[1],
				Secure:         true,
			}
			environmentVariables = append(environmentVariables, v)
		}
	}
	return environmentVariables

}

//TODO: in future move checkalkas to separate local package
func checkStringParam(name string, value string, error error) {

	utils.DebugMessage("Check required value: " + name + " value is: " + value)

	if value == "" {
		utils.DebugMessage("Value is not set")
		error = multierror.Append(
			error, errors.New(
				fmt.Sprintf("Required parameter '%s' not set", name)))
	}
}

func checkStringSliceParam(name string, value []string, error error) {

	utils.DebugMessage("Check required value: " + name)

	if len(value) < 1 {
		error = multierror.Append(
			error, errors.New(
				fmt.Sprintf("Required parameter '%s' not set", name)))
	}

	for _, valueItem := range value {
		if valueItem == "" {
			utils.DebugMessage("Value is not set")
			error = multierror.Append(
				error, errors.New(
					fmt.Sprintf("Required parameter '%s' not set", name)))
		}
	}
}

func checkEnvVarParam(envVars []string, error error) {

	utils.DebugMessage("Check env vars")

	if len(envVars) > 0 {
		for _, envVar := range envVars {
			matched, err := regexp.MatchString("[A-Z\\d]+=.*", envVar)

			if err != nil {
				error = multierror.Append(error, err)
			}

			if !matched {
				utils.DebugMessage(fmt.Sprintf("Env var '%s' mismatched format", envVar))

				error = multierror.Append(
					error, errors.New(
						fmt.Sprintf("Env var '%s' mismatched format, see help for details", envVar)))
			}
		}
	} else {
		utils.DebugMessage("Env var is empty, skipping..")
	}
}
