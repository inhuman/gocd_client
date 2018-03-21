package gocd

import (
	"github.com/inhuman/go-gocd"
	"github.com/hashicorp/go-multierror"
)

func DeletePackage(id string) (*go_gocd.ApiResponse, *multierror.Error) {
	Init()
	return Client.DeletePackage(id)
}

func CreatePackage(pkg go_gocd.Package) (*go_gocd.Package, *go_gocd.ApiResponse, *multierror.Error) {
	Init()
	return Client.CreatePackage(pkg)
}

func GetPackageRepoByName(repoName string) (*go_gocd.Repository, error) {
	Init()

	repos, err := Client.GetAllRepositories()

	if err != nil {
		return nil, err.ErrorOrNil()
	}

	for _, repo := range repos.Embedded.PackageRepositories {
		if repo.Name == repoName {
			return &repo, nil
		}
	}

	return nil, nil
}
