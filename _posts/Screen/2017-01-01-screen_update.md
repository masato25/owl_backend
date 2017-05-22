---
category: Screen
apiurl: '/api/v1/dashboard/screens'
title: "Update Screen"
type: 'PUT'
sample_doc: 'owl_screen.html'
layout: default
---

* [Session](#/authentication) Required
* id
  * required
* name
  * required
  * name must be unique, if exist will response error message to client

### Request

```
{"name":"updaredscreen","id":1268}
```

### Response

```Status: 200```
```
{"message":"ok"}
```
