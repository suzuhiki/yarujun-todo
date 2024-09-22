package types

type ShowTaskResponse struct {
	Id           string
	Title        string
	Deadline     string
	Waitlist_num string
	Done         bool
}

type LoginRequest struct {
	Name     string `form:"name" json:"name" binding:"required" example:"admin"`
	Password string `form:"password" json:"password" binding:"required" example:"admin"`
}

type LoginResponse struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expier" example:"2024-09-20"`
	Token  string `json:"token"`
}

type CreateAccountRequest struct {
	Name     string `form:"name" json:"name" binding:"required" example:"admin"`
	Password string `form:"password" json:"password" binding:"required" example:"admin"`
}

type CreateAccountResponse struct {
	Code int    `json:"code" example:"200"`
	Name string `form:"name" json:"name" example:"test"`
}

type CreateTaskRequest struct {
	Title        string `form:"title" json:"title" binding:"required" example:"やること"`
	Deadline     string `form:"deadline" json:"deadline" binding:"required" example:"2024-09-20"`
	Waitlist_num int    `form:"waitlist_num" json:"waitlist_num" example:1`
}

type GetUserIdResponse struct {
	UserId string `json:"user_id" example:"1"`
}
