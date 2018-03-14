package gocd

import (
	"github.com/inhuman/go-gocd"
)

var Client gocd.Client

func Init(host string, username string, password string) {
	Client = gocd.New(host, username, password)
}

func GetPipelineGroups() ([]*gocd.PipelineGroup, error) {
	return Client.GetPipelineGroups()
}

func GetPipelineStatus(name string) (*gocd.PipelineStatus, error) {
	return Client.GetPipelineStatus(name)
}