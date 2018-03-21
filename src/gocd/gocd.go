package gocd

import (
	"github.com/inhuman/go-gocd"
	"fmt"
	"os"
	"errors"
)

var Client go_gocd.Client

var host, username, password string

func Init() {

	if len(host) < 1 {
		fmt.Fprintf(os.Stdout, "GOCD client error: %v\n", errors.New("gocd server hostname not set, see help for details"))
		os.Exit(1)
	}

	Client = go_gocd.New(host, username, password)
}

func SetClientConfig(Host, Username, Password string) {
	host = Host
	username = Username
	password = Password
}
