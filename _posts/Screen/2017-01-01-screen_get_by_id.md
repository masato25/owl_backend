---
category: Screen
apiurl: '/api/v1/dashboard/screen/:id'
title: "Get Screen by id"
type: 'GET'
sample_doc: 'owl_screen.html'
layout: default
---

* [Session](#/authentication) Required


### Request

```
/api/v1/dashboard/screen/1268
```

### Response

```Status: 200```
```
{
  "graphs": [
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
      "graph_id": 4917,
      "graph_type": "h",
      "method": "",
      "position": 0,
      "screen_id": 1274,
      "timespan": 3600,
      "title": "grep_demo"
    }
  ],
  "scren": {
    "id": 1274,
    "pid": 0,
    "name": "screen_test2"
  }
}
```
