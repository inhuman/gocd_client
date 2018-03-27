package gocd

import (
	"os"
	"encoding/json"
	"utils"
	"fmt"
	"strings"
	"gopkg.in/urfave/cli.v1"
	"github.com/hashicorp/go-multierror"
	"github.com/inhuman/go-gocd"
)

func GetPipelineGroups() ([]*go_gocd.PipelineGroup, error) {
	Init()
	return Client.GetPipelineGroups()
}

func GetPipelineStatus(name string) (*go_gocd.PipelineStatus, error) {
	Init()
	return Client.GetPipelineStatus(name)
}

func CreatePipelineFromFile(filePath string) error {
	Init()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(file)

	var createPipelineData go_gocd.PipelineConfig

	if err = jsonParser.Decode(&createPipelineData); err != nil {
		return err
	}

	resp, err := Client.CreatePipeline(createPipelineData)
	if err != nil {
		return err
	}

	fmt.Println(resp.Message)

	return nil
}

func CreatePipelineFromTemplate(c *cli.Context) error {
	Init()

	//TODO: fix this to norm
	var tmp multierror.Error
	multiError := &tmp

	name := c.String("name")
	utils.CheckStringParam("name", name, multiError)

	group := c.String("group")
	utils.CheckStringParam("group", group, multiError)

	label := c.String("label")
	utils.CheckStringParam("label", label, multiError)

	lockBehavior := c.String("lock-behavior")
	utils.CheckStringParam("lockBehavior", lockBehavior, multiError)

	template := c.String("template")
	utils.CheckStringParam("template", template, multiError)

	materials := c.StringSlice("material")
	utils.CheckStringSliceParam("material", materials, multiError)

	envVarsSlice := c.StringSlice("var")
	utils.CheckEnvVarParam(envVarsSlice, multiError)

	envSecureVarsSlice := c.StringSlice("var-secure")
	utils.CheckEnvVarParam(envSecureVarsSlice, multiError)

	envSecureVars := hydrateSecureEnvVars(envSecureVarsSlice)
	envVars := hydrateEnvVars(envVarsSlice)
	allEnvVars := append(envSecureVars, envVars...)

	var materialAggregator []go_gocd.PipelineMaterial

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

	var pipelineConfig go_gocd.PipelineConfig

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

func DeletePipeline(name string) (*go_gocd.ApiResponse, *multierror.Error) {
	Init()
	return Client.DeletePipeline(name)
}


func getMaterialFromFile(filePath string) (*go_gocd.PipelineMaterial, error) {

	var material go_gocd.PipelineMaterial

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

func hydrateEnvVars(envVarsSlice []string) []go_gocd.EnvironmentVariable {
	var environmentVariables []go_gocd.EnvironmentVariable

	if len(envVarsSlice) > 0 {
		for _, envVar := range envVarsSlice {
			ar := strings.Split(envVar, "=")
			v := go_gocd.EnvironmentVariable{
				Name:  ar[0],
				Value: ar[1],
			}
			environmentVariables = append(environmentVariables, v)
		}
	}
	return environmentVariables
}

func hydrateSecureEnvVars(envSecureVarsSlice []string) []go_gocd.EnvironmentVariable {

	var environmentVariables []go_gocd.EnvironmentVariable

	if len(envSecureVarsSlice) > 0 {
		for _, envVar := range envSecureVarsSlice {
			ar := strings.Split(envVar, "=")
			v := go_gocd.EnvironmentVariable{
				Name:   ar[0],
				Value:  ar[1],
				Secure: true,
			}
			environmentVariables = append(environmentVariables, v)
		}
	}
	return environmentVariables

}
