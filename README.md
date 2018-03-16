#gocd-client

Client for gocd 

## Usage

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

##### Pipline status
```bash
$ gocd-client pipelines status --name pipeline1 
{
  "locked": false,
  "paused": false,
  "pausedBy": "",
  "pausedCause": "",
  "schedulable": false
}
```
##### Delete pipeline
```bash
$ gocd-client pipelines delete --name FromTemplate3
```

To enable debug set env 
```bash
export GOCD_CLIENT_DEBUG=1
```