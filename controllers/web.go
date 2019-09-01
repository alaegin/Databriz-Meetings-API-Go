package controllers

import (
	"Databriz-Meetings-API-Go/configs"
	"Databriz-Meetings-API-Go/db"
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/models"
	"Databriz-Meetings-API-Go/repository"
	"Databriz-Meetings-API-Go/services"
	"Databriz-Meetings-API-Go/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type WebController struct {
	client     *services.AzureClient
	memStorage *db.MemoryStorage
	userRepo   repository.UserRepository
}

type isDataActualResponse struct {
	ShouldUpdate bool  `json:"should_update"`
	Revision     int64 `json:"revision"`
}

type dataResponse struct {
	UserName      string             `json:"user_name"`
	UserEmail     string             `json:"user_email"`
	UserWorkItems *[]models.WorkItem `json:"user_work_items"`
}

func NewWebController() *WebController {
	client := services.NewAzureClient(
		viper.GetString(configs.AzureToken),
		viper.GetString(configs.AzureOrganization),
	)
	memoryStorage := db.GetMemoryStorage()
	userRepository := repository.NewUserRepository(db.GetDB())

	return &WebController{
		client:     client,
		memStorage: memoryStorage,
		userRepo:   userRepository,
	}
}

func (c *WebController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("data/revision/isActual", c.isDataActual)
	router.GET("data/get", c.getActualData)
}

// @Summary Проверка актуальности данных на фронте
// @Description Возвращает актуальную версию данных и флаг необходимости обновления
// @Tags Web
// @Produce json
// @Param revision query int true "Revision Id"
// @Success 200 {object} isDataActualResponse
// @Failure 400 {object} httputil.HTTPError "When user has provided wrong query param"
// @Router /v1/web/data/revision/isActual [get]
func (c *WebController) isDataActual(ctx *gin.Context) {
	revision := ctx.Query("revision")

	if revision == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "revision param must be provided")
		return
	}

	frontendRevision, err := utils.StringToInt64(revision)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, "revision param must be int")
		return
	}

	ctx.JSON(http.StatusOK, isDataActualResponse{
		ShouldUpdate: c.memStorage.ShouldUpdate(frontendRevision),
		Revision:     c.memStorage.GetDataRevision(),
	})
}

// @Summary Актуальные данные
// @Description Возвращает список работ выбранного пользователя
// @Tags Web
// @Produce json
// @Success 200 {object} dataResponse
// @Failure 400 {object} httputil.HTTPError "When nothing was selected from mobile app"
// @Router /v1/web/data/get [get]
func (c *WebController) getActualData(ctx *gin.Context) {
	request := c.memStorage.GetData()

	if request == nil {
		httputil.NewError(ctx, http.StatusBadRequest, "You must send data from mobile app before calling this method")
		return
	}

	workItems, err := c.client.WorkItems.MemberWorkItems(
		&services.MemberWorkItemsParams{
			ProjectId: request.ProjectId,
			TeamId:    request.TeamId,
			Iteration: request.IterationPath,
			UserEmail: request.UserEmail},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	userEntity := c.userRepo.GetByEmail(request.UserEmail)
	ctx.JSON(http.StatusOK, dataResponse{
		UserName:      userEntity.Name,
		UserEmail:     userEntity.Email,
		UserWorkItems: models.FromAzureWorkItems(workItems),
	})
}
