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
```

### 运行

```bash
$ cd gin-easy
$ go run main.go
```

