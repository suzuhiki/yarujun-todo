package cmd

import (
	"log"
	"yarujun/app/controller"
	_ "yarujun/app/model"
	_ "yarujun/app/responses"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	jwtMiddleware, err := NewJwtMiddleware()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", controller.ShowAllTask)

	// 認証
	r.POST("/login", jwtMiddleware.LoginHandler)
	r.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	r.GET("/test", controller.Test)
	return r
}
