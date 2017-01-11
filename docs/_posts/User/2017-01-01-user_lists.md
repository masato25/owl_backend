---
category: User
apiurl: '/api/v1/user/users'
title: 'List of Users'
type: 'GET'

layout: default
---

拿取使用者列表
* 需要[Session](#/authentication)

### Response

```Status: 200```
```[
  {
    "id": 1,
    "name": "root",
    "cnname": "",
    "email": "",
    "phone": "",
    "im": "",
    "qq": "904394234239",
    "role": 2
  },
  {
    "id": 32,
    "name": "owltester",
    "cnname": "翱鶚",
    "email": "root123@cepave.com",
    "phone": "99999999999",
    "im": "44955834958",
    "qq": "904394234239",
    "role": 0
  }
]```

For more example, see the [user](/doc/user.html).

For errors responses, see the [response status codes documentation](#/response-status-codes).
