# OWL_Backend

![Alt text](data/owl_backend_icon.jpg)

web api backend made for open-falcon

开源版本 open-falcon 後台, 将以下模组的操作功能api化:
* Fe
  * 使用者,群组管理
* Portal
  * HostGroup, Host 管理
  * Template, 告警策略管理
* Graph
  * Host/Counter列表查询
* No-Data
  * NoData策略设置api
* Dashboard
  * 具备数据查询接口功能

## Base
All error return code plaese refer [`here`](https://golang.org/src/net/http/status.go)

### Dashboard

#### Get endpoint by regexp `/api/v1/graph/endpoint?q=${string}`
* [GET]
* ex. http://localhost:3000/api/v1/graph/endpoint?q=b.%2B
* Error 401, 400

#### Get metric count list base on endpoint_id `/api/v1/graph/endpoint_counter?eid=${string}&metricQuery=${string}`
* [GET]
* ex. http://localhost:3000/api/v1/graph/endpoint_counter?eid=285533184,73482468&metricQuery=d.%B
* Error 401, 400

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/graph.html)

### User

#### Login user `/api/v1/user/login`
* [POST]
* ex. "name=xxx;password=xxx"
* Error: 400

#### Logout user `/api/v1/user/logout`
* [POST]
* [cookie]/[post_form] ex. "name=xxx;sig=xxx"
* Error: 400

#### Auth user / Session checking `/api/v1/user/auth_session`
* [GET]
* [cookie]/[post_form] ex. "name=xxx;sig=xxx"
* Error: 401

#### Create user `/api/v1/user/create`
* [POST]
* ex. "name=xxx;password=xxx;cnname=xxx;email=xxx;phone=xxx;im=xxx;qq=xxx"

#### Update user `/api/v1/user/update`
* [PUT]

#### Change user password `/api/v1/user/cgpasswd`
* [PUT]

#### User list `/api/v1/user/users`
* [GET]
* supprot regexp ex. q=a.+

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/user.html)
### Admin

#### (Admin) Change user role `/api/v1/admim/change_user_role`
* [PUT]

#### (Admin) Change user password `/api/v1/admim//change_user_passwd`
* [PUT]

#### (Admin) Delete a user `/api/v1/admin/delete_user`
* [DELETE]

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/admin.html)

### Team

#### Get team List `/api/v1/team`
* [GET]
* support regexp query ex. "q=a.+"

#### Get team by id `/api/v1/team/:team_id`
* [GET]
* ex. /api/v1/team/1

#### Create a team  `/api/v1/team`
* [POST]
* ex. params
* team_name `name of team`
* resume `descript of this team`
* users `team member list (user id)`

#### Update a team `/team`
* [PUT]
* team_id `team_id of db's primary key`
* resume `descript of this team`
* users `team member list (user id)`

#### Delete a team `/team/:team_id`
* [DELETE]
* ex. /api/v1/team/1

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/team.html)

#### Template / Strategy

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/template.html)

#### HostGroup / Host

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/hostgroup.html) *hostgroup

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/host.html) *host


### Plugin
#### Unbind Plugin of a hostgroup / Delete Plugin: `/plugin/:plugin_id`
* [DELETE]
* One Plugin only relate one hostgroup, so if you want unbind a plugin of a hostgroup, just use `delete` method to clean it.

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/plugin.html)

### Aggreator

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/aggreator.html)

### Expressions
For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/expression.html)

### NoData
For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/nodata.html)

# ToDo
* Screen

ps.
跨域问题實做可以參考 `sample_for_cross_api`
