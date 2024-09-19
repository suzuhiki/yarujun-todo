package controller

import (
	"net/http"
	"yarujun/app/model"
	"yarujun/app/types"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// @Summary アカウント作成
// @Tag 認証
// @Accept  json
// @Produce  json
// @Param   body	  body    types.CreateAccountRequest     true      "body param"
// @Success 200 {object} types.CreateAccountResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /create_account [post]
func CreateAccount(c *gin.Context) {
	var json types.CreateAccountRequest
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
// @Success 200 {object} types.SuccessResponse{data=[]types.TaskEntity}
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/tasks [get]
func ShowAllTask(c *gin.Context) {
	datas := model.GetAllTask()
	c.JSON(200, datas)
}

// @Summary タスクを作成する
// @Tag タスク
// @Accept  json
// @Security    BearerAuth
// @Param   user_id     query    string     true        "user_id"
// @Success 200 {object} types.CreateTaskRequest
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/tasks [post]
func CreateTask(c *gin.Context) {
	var json types.CreateTaskRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id := c.Query("user_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	err := model.CreateTask(user_id, json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// @Summary 現在のユーザーidを返す
// @Tag ユーザー
// @Produce  json
// @Security    BearerAuth
// @Success 200 {object} types.GetUserIdResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/current_user [get]
func GetCurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user_id := claims[jwt.IdentityKey].(string)
	var result types.GetUserIdResponse
	result.UserId = user_id
	c.JSON(200, result)
}

// @Summary hello worldを返す
// @Tag テスト
// @Produce  json
// @Success 200 {string} string "Hello, World!!!!!!!!"
// @Failure 400 {object} types.ErrorResponse
// @Router /test [get]
func Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!!!!!!!")
}

// @Summary ログイン
// @Tag 認証
// @Accept  json
// @Produce  json
// @Param   body	  body    types.LoginRequest     true      "body param"
// @Success 200 {object} types.LoginResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /login [post]
func login() {
}

// @Summary 認証情報の更新
// @Tag 認証
// @Produce  json
// @Security    BearerAuth
// @Success 200 {object} types.LoginResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/refresh_token [get]
func refresh_token() {
}
