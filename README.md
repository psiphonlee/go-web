# 目录结构

    ├── api
    |   ├── v1 # v1版本接口服务
    |       ├── system # 系统级服务
    |       └── enter.go # 统一入口
    ├── config # 配置信息相关结构体
    ├── core # 核心组件初始化代码
    ├── dao # dao层
    ├── global # 全局变量
    ├── initialize # 配置启动初始化
    ├── middleware # 中间件
    ├── model # 数据库结构体
    ├── router # 路由
    ├── service # 业务层
    ├── utils # 工具函数
    ├── config.yaml # 配置文件
    ├── go.mod # 包管理
    ├── main.go # 项目启动文件
    └── README.md # 项目README

# go gin 框架后台基础组件部署

> 参考连接：https://juejin.cn/post/7213297003869569081

## 1.初始化项目，引入 gin "github.com/gin-gonic/gin"

## 2.1 配置配置文件

2.1.1 根目录新建 config.yaml 文件，添加基本的配置信息<br>
2.2.2 根目录下新建 config 文件夹，用于存放所有配置对应的结构体:<br>
&emsp;&emsp;config 下新建 config.go 文件，
定义 Configuration 结构体，其 App 属性对应 config.yaml 中的 app 配置信息

## 2.2 viper 读取配置文件，并放入全局变量 "github.com/spf13/viper" "github.com/fsnotify/fsnotify"

2.2.1 根目录新建 global 文件夹，并创建 global.go 文件来存放全局变量<br>
2.2.2 import "gomap/config" , "github.com/spf13/viper"<br>
2.2.3 声明全局配置文件： EWA_CONFIG config.Configuration 以及 全局 viper：EWA_VIPER \*viper.Viper<br>
2.2.4 考虑实际工作中多环境开发、测试的场景，我们需要针对不同的环境使用不同的配置：<br>
&emsp;&emsp;根目录新建 core 文件夹 及 internal 子文件夹，internal 目录下新建 一个 constants.go，声明各个环境常量<br>
2.2.5 core 目录下新建 viper.go，编写 viper 配置初始化方法<br>
2.2.6 InitializeViper(path)函数可以以多种方式来指定 yaml 配置文件启动服务:<br>
&emsp;&emsp;1- path 传参:<br>
&emsp;&emsp;InitializeViper("config.yaml")<br>
&emsp;&emsp;2- 通过命令行传参给 flag:<br>
&emsp;&emsp;flag 格式：flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)<br>
&emsp;&emsp;命令行输入格式：go run main.go -c config.yaml , -c 后面的就是指定的配置文件名<br>
&emsp;&emsp;3- 命令行设置环境变量来指定 yaml 文件，注意 windows 和 linux 下格式区别:<br>
&emsp;&emsp;windows: $env:GVA_CONFIG="config.yaml"; go run main.go<br>
&emsp;&emsp;linux: GVA_CONFIG=config.yaml go run main.go<br>
&emsp;&emsp;4- main.go 初始化的时候通过 gin.SetMode(AppMode)设置:<br>
&emsp;&emsp;gin.SetMode("debug" or "test" or "release") 来指定<br>
&emsp;&emsp;优先级: 命令行 > 环境变量 > 默认值<br>

## 2.3 zap 日志模块封装 "go.uber.org/zap" "go.uber.org/zap/zapcore"

2.3.4- 在 config.yarml 中新增 zap 配置信息<br>
2.3.2 config 目录下新建 zap.go，在文件中增加对应的结构体和日志级别转换方法, 在 Cofiguration 结构体中添加 Zap 字段<br>
2.3.3 日志初始化方法：<br>
&emsp;&emsp;1- core 目录下新建 zap.go, 根目录下新建 utils 文件夹，并新建 directory.go，定义 PathExists 函数 检查存放日志的目录是否存在<br>
&emsp;&emsp;2- logger = zap.New(zapcore.Newtee(cores...))，
通过将多个 zapcore.Core 对象传递给 zapcore.NewTee 函数，创建了一个新的 zapcore.Core 对象。然后，将该 zapcore.Core 对象传递给 zap.New 函数，以创建一个新的 zap.Logger 对象。<br>
&emsp;&emsp;3- 这个 zap.Logger 对象可以用于实际的日志记录过程，根据配置的核心对象同时将日志条目发送给多个输出目标，实现多路日志记录的功能。<br>
2.3.4 core/internal 目录下新建 file_rotate_logs.go 文件，定义按时间分割日志的方法<br>
2.3.5 在 global.go 中添加全局变量 EWA_LOG \*zap.Logger<br>

## 2.4 封装 server 启动模块 "github.com/fvbock/endless"

2.4.1 core 目录下新建 server.go，定义服务启动的方法<br>
&emsp;&emsp;在 Cofiguration 结构体中添加 Zap 字段<br>
&emsp;&emsp;global 中挂载全局变量<br>

## 2.5 封装 gorm 连接 mysql 的方法 "gorm.io/gorm" "gorm.io/driver/mysql" "gorm.io/gorm/logger"

2.5.1 config 目录下新建 gorm_mysql.go，定义结构体及实现的方法, 然后在 Configuration 结构体中添加对应的字段<br>
2.5.2 根目录下新建 initialize 文件夹，这里封装一些数据库初始化操作<br>
&emsp;&emsp;initialize 目录下新建 gorm_mysql.go，定义 db 指针的初始化方法<br>
2.5.3 initialize 中新建 internal 文件夹，并新建 gorm_config.go<br>
&emsp;&emsp;我们把 gorm 配置抽离并封装成单独的方法，放在这个文件中因为后面其他的数据库连接也会使用到<br>
2.5.4 initialize/internal 目录下新建 logger.go，定义 gorm 日志记录器的写入器 \*writer<br>
2.5.5 初始化数据库，initialize 目录下新建 gorm.go 文件，设置启动数据库的类型<br>
2.5.6 main.go 中启动数据库<br>

## 2.6 封装 gorm 连接 postgresql 的方法 "gorm.io/driver/postgres"

2.6.1 用结构体对配置进行解析，然后将其挂载到全局变量上。config 目录下新建 gorm_pgsql.go，定义结构体及实现的方法，然后再 Cofiguration 添加对应的结构体和方法<br>
2.6.2 initialize 目录下新建 gorm_pgsql.go 文件，设置一下程序启动时需要的操作集合<br>
2.6.3 然后在 initialize/gorm.go 中增加使用 PostgreSQL 时的方法：<br>

## 2.7 路由封装

2.7.1 将我们 core/server.go 中启动服务的方法改造一下，在这里面做路由的初始化<br>
2.7.2 initialize 下新建 router.go 文件，实现路由初始化方法：<br>
2.7.3 在根目录下新建一个 router 文件夹 来管理所有的路由，并且每一级分组路由都统一用一个 enter.go 来管理<br>
&emsp;&emsp;在 enter.go 中，定义了 RouterGroup  的结构体，该结构体包含了一个名为  System  的字段，该字段的类型是  system.RouterGroup。同时，它还声明了一个名为  RouterGroupApp  的全局变量，并将其初始化为  new(RouterGroup)，这意味着  RouterGroupApp  是一个指向  RouterGroup  类型的指针，并且它的值为 nil。<br>
&emsp;&emsp;通常情况下，结构体是一种自定义类型，用于组合相关的字段，并将其作为单个实体来处理。在这个例子中，RouterGroup  结构体被用于定义路由分组信息，它的  System  字段会包含与路由分组相关的属性和方法。<br>
&emsp;&emsp;RouterGroupApp 变量是一个全局变量，因此它可以被其他代码文件访问。该变量的用途可能是将 RouterGroup 结构体或者 System 字段在项目的其他地方进行使用。<br>
&emsp;&emsp;通过创建这个结构体和全局变量，可以在应用级别上为同一类路由设置公共属性和方法，并便于在其他组件中复用,更好地对路由进行管理和组织在一个项目中。
例如，在一个 Web 应用程序中，通常会有许多不同的路由需要被注册、管理和处理。将这些路由按照业务逻辑或功能特点进行分组，可以让代码更加清晰易懂，同时也方便进行统一的权限控制、请求过滤等操作。使用结构体类型和全局变量来管理路由可以帮助开发者更好地组织代码，减少重复代码的出现，提高可维护性和可扩展性。
2.7.4 router 目录下新建 system 文件夹，并新建 enter.go 文件，里面定义了 system 分组路由，主要负责一些系统层面上的路由，比如登录、注册、权限管理等。<br>
2.7.5 system 目录下新建 sys_base.go 文件，里面定义 BaseRouter 结构体，BaseRouter 则主要负责最基础的登录、注册等路由。<br>
&emsp;&emsp;这里面则是定义了一个 BaseRouter 结构体和它的一个 InitBaseRouter 方法。该方法接收一个 gin 的 RouterGroup 类型参数，创建一个名为 "base" 的路由组，并为其添加一个 POST 请求路由 /login。当该路由被请求时，会返回一个 JSON 格式的字符串 "ok"。
该方法返回值类型为 gin.IRoutes 接口，因此实际上返回的是创建的 baseRouter 对象，可以在其他地方使用该对象以继续往该路由组中添加更多的路由。

## 2.8 自定义校验器 "github.com/go-playground/validator/v10" "regexp"

2.8.1 Gin 自带验证器返回的错误信息格式不太友好，本篇将进行调整，实现自定义错误信息，并规范接口返回的数据格式，分别为每种类型的错误定义错误码，前端可以根据对应的错误码实现后续不同的逻辑操作，结构体后面的 tag：\`binding:required`，这个 required 就是验证规则<br>
2.8.2 utils 目录下新建 validator.go 文件，用来存放所有跟校验相关的方法<br>
&emsp;&emsp;GetErrorMsg  函数来获取错误信息。它接收两个参数：
request  参数是被验证的请求结构体。
err  是用 Go 自带的验证器库  validator  验证参数时返回的错误。
函数会根据不同情况返回不同的错误信息。
如果传入的  err  参数属于 Go 自带的验证器库  validator  的  ValidationErrors  类型，即参数出现验证错误：
程序会判断请求结构体是否实现了  Validator  接口。
如果  request  实现了  Validator  接口，则可以自定义错误信息。这里的实现方式是：在  ValidatorMessages  中使用\<FieldName>.\<Tag>  作为 key，值为对应的错误信息。例如："name.required": "name 不能为空"  这个键值对就对应了  name  字段的  required  验证失败时输出的错误信息。
如果没有实现  Validator  接口，则直接返回默认的错误信息。
最后如果参数没有验证出现错误，则返回参数错误的提示信息  "Parameter error"<br>
2.8.3 根目录下新建 model 文件夹，该目录再新建 system/sys_user.go：
2.8.4 接着在 rouetr/system/sys_base.go 中新建一个 register 接口：
2.8.5 自定义验证规则 <br>
&emsp;&emsp;有一些验证规则在 Gin 框架中是没有的，这个时候我们就需要自定义验证器，验证规则将统一存放在 utils/validator.go 中，新增一个校验手机号的校验器。
2.8.6 initialize 目录下新建 other.go 文件，用来初始化一些其他的校验方法
2.8.7 在 main.go 中设置 自定义校验器初始化
2.8.7 这就可以再 model 里面的使用自定义的校验规则了，比如说对 mobile 字段的值使用 mobile 校验校验规则: <br>
&emsp;&emsp;tag 添加校验规则：binding:"required,mobile"`<br>
&emsp;&emsp;校验器添加返回信息："mobile.mobile": "手机号码格式不正确"<br>

## 2.9 引入 Redis "github.com/go-redis/redis/v8"

2.9.1 config.yaml 中增加 redis 配置：<br>
2.9.2 config 目录下新建 redis.go 文件, 定义 Redis 结构体，并在 Configuration 中引入<br>
2.9.3 initialize 目录下新建 redis.go，定义 redis 初始化方法<br>
2.9.4 然后在 core/server.go 初始化 redis 服务<br>
