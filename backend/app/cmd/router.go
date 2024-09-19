package cmd

import (
	"log"
	"yarujun/app/controller"

	"github.com/gin-gonic/gin"
)

// ginのルーターを設定
func GetRouter() *gin.Engine {
	jwtMiddleware, err := controller.NewJwtMiddleware()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	// ログイン
	v1.POST("/login", jwtMiddleware.LoginHandler)
	v1.POST("/create_account", controller.CreateAccount)
	v1.GET("/test", controller.Test)

	// 認証済みエンドポイント
	auth := v1.Group("/auth").Use(jwtMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	auth.GET("/tasks", controller.ShowAllTask)
	auth.GET("/current_user", controller.GetCurrentUser)

	return r
}
