# go-ecommerce

Simple, fast Golang Ecommerce API.

**Features**

- Authentication with E-mail & Google
- Users
- Products

**Planned Features**

- Shopping Cart
- Collections

## Deployment

```
docker build -t samjtro/go-ecommerce-api .
docker run -it -p 8080:8080 --env-file=variables.env samjtro/go-ecommerce-api
```