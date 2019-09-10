package httputil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewError(ctx *gin.Context, status int, err string) {
	er := HTTPError{
		StatusCode: status,
		Message:    err,
	}
	ctx.JSON(status, er)
}

func NewInternalAzureError(ctx *gin.Context, err error) {
	errStr := err.Error()
	fmt.Println(fmt.Sprintf("Error occurred = %s", errStr))
	NewError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error while requesting data from Azure. %s", errStr))
}

type HTTPError struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"status bad request"`
}
