package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/go-config"
	"github.com/goworkeryyt/go-config/env"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/srun"
	"github.com/goworkeryyt/go-core/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	// 获取程序运行环境，默认会读取 resources/active.yaml 文件中配置的运行环境
	global.ENV = env.Active()

	// 获取全局配置,默认根据运行环境加载对应配置文件
	global.CONFIG = goconfig.GlobalConfig()

	// 初始化zap日志
	global.LOG = zap.Zap()

	// 初始化数据库连接
	global.DB = db.Gorm()

	// 获取配置文件原始内容,这样方便在程序中全局拿到自己定义的配置子项
	global.VP = global.CONFIG.Viper

	// 启动 http 服务
	r := gin.Default()
	// 健康监测
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	// 启动服务
	srun.RunHttpServer(r)
}
