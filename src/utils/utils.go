package utils

import (
	"github.com/hokaccha/go-prettyjson"
	"fmt"
	"strings"
	"regexp"
	"os"
	"github.com/hashicorp/go-multierror"
	"errors"
)

func PrettyPrintStruct(strct interface{}) {

	s, _ := prettyjson.Marshal(strct)
	fmt.Println(string(s))
}

func DebugMessage(string string) {
	if os.Getenv("GOCD_CLIENT_DEBUG") == "1" {
		fmt.Println(string)
	}
}

func InArrayRegexp(string string, regexpArr []string) bool {

	for _, regex := range regexpArr {

		var re = regexp.MustCompile(regex)

		if re.Match([]byte(string)) {
			return true
		}
	}

	return false
}

func ParseDcService(serviceName string) (string, string) {

	stringArr := strings.Split(serviceName, "[")
	stringArr2 := strings.Split(stringArr[1], "]")

	return stringArr2[0], stringArr2[1]
}

func CheckStringParam(name string, value string, error *multierror.Error) {

	DebugMessage("Check required value: " + name + " value is: " + value)

	if value == "" {
		DebugMessage("Value is not set")
		error = multierror.Append(
			error, errors.New(
				fmt.Sprintf("Required parameter '%s' not set", name)))
	}
}

func CheckStringSliceParam(name string, value []string, error *multierror.Error) {

	DebugMessage("Check required value: " + name)

	if len(value) < 1 {
		error = multierror.Append(
			error, errors.New(
				fmt.Sprintf("Required parameter '%s' not set", name)))
	}

	for _, valueItem := range value {
		if valueItem == "" {
			DebugMessage("Value is not set")
			error = multierror.Append(
				error, errors.New(
					fmt.Sprintf("Required parameter '%s' not set", name)))
		}
	}
}

func CheckEnvVarParam(envVars []string, error *multierror.Error) {

	DebugMessage("Check env vars")

	if len(envVars) > 0 {
		for _, envVar := range envVars {
			matched, err := regexp.MatchString("[A-Z\\d]+=.*", envVar)

			if err != nil {
				error = multierror.Append(error, err)
			}

			if !matched {
				DebugMessage(fmt.Sprintf("Env var '%s' mismatched format", envVar))

				error = multierror.Append(
					error, errors.New(
						fmt.Sprintf("Env var '%s' mismatched format, see help for details", envVar)))
			}
		}
	} else {
		DebugMessage("Env var is empty, skipping..")
	}
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
