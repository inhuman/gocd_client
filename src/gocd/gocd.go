package gocd

import (
	"github.com/inhuman/go-gocd"
)

var Client go_gocd.Client

func Init(host, username, password string) error {
	Client = go_gocd.New(host, username, password)
	return nil
}
