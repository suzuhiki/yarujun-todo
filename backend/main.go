package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"yarujun/app/cmd"
	_ "yarujun/docs"

	_ "github.com/lib/pq"
)

type EMPLOYEE struct {
	ID     string
	NUMBER string
}

// @title gin-swagger todos
// @version 1.0
// @license.name suzuhiki
// @description このswaggerはyarujunのAPIを定義しています。 JWTトークンの前に"Bearer"を追加してください。
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	// start api server
	r := cmd.GetRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
