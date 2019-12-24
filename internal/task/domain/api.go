package domain

type Response struct {
	ResponseData interface{} `json:"response"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}
