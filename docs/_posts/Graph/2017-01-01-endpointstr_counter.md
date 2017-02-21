---
category: Graph
apiurl: '/api/v1/graph/endpointsrt_counter'
title: "Get Counter of Endpoint by endpoint names"
type: 'GET'
sample_doc: 'graph.html'
layout: default
---

* [Session](#/authentication) Required
* params:
  * endpoints: endpoint name lists
  * q: 使用 regex 查询字符
    * option 参数

### Response

```Status: 200```
```[
  "df.inodes.free.percent/fstype=ext4,mount=/",
  "disk.io.read_sectors/device=sda",
  "disk.io.write_merged/device=sdh",
  "net.if.total.packets/iface=eth4",
  "snmp.Udp.RcvbufErrors"
]```
