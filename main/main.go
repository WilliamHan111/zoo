package main

import (
	"log"
	"net/http"
	"time"

	"github.com/WilliamHan111/zoo/consts"
	"github.com/WilliamHan111/zoo/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// 运行模式
var runMode string

func main() {
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
	err := g.Run(":8848")
	if err != nil {
		log.Panicf("gin engine run error: %s", err)
	}
}
