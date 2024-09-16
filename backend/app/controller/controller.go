package controller

import (
	"yarujun/app/model"

	"github.com/gin-gonic/gin"
)

func ShowAllTask(c *gin.Context) {
	datas := model.GetAll()
	c.JSON(200, datas)
}
