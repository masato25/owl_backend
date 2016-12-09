# OWL_Backend
web api backend made for open-falcon


## Base
All error return code plaese [`refer here`](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html)

### Graph

#### Get endpoint by regexp `/api/v1/graph/endpoint?q=${string}`
* [GET]
* ex. http://localhost:3000/api/v1/graph/endpoint?q=b.%2B
* Error 401, 400

#### Get metric count list base on endpoint_id `/api/v1/graph/endpoint_counter?eid=${string}&metricQuery=${string}`
* [GET]
* ex. http://localhost:3000/api/v1/graph/endpoint_counter?eid=285533184,73482468&metricQuery=d.%B
* Error 401, 400

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

For more api please visit [`here`](https://htmlpreview.github.io/?https://github.com/masato25/owl_backend/blob/master/apidoc.html)
