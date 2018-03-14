#gocd-client

Client for gocd 

## Usage

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