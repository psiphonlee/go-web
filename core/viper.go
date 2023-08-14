package core

import (
	"flag"
	"fmt"
	"gomap/core/internal"
	"gomap/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*
InitializeViper()函数可以以多种方式选择配置文件启动服务
1. path 传参
2. 命令行传参给flag
3. 初始化的时候通过gin.SetMode()设置
优先级: 命令行 > 环境变量 > 默认值
*/
func InitViper(path ...string) *viper.Viper {
	// config 变量用来接收命令行的参数，用于指定读取哪个配置文件
	var config string

	if len(path) == 0 {
		// 如果没有InitViper指定配置文件，就检查命令行参数
		// 定义命令行flag参数，格式：flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
		// 命令行指定配置文件启动命令：go run main.go -c config.yaml , -c后面的就是指定的配置文件名
		flag.StringVar(&config, "c", "", "choose config file.")

		// 定义好命令行flag参数后，需要通过调用flag.Parse()解析命令行参数
		flag.Parse()

		// 如果没有指定命令行参数
		if config == "" {
			/*
				1. 检测是否通过命令行设置环境变量来指定yaml文件，注意windows和linux下格式区别：
				   windows: $env:GVA_CONFIG="config.yaml"; go run main.go
				   linux: GVA_CONFIG=config.yaml go run main.go
				2. 环境变量没有指定的话就检测是否通过gin.SetMode("debug" or "test" or "release") 来指定
			*/
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用%s模式的gin.debug环境,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用%s模式的gin.release环境,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用%s模式的gin.test环境,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				// internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量指定的yaml文件启动服务：%s\n", internal.ConfigEnv, config)
			}

		} else {
			// 命令行参数不为空
			fmt.Printf("您正在使用命令行的-c 传参指定的yaml文件启动服务：%s\n", config)
		}
	} else {
		// 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func InitViper(path)函数传参指定yaml文件启动服务：%s\n", config)
	}

	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("没有找到指定的yaml文件: %s \n", err))
	}

	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = vip.Unmarshal(&global.EWA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = vip.Unmarshal(&global.EWA_CONFIG); err != nil {
		fmt.Println(err)
	}
	fmt.Println("====1-viper====: viper init config success")
	return vip
}
