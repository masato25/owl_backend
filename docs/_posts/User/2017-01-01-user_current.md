---
category: User
apiurl: '/api/v1/user/current'
title: 'Current User info'
type: 'GET'

layout: default
---

拿取當前使用者資訊
* 需要[Session](#/authentication)

### Response

```Status: 200```
```{
  "id": 2,
  "name": "root",
  "cnname": "",
  "email": "",
  "phone": "",
  "im": "",
  "qq": "",
  "role": 2
}```

For more example, see the [user](/doc/user.html).

For errors responses, see the [response status codes documentation](#/response-status-codes).
