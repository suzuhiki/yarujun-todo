package types

// SuccessResponse ...
type SuccessResponse struct {
	Data interface{} `json: "data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
