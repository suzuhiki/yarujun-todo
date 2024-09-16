package main

import (
	"yarujun/app/cmd"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
// @description このswaggerはyarujunのAPIを定義しています。
func main() {
	// start api server
	r := cmd.GetRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
