---
category: Alarm
apiurl: '/api/v1/alarm/event_note'
title: "Get status note of event case"
type: 'GET'
sample_doc: 'alarm.html'
layout: default
---

* [Session](#/authentication) Required
* 使用者可以对告警留言并人工切换状态
* 参数:
* startTime: 开始区间
* endTime: 结束区间
* status: 人工判定告警状态 ["in progress", "unresolved", "resolved", "ignored", "comment"]
* event_id: 某一单向告警id
* limit: 设定笔数返回上线 [预设及最大上线值:50]
* page: 后端分页页数

### Request

`/api/v1/alarm/event_note?startTime=1466956800&endTime=1480521600&event_id=s_165_cef145900bf4e2a4a0db8b85762b9cdb&status=ignored
`


### Response

```Status: 200```
```[
  {
    "event_caseId": "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
    "note": "test",
    "case_id": "",
    "status": "ignored",
    "timestamp": "2016-06-23T05:39:09+08:00",
    "user": "root"
  },
  {
    "event_caseId": "s_165_9d223f126e7ecb3477cd6806f1ee9656",
    "note": "Ignored by user",
    "case_id": "",
    "status": "ignored",
    "timestamp": "2016-06-23T05:38:56+08:00",
    "user": "root"
  }
]```
