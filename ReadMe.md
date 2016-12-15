# OWL_Backend
web api backend made for open-falcon


## Base
All error return code plaese [`refer here`](https://golang.org/src/net/http/status.go)

### Graph

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

#### Login User `/api/v1/user/login`
* [POST]
* ex. "name=xxx;password=xxx"
* Error: 400

#### Logout User `/api/v1/user/logout`
* [POST]
* [cookie]/[post_form] ex. "name=xxx;sig=xxx"
* Error: 400

#### Auth User `/api/v1/user/auth_session`
* [GET]
* [cookie]/[post_form] ex. "name=xxx;sig=xxx"
* Error: 401

#### Create User `/api/v1/user/create`
* [POST]
* ex. "name=xxx;password=xxx;cnname=xxx;email=xxx;phone=xxx;im=xxx;qq=xxx"

#### Update User `/api/v1/user/update`
* [PUT]

#### Change User password `/api/v1/user/cgpasswd`
* [PUT]

#### User list `/api/v1/user/users`
* [GET]
* supprot regexp ex. q=a.+

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/user.html)
### Admin

#### (Admin) Change User role `/api/v1/admim/change_user_role`
* [PUT]

#### (Admin) Change User password `/api/v1/admim//change_user_passwd`
* [PUT]

#### (Admin) Delete A user `/api/v1/admin/delete_user`
* [DELETE]

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/admin.html)

### Team

#### Get Team List `/api/v1/team`
* [GET]
* support regexp query ex. "q=a.+"

### Get Team by id `/api/v1/team/:team_id`
* [GET]
* ex. /api/v1/team/1

### Create A team  `/api/v1/team`
* [POST]
* ex. params
* team_name `name of team`
* resume `descript of this team`
* users `team member list (user id)`

### Update A team `/team`
* [PUT]
* team_id `team_id of db's primary key`
* resume `descript of this team`
* users `team member list (user id)`

### Delete A team `/team/:team_id`
* [DELETE]
* ex. /api/v1/team/1

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/team.html)

### Template / Strategy

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/doc/template.html)
