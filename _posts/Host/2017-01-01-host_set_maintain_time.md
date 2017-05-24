---
category: Host
apiurl: '/api/v1/host/set_maintain_time'
title: "Set maintain time for hosts"
type: 'POST'
sample_doc: 'host.html'
layout: default
---

* [Session](#/authentication) Required
* hosts
  * []string
  * list of hosts
* start_time
  * int
  * unix time
  * start_time & endt_time to 0 mean unset maintain time
* end_time
  * int
  * unix time
  * start_time & endt_time to 0 mean unset maintain time


### Request
```{"start_time":1495447254,"hosts":["hostA","hostB"],"end_time":1495533650}```

### Response

```Status: 200```
```[
  {
    "id": 8,
    "hostname": "hostB",
    "ip": "0.0.0.1",
    "agent_version": "5.2.0",
    "plugin_version": "e1571b9944626b1aa7ab3075262e9b171854b4c7",
    "maintain_begin": 1495447254,
    "maintain_end": 1495533650
  },
  {
    "id": 1,
    "hostname": "hostA",
    "ip": "0.0.0.2",
    "agent_version": "5.1.4",
    "plugin_version": "12155256cec3926186de22e282e67f4ce11cdbf7",
    "maintain_begin": 1495447254,
    "maintain_end": 1495533650
  }
]```
