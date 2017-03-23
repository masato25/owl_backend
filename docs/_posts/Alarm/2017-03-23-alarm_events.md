---
category: Alarm
apiurl: '/api/v1/alarm/events'
title: "Get events list from one event case"
type: 'GET'
sample_doc: 'alarm.html'
layout: default
---

* [Session](#/authentication) Required
* 支援GET && POST
* GET: "/api/v1/alarm/events?startTime=1466956800&endTime=1480521600"
* POST: `form post` & `json body post` 皆可以使用
* 用于储存每一次的告警历史, 触发及恢复时机点
* 参数:
* startTime: 开始区间
* endTime: 结束区间
* status: 系统判定告警状态 [0, 1] 0表触发, 1表恢复
* event_id: 某一单向告警id [拿取单笔告告警状态]
* limit: 设定笔数返回上线 [预设及最大上线值:50]
* page: 后端分页页数

### Request

```{
 	startTime: 1466956800,
 	endTime: 1480521600,
 	status: 0,
 	event_id: "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
 	limit: 50,
 	page: 1
}```

### Response

```Status: 200```
```[
  {
    "id": 635166,
    "event_caseId": "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
    "step": 0,
    "cond": "10.649350649350648 != 66",
    "status": 0,
    "timestamp": "2016-06-23T04:55:00+08:00"
  },
  {
    "id": 635149,
    "event_caseId": "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
    "step": 0,
    "cond": "13.486005089058525 != 66",
    "status": 0,
    "timestamp": "2016-06-23T04:50:00+08:00"
  }
]```
