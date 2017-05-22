---
category: Screen
apiurl: '/api/v1/dashboard/screens'
title: "Create Screen"
type: 'POST'
sample_doc: 'owl_screen.html'
layout: default
---

* [Session](#/authentication) Required
* name
  * required
  * name must be unique, if exist will response error message to client
* pid
  * option
  * parent id [Inherit]
  * owl-light don't have this concept. this only for open-falcon

### Request

```
{"pid":0,"name":"screen_test2"}
```

### Response

```Status: 200```
```
{"id":1268,"pid":0,"name":"screen_test2"}
```
