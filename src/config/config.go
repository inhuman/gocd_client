package config

import (
	"utils"
)

type AppConfig struct {
	Repository struct {
		Rpm struct {
			Default string
		}
	}
}

var Config AppConfig

func Init() error {
	config := AppConfig{}
	config.Repository.Rpm.Default = utils.Getenv("GOCD_DEFAULT_RPM_REPOSITORY", "artifactory-rpm")
	Config = config
	return nil
}
