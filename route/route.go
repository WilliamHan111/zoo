package route

import (
	"github.com/WilliamHan111/zoo/pkg/handler"
	"github.com/gin-gonic/gin"
)

// 路由
func LoadRoute(g *gin.Engine) {
	//dandler的控制层
	ctl := handler.NewController()
	//动物管理接口
	r := g.Group("/api/v1/animal")
	r.POST("/check_animal", ctl.AnimalManager.CheckAnimal)
	//人员管理接口

	//车辆管理接口
}
