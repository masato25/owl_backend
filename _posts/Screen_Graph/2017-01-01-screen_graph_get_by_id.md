---
category: Screen_Graph
apiurl: '/api/v1/dashboard/graph/:id'
title: "Get graph by id"
type: 'GET'
sample_doc: 'owl_scren_graph.html'
layout: default
---

* [Session](#/authentication) Required

### Request

```
/api/v1/dashboard/graph/4913
```

### Response

```Status: 200```
```
{
  "counters": [
    "c1",
    "c2"
  ],
  "endpoints": [
    "a",
    "b"
  ],
  "falcon_tags": "",
  "graph_id": 4913,
  "graph_type": "h",
  "method": "",
  "position": 0,
  "screen_id": 1270,
  "timespan": 3600,
  "title": "no exist"
}
```
