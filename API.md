### API文档

> *表示需要认证，认证形式：header前加上Bearer字段，内容为登陆接口返回的token

#### POST /api/v1/register 用户注册

req

```json
{
  "username": "user",
  "password": "123456",
  "email": "user@eamil.com",
  "telephone": "18612341234"
}
```

res

```json
{
  "success": true,
  "error": "",
  "data": null
}
```

#### POST /api/v1/login 用户登录

req

```json
{
  "username": "user",
  "password": "123456"
}
```

res

```json
{
  "success": true,
  "error": "",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTEyOTY0NDAsImlkIjoiNjAwM2Q3MmIwZWIxZTI5NTZiNDA3NjMzIn0.MCq6vFY64hQs6Rg1059401IH7Gw1E30NCShUolkZ4wM"
  }
}
```

#### *GET /api/v1/user 用户资料

res

```json
{
  "success": true,
  "error": "",
  "data": {
    "email": "user@eamil.com",
    "telephone": "18612341234",
    "username": "user"
  }
}
```

#### *DELETE /api/v1/user 用户删除

res

```json
{
  "success": true,
  "error": "",
  "data": null
}
```

