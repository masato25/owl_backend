---
category: Screen
apiurl: '/api/v1/dashboard/screens'
title: "Get Screen List"
type: 'GET'
sample_doc: 'owl_screen.html'
layout: default
---

* [Session](#/authentication) Required
* limit
* page
* desc
  * value: true or false
  * Order by desc on id
* key_word
  * search key words of screen name

### Request

```
/api/v1/screens?limit=10&page=1&desc=true&key_word=test
```

### Response

```Status: 200```
```
{
  "current_page": 1,
  "data": [
    {
      "graph_names": [],
      "id": 962,
      "pid": 0,
      "name": "FCM",
      "creator": "root"
    },
    {
      "graph_names": [
        "net.if.total.bytes",
        "net.if.total.bytes"
      ],
      "id": 965,
      "pid": 962,
      "name": "net.if.total.bytes",
      "creator": "root"
    },
    {
      "graph_names": [],
      "id": 967,
      "pid": 0,
      "name": "c01.i36",
      "creator": "root"
    },
    {
      "graph_names": [
        "cpu",
        "mysql_test",
        "cpu-1"
      ],
      "id": 968,
      "pid": 967,
      "name": "cpu",
      "creator": "root"
    },
    {
      "graph_names": [
        "cpu"
      ],
      "id": 969,
      "pid": 962,
      "name": "cpu",
      "creator": "root"
    },
    {
      "graph_names": [
        "net.if.total.bytes"
      ],
      "id": 970,
      "pid": 967,
      "name": "net.if.total.bytes",
      "creator": "root"
    },
    {
      "graph_names": [
        "mem"
      ],
      "id": 971,
      "pid": 967,
      "name": "mem",
      "creator": "root"
    },
    {
      "graph_names": [
        "mem"
      ],
      "id": 972,
      "pid": 962,
      "name": "mem",
      "creator": "root"
    },
    {
      "graph_names": [
        "cpu.load"
      ],
      "id": 973,
      "pid": 962,
      "name": "cpu.load",
      "creator": "root"
    },
    {
      "graph_names": [
        "cpu.load"
      ],
      "id": 975,
      "pid": 967,
      "name": "cpu.load",
      "creator": "root"
    }
  ],
  "desc_order": false,
  "key_word": "",
  "order_by": "id",
  "totall_count": 123,
  "totall_page": 13
}
```
