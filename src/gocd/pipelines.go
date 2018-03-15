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

	var createPipelineData gocd.CreatePipelineData

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

	var multiError *multierror.Error

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

	var materialAggregator []gocd.Material

	for _, path := range materials {
		material, err := getMaterialFromFile(path)

		if err != nil {
			multiError = multierror.Append(multiError, err)
		}
		materialAggregator = append(materialAggregator, *material)
	}

	if multiError != nil {
		return multiError.ErrorOrNil()
	}

	var createPipelineData gocd.CreatePipelineData

	createPipelineData.Group = group
	createPipelineData.Pipeline.Name = name
	createPipelineData.Pipeline.LabelTemplate = label
	createPipelineData.Pipeline.Template = template
	createPipelineData.Pipeline.LockBehavior = lockBehavior
	createPipelineData.Pipeline.Materials = materialAggregator

	resp, err := Client.CreatePipeline(createPipelineData)
	if err != nil {
		return err
	}

	utils.PrettyPrintStruct(resp)

	return nil
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
