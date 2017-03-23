---
category: Alarm
apiurl: '/api/v1/alarm/event_note'
title: "Add status note to a event case"
type: 'POST'
sample_doc: 'alarm.html'
layout: default
---

* [Session](#/authentication) Required
* 使用者可以对告警留言并人工切换状态
* 参数:
* event_id: 某一单向告警id
* status: 人工判定告警状态 ["in progress", "unresolved", "resolved", "ignored", "comment"], comment之外的留言会改变event_case的process_status状态
* note: 对于告警留言
* case_id: 填入外部对应的系统公单号

### Request

```{
	"event_id": "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
	"note": "closed case",
	"case_id": "k00001",
	"status": "resolved"
}```

### Response

```Status: 200```
```{
  "id": "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
  "message": "add note to s_165_cef145900bf4e2a4a0db8b85762b9cdb successfuled"
}```
