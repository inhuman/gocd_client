package gocd

import (
	"github.com/inhuman/go-gocd"
)

var Client go_gocd.Client

func Init(host string, username string, password string) {
	Client = go_gocd.New(host, username, password)
}

