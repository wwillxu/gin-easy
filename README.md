## gin-easy

> 基于gin的golang-web快速开发框架

### 环境

- Go

- mongoDB

### 特性

- 自定义配置
- Restful API
- 基础中间件
- 序列化返回

### 开始

```bash
$ git clone https://github.com/willxu24/gin-easy.git
```

### 配置

```bash
$ cd gin-easy/config
$ vim example.yml
```

example.yml

```yml
app:
  port: ":8080"
  prefix: "/api/v1"

db:
  uri: "mongodb://root:1234@localhost:27017"
  name: "gin-easy"

jwt:
  secret: "gin-easy"

dev: true
```

### 运行

```bash
$ cd gin-easy
$ go run main.go
```

### API文档

> *表示需要认证，认证形式：header前加上Bearer字段，内容为登陆接口返回的token

#### POST /api/v1/register 用户注册

req

```json
{
    "username":"user",
    "password":"123456",
    "email":"user@eamil.com",
    "telephone":"18612341234"
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
    "username":"user",
    "password":"123456"
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

