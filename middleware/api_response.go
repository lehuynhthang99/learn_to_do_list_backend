package middleware

import "net/http"

const (
	Success = iota
	ErrorUnknown
	InvalidInput
	DatabaseError
	NotExisted
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (customContext *CustomContext) Success(data interface{}) error {
	return customContext.JSON(http.StatusOK, APIResponse{
		Code: Success,
		Data: data,
	})
}

func (customContext *CustomContext) SuccessWithMessage(message string, data interface{}) error {
	return customContext.JSON(http.StatusOK, APIResponse{
		Code:    Success,
		Message: message,
		Data:    data,
	})
}

func (customContext *CustomContext) ErrorUnknown(statusCode int, message string) error {
	return customContext.JSON(statusCode, APIResponse{
		Code:    ErrorUnknown,
		Message: message,
	})
}

func (customContext *CustomContext) ErrorWithCustomCode(statusCode int, customCode int, message string) error {
	return customContext.JSON(statusCode, APIResponse{
		Code:    customCode,
		Message: message,
	})
}
