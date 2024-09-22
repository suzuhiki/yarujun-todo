package controller

import (
	"net/http"
	"strconv"
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
// @Param   user_id     query    string     true        "user_id"
// @Param   sort     query    string     false        "deadline or waitlist_num"
// @Success 200 {object} types.SuccessResponse{data=[]types.ShowTaskResponse}
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/tasks [get]
func ShowAllTask(c *gin.Context) {
	user_id := c.Query("user_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	sort := c.Query("sort")
	if sort != "deadline" && sort != "waitlist_num" {
		sort = "deadline"
	}

	datas, err := model.GetAllTask(user_id, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, datas)
}

// @Summary タスクを作成する
// @Tag タスク
// @Accept  json
// @Security    BearerAuth
// @Param   user_id     query    string     true        "user_id"
// @Param   body	  body    types.CreateTaskRequest     true      "body param"
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

// @Summary タスクを完了にする
// @Tag タスク
// @Produce  json
// @Security    BearerAuth
// @Param   user_id     query    string     true        "user_id"
// @Param   task_id     query    string     true        "task_id"
// @Param   status     query    bool     true        "status"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/tasks/status [put]
func PutDoneTask(c *gin.Context) {
	user_id := c.Query("user_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	task_id := c.Query("task_id")
	if task_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_id is required"})
		return
	}
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}
	bool_status, err := strconv.ParseBool(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = model.UpdateDoneTask(user_id, task_id, bool_status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// @Summary タスクを削除する
// @Tag タスク
// @Produce  json
// @Security    BearerAuth
// @Param   user_id     query    string     true        "user_id"
// @Param   task_id     query    string     true        "task_id"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /auth/tasks [delete]
func DeleteTask(c *gin.Context) {
	user_id := c.Query("user_id")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	task_id := c.Query("task_id")
	if task_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_id is required"})
		return
	}

	err := model.DeleteTask(user_id, task_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
