package types

type TaskEntity struct {
	Title        string
	Memo         string
	Deadline     string
	Waitlist_num string
}

type LoginRequest struct {
	Name     string `form:"name" json:"name" binding:"required" example:"taro"`
	Password string `form:"password" json:"password" binding:"required" example:"tarodesu"`
}

type LoginResponse struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expier" example:"2024-09-20T03:12:53+09:00"`
	Token  string `json:"token"`
}

type CreateAccountRequest struct {
	Name     string `form:"name" json:"name" binding:"required" example:"taro"`
	Password string `form:"password" json:"password" binding:"required" example:"tarodesu"`
}

type CreateAccountResponse struct {
	Code int    `json:"code" example:"200"`
	Name string `form:"name" json:"name" example:"test"`
}

type CreateTaskRequest struct {
	Title        string `form:"title" json:"title" binding:"required" example:"やること"`
	Memo         string `form:"memo" json:"memo" example:"概要"`
	Deadline     string `form:"deadline" json:"deadline" example:"2024-09-20T03:12:53+09:00"`
	Waitlist_num int    `form:"waitlist_num" json:"waitlist_num" example:1`
}

type GetUserIdResponse struct {
	UserId string `json:"user_id" example:"1"`
}
