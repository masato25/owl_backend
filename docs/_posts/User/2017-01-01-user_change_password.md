---
category: User
apiurl: '/api/v1/user/cgpasswd'
title: 'Change Password'
type: 'PUT'

layout: default
---

更新使用者
* 需要[Session](#/authentication)

### Request
```{
  "new_password": "test1",
  "old_password": "test1"
}```

### Response

```Status: 200```
```{"message":"password updated!"}```

For more example, see the [user](/doc/user.html).

For errors responses, see the [response status codes documentation](#/response-status-codes).
