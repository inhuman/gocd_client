package gocd

import (
	"github.com/inhuman/go-gocd"
	"github.com/hashicorp/go-multierror"
)

func DeletePackage(id string) (*go_gocd.ApiResponse, *multierror.Error) {
	Init()
	return Client.DeletePackage(id)
}
