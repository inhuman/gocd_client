package main

import (
	"testing"
	"github.com/rendon/testcli"
	"os"
	"io/ioutil"
)

// before test:
// go build -v -i -x --ldflags '-extldflags "-static"' -o bin/gocd-client src/main.go
// export PATH=$PATH:$GOPATH/bin
// go test -v -cover ./src/...

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

func TestRunNoHostname(t *testing.T) {

	testcli.Run("gocd-client", "pipelines", "status", "--name", "test")
	if !testcli.Failure() {
		t.Fatalf("Expected to failure, but succeed: %s", testcli.Error())
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