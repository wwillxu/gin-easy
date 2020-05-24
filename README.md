## gin-easy

> 基于gin的golang-web快速开发框架

### 环境

- Go 

- mongoDB 

### 特性

- Restful API
- 结构化输出
- 配置文件定义
- 参数验证
- 跨域、JWT中间件
- 自动化API文档

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
DATABASE_NAME = test_db

#jwt
JWT_KEY = test_key
```

### 运行

```bash
$ cd gin-easy
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /docs/*any                --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (4 handlers)
[GIN-debug] POST   /user/register            --> gin-easy/controllers.UserRegisterHandler (4 handlers)
[GIN-debug] POST   /user/login               --> gin-easy/controllers.UserLoginHandler (4 handlers)
[GIN-debug] POST   /user/delete              --> gin-easy/controllers.UserDeleteHandler (5 handlers)
[GIN-debug] Environment variable PORT="8080"
[GIN-debug] Listening and serving HTTP on :8080
```

### API文档

运行后浏览器访问http://localhost:8080/docs/index.html