# TRD: API Management

## API Design

```
GET /customer?keyword=xxx&limit=10&cursor=
POST /customer
GET /customer/:id
PUT /customer/:id
DELETE /customer/:id

GET /order?keyword
POST /order
GET /order/:id
PUT /order/:id
DELETE /order/:id

POST /auth/login
POST /auth/register
```

## ERD

![erd](https://github.com/rezaig/dbo-service/assets/121090402/1a5d2c94-18e5-4c76-ada8-3fb323a90b1b)

```
Table customer {
  id bigint [pk]
  name varchar(255)
  email varchar(255)
  phone_number varchar(20)
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Table account {
  id bigint [pk]
  username varchar(255)
  password varchar(255)
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Table order {
  id bigint [pk]
  product_id bigint
  customer_id bigint
  quantity int
  shipping_address varchar(255)
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Table product {
  id bigint [pk]
  name varchar(255)
  description varchar(255)
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}

Ref: order.product_id > product.id
Ref: order.product_id > customer.id
```