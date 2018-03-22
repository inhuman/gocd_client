#gocd-client

Client for gocd 

## Usage
#### Auth
_Authorization params can be given by params on every client call_ 

Without user credentials
```bash 
$ gocd-client --host example pipelines ...
```
With user credentials
```bash
$ gocd-client --host example --user admin --password hardPassword pipelines ...
```

_Authorization params can be exported as env variables_
```bash
$ export GOCD_HOST=http://localhost:8153
$ export GOCD_USERNAME=admin
$ export GOCD_PASSWORD=hardPassword
$ gocd-client pipelines ...
``` 

#### Create pipeline
##### From file
```bash
$ gocd-client pipelines create --file path/to/createPipelineData.json
```
When using --file parameter for create pipeline other sub-command params will be ignored

##### From template
```bash
$ gocd-client pipelines create --template test_template --name FromTemplate3 --group first --label 'git-${COUNT}'  --material tests/fixtures/material1.json --material tests/fixtures/material2.json --var 'ADF=123' --var-secure 'PASSWORD=234'
```

##### Delete pipeline
```bash
$ gocd-client pipelines delete --name FromTemplate3
```

#### Packages 
##### Add new package
Minimum params
```bash
$ gocd-client package create --name TestPackage --spec TestPackage-1.0.1-fc32.src
```
Full list
```bash
$ gocd-client package create --name TestPackage --spec TestPackage-1.0.1-fc32.src --disable-auto-update --id 5035e11a-fd2d-44df-b73e-6f8fa49a4275 --configuration "key1=value1" --configuration "key2=value2" --repo repo-name
```

##### Delete package
Deleting package only possible if it not associated with any pipeline
```bash
$ gocd-client package delete --name TestPackage
```

#### Env vars to config params
Enable debug set env 
```bash
export GOCD_CLIENT_DEBUG=1
```

Change default rpm repo name (default: artifactory-rpm)
```bash
export GOCD_DEFAULT_RPM_REPOSITORY=artifactory-new-rpm
```