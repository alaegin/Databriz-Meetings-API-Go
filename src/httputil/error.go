package httputil

import (
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

func NewInternalAzureError(ctx *gin.Context) {
	NewError(ctx, http.StatusInternalServerError, "Error while requesting data from Azure")
}

type HTTPError struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"status bad request"`
}
