package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error()}
}

type CustomError struct {
	Msg string
	Err error
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}
