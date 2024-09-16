package cmd

import (
	"net/http"
	"yarujun/app/controller"
	_ "yarujun/app/model"
	_ "yarujun/app/responses"

	"github.com/gin-gonic/gin"
)

// GetTodos ...
// @Summary Todo一覧を配列で返す
// @Tags Todo
// @Produce  json
// @Success 200 {object} responses.SuccessResponse{data=[]model.TaskEntity}
// @Failure 400 {object} responses.ErrorResponse
// @Router /todos [get]
func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.ShowAllTask)
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!!!!!!!!")
	})
	return r
}
