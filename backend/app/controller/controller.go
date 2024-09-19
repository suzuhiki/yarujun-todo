package controller

import (
	"net/http"
	"yarujun/app/model"

	"github.com/gin-gonic/gin"
)

// @Summary アカウント作成
// @Tag 認証
// @Accept  json
// @Produce  json
// @Param   body	  body    createAccountRequest     true      "body param"
// @Success 200 {object} createAccountResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /create_account [post]
func CreateAccount(c *gin.Context) {
	var json createAccountRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := model.CreateAccount(json.Name, json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	Name     string `form:"name" json:"name" binding:"required" example:"testaro"`
	Password string `form:"password" json:"password" binding:"required" example:"test"`
}

type loginResponse struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expier" example:"2024-09-20T03:12:53+09:00"`
	Token  string `json:"token"`
}

type createAccountRequest struct {
	Name     string `form:"name" json:"name" binding:"required" example:"taro"`
	Password string `form:"password" json:"password" binding:"required" example:"tarodesu"`
}

type createAccountResponse struct {
	Code int    `json:"code" example:"200"`
	Name string `form:"name" json:"name" example:"test"`
}
