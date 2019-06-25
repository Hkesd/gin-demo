gin框架虽然是开箱即用的，但是在我们写业务逻辑之前还是有必要做一些基础封装的，本文试图从入门的角度介绍如何使用基于gin构建出符合业务需求的API框架
<!-- more -->

## 系统环境&依赖
#### go
> go version go1.12.5 darwin/amd64
* 项目采用go mod来管理依赖

#### web
* `github.com/gin-gonic/gin`

#### 配置文件
* `yaml`

#### 日志
* `github.com/sirupsen/logrus`
* `github.com/lestrrat/go-file-rotatelogs`
* `github.com/rifflock/lfshook`

#### 数据库（orm）
* `github.com/jinzhu/gorm`  

## 基本思想
* 基于gin的基本功能，加上日志、orm以及一些文件夹的组织规范来构建该框架
#### 配置文件
* 采用yaml格式

#### 日志
* 采用`logrus`来构建我们的日志，但是logrus原生不支持文件分割，所以要通过hook机制来解决
* `app.log` 包含DebugLevel、InfoLevel
* `app-error.log` 包含WarnLevel、ErrorLevel、FatalLevel、PanicLevel

#### orm
* 采用gorm来实现orm功能

## 目录结构
```
.
├── Makefile
├── README.md
├── bin //dev 环境编译输出目录
│   └── logs
│       ├── app-error.log.2019-06-25
│       └── app.log.2019-06-25
├── common
│   ├── applog 
│   │   ├── gin-recovery.go //release 模式下recovery
│   │   └── log.go //日志处理
│   └── error
│       └── api-err-desc.go //API统一错误描述
├── config
│   └── config.go //配置文件解析
├── controller //控制器
│   └── base.go
├── dev.yaml
├── go.mod
├── go.sum
├── main.go //main
├── middleware //中间件
│   └── base.go
├── model //模型
│   └── model.go
├── release //release 编译输出
│   ├── config //release 配置文件
│   └── logs //release 日志路径
├── repository //数据仓库，原则上model只存放定义，操作模型放在数据仓库
├── route //路由
│   └── route.go
├── service //sevice 复杂的业务逻辑通过sevice来实现
└── util //工具
```

## 运行
#### 基本步骤
```
`cp .example.yaml dev.yaml` 然后修改dev.yaml

go run main.go -c dev.yaml
```

#### Makefile
```
make i

make r
```

## 后续要做的
* 增加参数校验
* 增加上线部署说明

## 代码地址
* [https://github.com/Hkesd/gin-demo]()