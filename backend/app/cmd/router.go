package cmd

import (
	"log"
	"yarujun/app/controller"
	_ "yarujun/app/model"
	_ "yarujun/app/responses"

	"github.com/gin-gonic/gin"
)

// ginのルーターを設定
func GetRouter() *gin.Engine {
	jwtMiddleware, err := NewJwtMiddleware()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	// ログイン
	v1.POST("/login", jwtMiddleware.LoginHandler)
	v1.GET("/test", controller.Test)

	// 認証済みエンドポイント
	auth := v1.Group("/auth").Use(jwtMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	auth.GET("/tasks", controller.ShowAllTask)

	return r
}
