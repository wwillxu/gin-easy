## Gin Easy

> an easy gin example for golang beginner, which can help build web app very quickly.

### Environmental requirements

- Go
- MySQL

### Features

- RESTful API
- Gin
- Gorm
- Cron
- Jwt
- ini

### Install

```bash
$ git clone git@github.com:WillXu24/GinEasy.git gineasy
```

### Configuration

 `conf/example.ini`

```yaml
[app]
Port = :8080
RunMode = debug

[jwt]
Key = test_key

[database]
ConnectUri = user:password@/dbname?charset=utf8&parseTime=True&loc=Local
```

### Run

```bash
$ cd gineasy
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/ping              --> gineasy/routers.pingHandler (4 handlers)
[GIN-debug] POST   /api/v1/user              --> gineasy/controllers/v1.RegisterHandler (4 handlers)
[GIN-debug] POST   /api/v1/user/login        --> gineasy/controllers/v1.LoginHandler (4 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> gineasy/controllers/v1.UserProfileHandler (4 handlers)
[GIN-debug] POST   /api/v1/hello             --> gineasy/routers.pingHandler (5 handlers)
[GIN-debug] GET    /api/v1/me                --> gineasy/controllers/v1.MyProfileHandler (5 handlers)
[GIN-debug] POST   /api/v1/me                --> gineasy/controllers/v1.UpdateProfileHandler (5 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

### API Doc

#### Json Response

```json
// success
{
    "code": 200,
    "msg": "ok",
    "data": "success response"
}
// fail
{
    "code": 500,
    "msg": "error response",
    "data": "Null"
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

```json
// request
{
	"username":"user",
	"password":"123",
	"invite_code":"123"
}
// success response
{
    "code": 200,
    "msg": "ok",
    "data": "register success"
}
```

#### User Login： POST   /user/login

```json
// request
{
	"username":"user1",
	"password":"123"
}
// success response
{
    "code": 200,
    "msg": "ok",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjYyMDg3NzIsImlkIjoxLCJvcmlnX2lhdCI6MTU2NjE5Nzk3Mn0.qnuqqZ1bJ56AUQtc5C8-1oBDHKkmAzz4hE_OqErpfhs"
}
```

#### User Profile： GET   /user/:id

```json
// request:1
// success response
{
    "code": 200,
    "msg": "ok",
    "data": {
        "id": 1,
        "created_at": "2019-08-18T16:49:18+08:00",
        "username": "user1",
        "nickname": "anonymity",
        "introduction": "Hi!",
        "avatar": "123"
    }
}
```



#### Get My Profile：  GET   /me

```json
// request:nil
// success response
{
    "code": 200,
    "msg": "ok",
    "data": {
        "id": 1,
        "created_at": "2019-08-18T16:49:18+08:00",
        "username": "user1",
        "nickname": "anonymity",
        "introduction": "Hi!",
        "avatar": "123"
    }
}
```



#### Update My Profile： POST      /me

```json
// request
{
	"username":"user1",
	"password":"123",
	"nickname":"anonymity",
	"introduction":"Hi!",
	"avatar":"321"
}
// success response
{
    "code": 200,
    "msg": "ok",
    "data": "update success"
}
```

