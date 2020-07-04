## gin-easy

> 基于gin的golang-web快速开发框架



### 环境

- Go 

- mongoDB 



### 特性

- 自定义配置
- Restful API
- 基础中间件
- 参数验证器
- 序列化返回



### 开始

```bash
$ git clone https://github.com/willxu24/gin-easy.git
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
DATABASE_URI = mongodb://localhost:27017 # 地址格式：mongodb://[user]:[password]@[ip]:[port]
DATABASE_NAME = gin-easy

#jwt
JWT_KEY = gin-easy
```



### 运行

```bash
$ cd gin-easy
$ go run main.go
```



### 移植开发

```bash
$ cd gin-easy
$ rm -r ./.git/
```

替换`go.mod`以及7个`.go`文件中共计16个`gin-easy`字段为你的工程名



### API文档

> *表示需要认证，认证形式：header前加上Bearer字段，内容为登陆接口返回的token

#### POST /api/v1/register 用户注册

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
    "code": 20000,
    "error": "",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTM0MzU3NDIsImlkIjoiNWVmMzRlYzU0ZWM3ZTFhYzA4MzNhYTcwIn0.pGAui6_2Pd9H1517oi3B9BZCctEvTwjgt53iYxjAsMg"
    }
}
```



#### *GET /api/v1/user 用户资料

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



#### *DELETE /api/v1/user 用户删除

res

```json
{
    "code": 20000,
    "error": "",
    "data": null
}
```

