package main

import (
	"fisco-cert-app/controller"
	"fisco-cert-app/dao/fisco"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/dao/redis"
	_ "fisco-cert-app/docs"
	"fisco-cert-app/logger"
	"fisco-cert-app/pkg/snowflake"
	"fisco-cert-app/pkg/watermark"
	"fisco-cert-app/router"
	"fisco-cert-app/setting"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"os"
	// swagger embed files
	// gin-swagger middleware
)

/**
CLD web服务模板
*/

// @title           fisco-cert-app 社区
// @version         1.0
// @description     关于fisco-cert-app社区接口文档
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  1239989762@qq.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8084
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}

	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	configs, err := conf.ParseConfigFile("conf/fisco.toml")
	if err != nil {
		fmt.Printf("ParseConfigFile fisco-bcos failed, err: %v", err)
		return
	}

	if err = fisco.Init(&configs[0]); err != nil {
		fmt.Printf("init fisco failed, err:%v", err)
		return
	}
	defer fisco.Close()

	watermark.Init(setting.Conf.Tencent)

	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 程序退出记得关闭
	defer mysql.Close()
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	if err := controller.InitTrans(setting.Conf.Local); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	r := router.SetupRouter(setting.Conf.Mode)
	err = r.Run(fmt.Sprintf(":%v", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
