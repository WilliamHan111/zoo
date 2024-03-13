package main

import (
	"log"
	"net/http"
	"time"

	"github.com/WilliamHan111/zoo/consts"
	"github.com/WilliamHan111/zoo/pkg/conf"
	"github.com/WilliamHan111/zoo/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	runMode    string
	configPath string
)

func main() {
	//配置文件
	configPath = "../conf/basic.yaml"

	//配置初始化
	err := conf.InitConfigs(configPath)
	if err != nil {
		panic(err)
	}

	//创建gin
	g := gin.Default()
	pprof.Register(g)
	if runMode != consts.ReleaseMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
			AllowHeaders:     []string{"*", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	))
	g.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	g.Any("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//加载路由
	route.LoadRoute(g)

	//启动http服务
	err = g.Run(conf.AllConfig.Web.HttpPort)
	if err != nil {
		log.Panicf("gin engine run error: %s", err)
	}
}
