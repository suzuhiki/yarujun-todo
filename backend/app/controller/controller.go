package controller

import (
	"net/http"
	"yarujun/app/model"

	"github.com/gin-gonic/gin"
)

// @Summary Todo一覧を配列で返す
// @Tag 一覧画面
// @Produce  json
// @Success 200 {object} responses.SuccessResponse{data=[]model.TaskEntity}
// @Failure 400 {object} responses.ErrorResponse
// @Router /tasks [get]
func ShowAllTask(c *gin.Context) {
	datas := model.GetAll()
	c.JSON(200, datas)
}

// @Summary hello worldを返す
// @Tag テスト
// @Produce  json
// @Success 200 {string} string "Hello, World!!!!!!!!"
// @Failure 400 {object} responses.ErrorResponse
// @Router /test [get]
func Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!!!!!!!")
}

// @Summary ログイン
// @Tag 認証
// @Accept  json
// @Produce  json
// @Param   body	  body    loginRequest     true      "body param"
// @Success 200 {object} loginResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /login [post]
func login() {
}

type loginRequest struct {
	Email    string `form:"email" json:"email" binding:"required" example:"test@example.com"`
	Password string `form:"password" json:"password" binding:"required" example:"test"`
}

type loginResponse struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expier" example:"2024-09-20T03:12:53+09:00"`
	Token  string `json:"token"`
}
