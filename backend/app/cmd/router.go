package cmd

import (
	"yarujun/app/controller"
	_ "yarujun/app/model"
	_ "yarujun/app/responses"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.ShowAllTask)
	r.GET("/test", controller.Test)
	return r
}
