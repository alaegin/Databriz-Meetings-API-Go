package controllers

import (
	"Databriz-Meetings-API-Go/db"
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MobileController struct {
	memStorage *db.MemoryStorage
}

func NewMobileController() *MobileController {
	memoryStorage := db.GetMemoryStorage()

	return &MobileController{
		memStorage: memoryStorage,
	}
}

func (c *MobileController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("control/show", c.showMemberWorkItems)
}

// @Summary Переключение фронта
// @Description Переключает отображающегося ппользователя на фронте
// @Tags Mobile
// @Produce json
// @Success 200
// @Failure 400 {object} httputil.HTTPError "When user has not provided correct request body"
// @Router /v1/mobile/control/show [post]
func (c *MobileController) showMemberWorkItems(ctx *gin.Context) {
	var requestBody models.ShowRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		httputil.NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	c.memStorage.StoreData(requestBody)

	ctx.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	})
}
