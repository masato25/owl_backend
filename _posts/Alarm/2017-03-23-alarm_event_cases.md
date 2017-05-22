---
category: Alarm
apiurl: '/api/v1/alarm/eventcases'
title: "Get event cases list"
type: 'GET'
sample_doc: 'alarm.html'
layout: default
---

* [Session](#/authentication) Required
* 支援GET && POST
* GET: "/api/v1/alarm/eventcases?startTime=1466956800&endTime=1480521600"
* POST: `form post` & `json body post` 皆可以使用
* 一个告警设置只会产生一个eventcases, 每一次的告警历史会储存于events中
* 参数:
* startTime: 开始区间
* endTime: 结束区间
* priority: 优先等级 [0 ~ 4]
* status: 系统判定告警状态 ["OK", "PROBLEM"] *逗号分隔支援多笔查询
* process_status: 人工判定告警状态 ["in progress", "unresolved", "resolved", "ignored"] *逗号分隔支援多笔查询
* metrics: 告警监控项搜寻 [regexp搜寻]
* event_id: 某一单向告警id [拿取单笔告告警状态]
* limit: 设定笔数返回上线, 如果搭配page一起使用表示单页的上线,如果page无设定表示单页返回上线. [预设及最大上线值 => 页面: 50 / 单页: 2000]
* page: 后端分页页数

### Request

```{
 	startTime: 1466956800,
 	endTime: 1480521600,
 	priority: 0,
 	status: "OK,PROBLEM",
 	process_status: "unresolved",
 	metrics: "cpu.+",
 	event_id: "s_165_cef145900bf4e2a4a0db8b85762b9cdb",
 	limit: 50,
 	page: 1
}```

### Response

```Status: 200```
```[
  {
    "id": "s_46_1ac45122afb893adc02fbd30154ac303",
    "endpoint": "agent4",
    "metric": "cpu.iowait",
    "func": "all(#3)",
    "cond": "48.33759590792839 > 40",
    "note": "CPU I/O wait超过40",
    "step": 1,
    "current_step": 1,
    "priority": 1,
    "status": "PROBLEM",
    "timestamp": "2016-08-01T06:25:00+08:00",
    "update_at": "2016-08-01T06:25:00+08:00",
    "closed_at": null,
    "closed_note": "",
    "user_modified": 0,
    "tpl_creator": "root",
    "expression_id": 0,
    "strategy_id": 46,
    "template_id": 126,
    "process_note": 16907,
    "process_status": "ignored"
  },
  {
    "id": "s_50_6438ac68b30e2712fb8f00d894c46e21",
    "endpoint": "agent5",
    "metric": "cpu.idle",
    "func": "avg(#3)",
    "cond": "95.16331658291456 <= 98",
    "note": "cpu空闲值报警",
    "step": 1,
    "current_step": 1,
    "priority": 3,
    "status": "PROBLEM",
    "timestamp": "2016-07-03T16:13:00+08:00",
    "update_at": "2016-07-03T16:13:00+08:00",
    "closed_at": null,
    "closed_note": "",
    "user_modified": 0,
    "tpl_creator": "root",
    "expression_id": 0,
    "strategy_id": 50,
    "template_id": 53,
    "process_note": 1181,
    "process_status": "ignored"
  }
]```
