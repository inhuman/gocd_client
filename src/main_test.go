package main

import (
	"testing"
	"github.com/rendon/testcli"
	"os"
	"io/ioutil"
)

// before test
// build app
// export PATH

func TestRunWithNoParams(t *testing.T) {

	testcli.Run("gocd-client")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	f, err := os.Open("../tests/responses/response_no_param")
	if err != nil {
		t.Fatal(err)
	}
	file, err := ioutil.ReadAll(f)

	if !testcli.StdoutContains(string(file)) {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), file)
	}

}

//TODO fix tests

func TestRunNoHostname(t *testing.T) {

	testcli.Run("gocd-client")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	f, err := os.Open("../tests/responses/without_hostname")
	if err != nil {
		t.Fatal(err)
	}
	file, err := ioutil.ReadAll(f)

	if !testcli.StdoutContains(string(file)) {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), file)
	}

}