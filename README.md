
![cover.png](https://i.loli.net/2019/08/13/k46WnDILMwutj9s.png)

​												[![GitHub stars](https://img.shields.io/github/stars/WillXu24/GoEasy)](https://github.com/WillXu24/GoEasy/stargazers) [![GitHub forks](https://img.shields.io/github/forks/WillXu24/GoEasy)](https://github.com/WillXu24/GoEasy/network) [![GitHub issues](https://img.shields.io/github/issues/WillXu24/GoEasy)](https://github.com/WillXu24/GoEasy/issues) [![GitHub license](https://img.shields.io/github/license/WillXu24/GoEasy)](https://github.com/WillXu24/GoEasy)	

> an easy gin example for golang beginner, which can help build web app very quickly.

[中文文档](https://github.com/WillXu24/GoEasy/blob/master/README-zh.md)

### Environmental requirements

- Go
- MongoDB

### Features

- RESTful API
- Gin
- Cron
- Jwt
- Conf configurable

### Install

```bash
$ git clone git@github.com:WillXu24/GoEasy.git
```

### Configuration

```ini
MONGODB_URI="mongodb://localhost:27017"
MONGODB_DBNAME="mydb"
JWT_REALM="Test"
JWT_SIGN_METHOD="HS256"
JWT_KEY="TestKey"
PORT=":8080"
GIN_MODE="debug"
```

### Run

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

### API Doc

#### Json Response

```json
//success
{
    "Success": true,
    "Error": "",
    "Data": "response data"
}
//fail
{
    "Success": false,
    "Error": "errer response",
    "Data": NULL 
}
```

#### User Author

- Method： JSON Web Token

- How： put token into HTTP  header of Authorization and begin with Bearer

- example：

  example token

  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NTk1NTEsImlkIjoiNWQ0ZWVlYjdhNTBmYTMxMmU0ODYyYzdiIiwib3JpZ19pYXQiOjE1NjU0NTU5NTF9.gU_OsqXlAcFWS5qN7sbGVqQdrFVOkAUOoPffTH6q9sk
  ```

  HTTP header

  ```
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0NTk1NTEsImlkIjoiNWQ0ZWVlYjdhNTBmYTMxMmU0ODYyYzdiIiwib3JpZ19pYXQiOjE1NjU0NTU5NTF9.gU_OsqXlAcFWS5qN7sbGVqQdrFVOkAUOoPffTH6q9sk
  ```

#### User Register： POST    /user

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

#### User Login： POST   /user/login

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

#### Get Profile： GET   /user/:id

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
        "Nickname": "jack",
        "Introduction": "I'm Jack！",
        "TotalCollCount": 0,
        "TotalFollowCount": 0,
        "CreateTime": 1565580285
    }
}
```



#### My Profile：  GET   /me

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
        "Nickname": "jack",
        "Introduction": "I'm Jack！",
        "TotalCollCount": 0,
        "TotalFollowCount": 0,
        "CreateTime": 1565580285
    }
}
```

#### Update My Profile： POST      /me

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
