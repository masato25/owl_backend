# OWL_Backend
web api backend made for open-falcon


## Base
All error return code plaese [`refer here`](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html)
### Graph

#### Get endpoint by regexp `/endpoint?q=${string}`
* [GET]
* ex. http://localhost:3000/endpoint?q=b.%2B

#### Get metric count list base on endpoint_id `/endpoint_counter?eid=${string}`
* [GET]
* ex. http://localhost:3000/endpoint_counter?eid=285533184,73482468


### User

#### Login User `/user/login`
* [POST]
* ex. "name=xxx;password=xxx"

#### Login User `/user/logout`
* [POST]
* [cookie]/[post_form] ex. "name=xxx;sig=xxx"
