package dto

type ErrorBody struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}
