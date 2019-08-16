# 配置文件模块实现

go-profile 灵感源自 Spring framework，并采用了 [github.com/joho/godotenv](https://github.com/joho/godotenv) 库的实现，
从而得以完成配置文件的灵活设置

## 使用方式

```
go get -u github.com/ltyyz/goprofile
```

将配置文件放在 config/ 文件夹下，默认导入的文件为 application.env，一个典型的配置文件样例：

```
server.port=8888
server.key=12345678
```

在程序中使用：

```go
goprofile.Load()

os.Getenv("server.port")
```

来获取配置文件中的值

> 由于 godotenv 使用了环境变量来实现配置文件，所以要避开系统环境变量名（前缀多一点）

## 多环境配置文件

实现方式类似于 Spring，采用传入参数 profile 来实现多环境配置

比如开发环境启动参数为：-profile=dev 程序自动会导入：

- config/application.dev.env
- config/application.env

两个文件的配置

同样支持多文件导入：-profile=dev,db 程序自动导入：

- config/application.dev.env
- config/application.db.env
- config/application.env

三个文件的配置

> 注意优先级：applicatiohn.dev.env > applicatiohn.db.env > applicatiohn.env，
> 相同 key 前者会覆盖后者


## 最佳实践

开发环境不用传参数，配置文件统一使用 application.env

测试环境参数：-profile=test，配置文件使用 application.test.env application.env 前者会覆盖后者

正式环境参数：-profile=prod，配置文件使用 application.prod.env application.env 前者会覆盖后者
