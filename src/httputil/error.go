package httputil

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err string) {
	er := HTTPError{
		StatusCode: status,
		Message:    err,
	}
	ctx.JSON(status, er)
}

type HTTPError struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"status bad request"`
}
