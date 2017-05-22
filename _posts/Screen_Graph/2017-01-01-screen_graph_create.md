---
category: Screen_Graph
apiurl: '/api/v1/dashboard/graph'
title: "Create a graph and bind to exist screen"
type: 'POST'
sample_doc: 'owl_scren_graph.html'
layout: default
---

* [Session](#/authentication) Required
* screen_id
  * int64 [required]
* title
  * string [required]
  * name of graph
* endpoints
  * []string [required]
* counters
  * []string [required]
* timespan
  * int64
  * 时间区段 (秒)
    * default 3600
* graph_type
  * string [required]
  * 视角: h (endpoint view), k (counter view), a (combo view)
  * accept values:
    * h
    * a
    * k
* method
  * string
  * accept values:
    * SUM  
    * 空值
* position
  * int64
  * 排序
    * 预设值為0
* falcon_tags
  * string
  * owl-light not this concept, keep empty. (open-falcon only)

### Request

```
{
  "title": "imtitle",
  "screen_id": 1266,
  "endpoints": ["a-host","b-host"],
  "counters": ["c.1","c.2"],
  "graph_type": 'h'
}
```

### Response

```Status: 200```
```
{"id":4912}
```
* return created graph id
