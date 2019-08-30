package controllers

import (
	"Databriz-Meetings-API-Go/db"
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebController struct{}

type isDataActualResponse struct {
	ShouldUpdate bool `json:"should_update"`
	Revision     int  `json:"revision"`
}

func NewWebController() *WebController {
	return &WebController{}
}

func (c *WebController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("data/revision/isActual", c.isDataActual)
	//router.GET("data/get", c.getActualData)
}

// @Summary Проверка актуальности данных на фронте
// @Description Возвращает актуальную версию данных и флаг необходимости обновления
// @Tags Web
// @Produce json
// @Param revision query int true "Revision Id"
// @Success 200 {object} isDataActualResponse
// @Failure 400 {object} httputil.HTTPError "When user has provided wrong query param"
// @Router /v1/web/data/revision/isActual [get]
func (WebController) isDataActual(ctx *gin.Context) {
	revision := ctx.Query("revision")

	if revision == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "revision param must be provided")
		return
	}

	frontendRevision, err := utils.StringToInt(revision)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, "revision param must be int")
		return
	}
	storage := db.GetMemoryStorage()

	ctx.JSON(http.StatusOK, isDataActualResponse{
		ShouldUpdate: storage.ShouldUpdate(frontendRevision),
		Revision:     storage.GetDataRevision(),
	})
}

//func (WebController) getActualData(ctx *gin.Context) {
//
//}
