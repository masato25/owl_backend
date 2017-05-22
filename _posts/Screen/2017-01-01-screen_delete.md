---
category: Screen
apiurl: '/api/v1/dashboard/screen/:id'
title: "Delete Screen by id"
type: 'DELETE'
sample_doc: 'owl_screen.html'
layout: default
---

* [Session](#/authentication) Required


### Request

```
/api/v1/screen/1268
```

### Response

```Status: 200```
```
{
  "deleted_rows": 1,
  "deleted_graph_ids": [
    4901
  ],
  "message": "ok"
}
```
