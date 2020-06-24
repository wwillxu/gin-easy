## gin-easy

> 基于gin的golang-web快速开发框架

### 环境

- Go 

- mongoDB 

### 特性

- 配置文件定义
- Restful API
- 基础中间件
- 参数验证
- 结构化返回

### 开始

```bash
$ git clone https://github.com/WillXu24/gin-easy.git
```

### 配置

```bash
$ cd gin-easy/config
$ vim .env
```

example.env

```env
#app
PORT = 8080

#database
DATABASE_URI = mongodb://localhost:27017
DATABASE_NAME = gin-easy

#jwt
JWT_KEY = gin-easy
```

### 运行

```bash
$ cd gin-easy
$ go run main.go
```

### API文档

> *表示需要认证，认证形式：header前加上Bearer字段，内容为登陆接口返回的token

#### POST /api/v1/user/register 用户注册

req

```json
{
    "username":"user",
    "password":"123456",
    "email":"user@eamil.com",
    "telephone":"666699999"
}
```

res

```json
{
    "code": 20000,
    "error": "",
    "data": null
}
```



#### POST /api/v1/user/login 用户登录

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
    "code": 20000,
    "error": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTM0MzU3NDIsImlkIjoiNWVmMzRlYzU0ZWM3ZTFhYzA4MzNhYTcwIn0.pGAui6_2Pd9H1517oi3B9BZCctEvTwjgt53iYxjAsMg"
    }
}
```



#### *GET /api/v1/user/profile 用户资料

res

```json
{
    "code": 20000,
    "error": "",
    "data": {
        "email": "user@eamil.com",
        "telephone": "666699999",
        "username": "user"
    }
}
```



#### *DELETE /api/v1/user/account 用户删除

res

```json
{
    "code": 20000,
    "error": "",
    "data": null
}
```

