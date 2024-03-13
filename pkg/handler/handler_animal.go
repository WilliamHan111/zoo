package handler

import (
	"net/http"

	"github.com/WilliamHan111/zoo/pkg/baidu/animal"
	"github.com/gin-gonic/gin"
)

type AnimalManager struct {
}

func (ctl *AnimalManager) CheckAnimal(c *gin.Context) {
	//动物识别
	result := animal.CheckAnimal()
	//接口返回
	c.JSON(http.StatusOK, result)
}
