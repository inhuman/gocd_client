#gocd-client

Client for gocd 

## Usage

#### Create pipeline
##### From file
```bash
$ gocd-client pipelines create --file path/to/createPipelineData.json
```
When using --file parameter for create pipeline other sub-command params will be ignored

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




To enable debug set env 
```bash
export GOCD_CLIENT_DEBUG=1
```