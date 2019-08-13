# GoEasy

基于Gin的简单整合框架，帮助初学者快速搭建web应用

### 环境

- Go
- MongoDB

### 特性

- RESTful API
- Gin
- Cron
- Jwt
- Conf configurable

### 安装

```bash
$ git clone git@github.com:WillXu24/GoEasy.git
```

### 配置

```ini
MONGODB_URI="mongodb://localhost:27017"
MONGODB_DBNAME="mydb"
JWT_REALM="Test"
JWT_SIGN_METHOD="HS256"
JWT_KEY="TestKey"
PORT=":8080"
GIN_MODE="debug"
```

### 运行

```bash
$ cd GoEasy
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/ping              --> OnlinePhotoAlbum/routers.pingHandler (4 handlers)
[GIN-debug] POST   /api/v1/user              --> OnlinePhotoAlbum/controllers/v1.RegisterHandler (4 handlers)
[GIN-debug] POST   /api/v1/user/login        --> OnlinePhotoAlbum/controllers/v1.LoginHandler (4 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> OnlinePhotoAlbum/controllers/v1.UserProfileHandler (4 handlers)
[GIN-debug] GET    /api/v1/me                --> OnlinePhotoAlbum/controllers/v1.MyProfileHandler (5 handlers)
[GIN-debug] POST   /api/v1/me                --> OnlinePhotoAlbum/controllers/v1.UpdateProfileHandler (5 handlers)
[GIN-debug] Listening and serving HTTP on :8080

```

### API 文档

#### json响应格式

```json
//成功
{
    "Success": true,
    "Error": "错误信息",
    "Data": ""
}
//失败
{
    "Success": false,
    "Error": "",
    "Data": "响应数据"
}
```

#### 用户身份认证

- 方式：JSON Web Token

- 备注：将token放入http头的Authorization字段，并使用Bearer开头

- 举例：

  若token为

  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NTk1NTEsImlkIjoiNWQ0ZWVlYjdhNTBmYTMxMmU0ODYyYzdiIiwib3JpZ19pYXQiOjE1NjU0NTU5NTF9.gU_OsqXlAcFWS5qN7sbGVqQdrFVOkAUOoPffTH6q9sk
  ```

  则HTTP头为

  ```
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NTk1NTEsImlkIjoiNWQ0ZWVlYjdhNTBmYTMxMmU0ODYyYzdiIiwib3JpZ19pYXQiOjE1NjU0NTU5NTF9.gU_OsqXlAcFWS5qN7sbGVqQdrFVOkAUOoPffTH6q9sk
  ```

#### 用户注册： POST    /user

request

```json
{
    "username":"13011112222",
    "password":"123",
    "invite_code":"456"
}
```

response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": "register success"
}
```

#### 用户登录： POST   /user/login

request

```json
{
    "username":"13011112222",
    "password":"123",
}

```

response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NTk1NTEsImlkIjoiNWQ0ZWVlYjdhNTBmYTMxMmU0ODYyYzdiIiwib3JpZ19pYXQiOjE1NjU0NTU5NTF9.gU_OsqXlAcFWS5qN7sbGVqQdrFVOkAUOoPffTH6q9sk"
}
```

#### 查看资料： GET   /user/:id

request  `:id`

response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": {
        "Id": "5d50dbfddb518a0bdfa4b818",
        "Username": "user1",
        "Nickname": "沙雕本雕",
        "Introduction": "说点什么介绍自己吧！",
        "TotalCollCount": 0,
        "TotalFollowCount": 0,
        "CreateTime": 1565580285
    }
}
```



#### 我的资料：  GET   /me

no request 

response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": {
        "Id": "5d50dbfddb518a0bdfa4b818",
        "Username": "user1",
        "Nickname": "沙雕本雕",
        "Introduction": "说点什么介绍自己吧！",
        "TotalCollCount": 0,
        "TotalFollowCount": 0,
        "CreateTime": 1565580285
    }
}
```

#### 修改资料： POST      /me

request

```json
{
    "username":"13011112222",
    "password":"123",
    "introduction":"apple"
}

```

response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": "update success"
}
```
