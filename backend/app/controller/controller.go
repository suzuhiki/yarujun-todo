package controller

import (
	"net/http"
	"yarujun/app/model"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {

}

// @Summary Todo一覧を配列で返す
// @Tag 一覧画面
// @Produce  json
// @Security    BearerAuth
// @Success 200 {object} responses.SuccessResponse{data=[]model.TaskEntity}
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/tasks [get]
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

// @Summary 認証情報の更新
// @Tag 認証
// @Produce  json
// @Security    BearerAuth
// @Success 200 {object} loginResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/refresh_token [get]
func refresh_token() {
}

type loginRequest struct {
	Id       string `form:"id" json:"id" binding:"required" example:"testaro"`
	Password string `form:"password" json:"password" binding:"required" example:"test"`
}

type loginResponse struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expier" example:"2024-09-20T03:12:53+09:00"`
	Token  string `json:"token"`
}

type createAccountRequest struct {
	Id       string `form:"id" json:"id" binding:"required" example:"testaro"`
	Password string `form:"password" json:"password" binding:"required" example:"test"`
}
