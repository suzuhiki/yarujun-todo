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

	// 認証
	r.POST("/login", jwtMiddleware.LoginHandler)
	r.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	// 認証済みのみ
	r.Use(jwtMiddleware.MiddlewareFunc()).GET("/", controller.ShowAllTask)

	r.GET("/test", controller.Test)
	return r
}
