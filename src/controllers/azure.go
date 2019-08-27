package controllers

import (
	"../config"
	"../httputil"
	"../models"
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var interactor utils.AzureInteractor
var client *services.AzureClient

type AzureController struct{}

func NewAzureController() *AzureController {

	client = services.NewAzureClient(
		viper.GetString(config.AzureToken),
		viper.GetString(config.AzureOrganization),
	)

	interactor = utils.AzureInteractor{
		OrganizationName: viper.GetString(config.AzureOrganization),
		AuthToken:        viper.GetString(config.AzureToken),
	}

	return &AzureController{}
}

func (AzureController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("getProjects", getProjectsList)
	router.GET("getProjectTeams/:projectId", getProjectTeams)

	router.GET("getTeamMembers/:projectId/:teamId", getTeamMembers)
	router.GET("getTeamIterations/:projectId/:teamId", getTeamIterations)
	router.GET("getMemberWorkItems/:projectId/:teamId", getMemberWorkItems)
}

// @Summary Список проектов
// @Description Возвращает список проектов организации
// @Tags Azure
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getProjects [get]
func getProjectsList(ctx *gin.Context) {
	projects, _, err := client.Projects.Projects(
		&services.ProjectsParams{ApiVersion: "5.1"},
	)
	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureProjectsList(projects))
}

// @Summary Список команд
// @Description Возвращает список команд проекта
// @Tags Azure
// @Produce json
// @Param projectId path string true "Project Id"
// @Success 200 {array} models.Team
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getProjectTeams/{projectId} [get]
func getProjectTeams(ctx *gin.Context) {
	projectId := ctx.Param("projectId")

	if projectId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId must be provided")
		return
	}

	projectTeams, _, err := client.Projects.ProjectTeams(
		projectId,
		&services.ProjectsParams{ApiVersion: "5.0"},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureTeamsList(projectTeams))
}

// @Summary Список участников команды
// @Description Возвращает список участников команды проекта
// @Tags Azure
// @Produce json
// @Param projectId path string true "Project Id"
// @Param teamId path string true "Team Id"
// @Success 200 {array} models.Member
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getTeamMembers/{projectId}/{teamId} [get]
func getTeamMembers(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	teamId := ctx.Param("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	teamMembers, _, err := client.Teams.TeamMembers(
		projectId,
		teamId,
		&services.TeamsParams{ApiVersion: "5.1"},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureMembersList(teamMembers))
}

// @Summary Список спринтов команды
// @Description Возвращает список спринтов команды
// @Tags Azure
// @Produce json
// @Param projectId path string true "Project Id"
// @Param teamId path string true "Team Id"
// @Success 200 {array} models.Iteration
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getTeamIterations/{projectId}/{teamId} [get]
func getTeamIterations(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	teamId := ctx.Param("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	teamIterations, _, err := client.Teams.TeamIterations(
		projectId,
		teamId,
		&services.TeamsParams{ApiVersion: "5.1"},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureIterations(teamIterations))
}

// @Summary Задачи определенного участника команды
// @Tags Azure
// @Produce json
// @Param projectId path string true "Project Id"
// @Param teamId path string true "Team Id"
// @Param userEmail query string true "User Email"
// @Param iteration query string true "Iteration Name"
// @Success 200 {object} azure.WorkItemsResponse
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getMemberWorkItems/{projectId}/{teamId} [get]
func getMemberWorkItems(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	teamId := ctx.Param("teamId")
	userEmail := ctx.Query("userEmail")
	iteration := ctx.Query("iteration")

	if projectId == "" || teamId == "" || userEmail == "" || iteration == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId, teamId, userId, iteration must be provided")
		return
	}

	workItemsWiql, err := interactor.GetWorkItemsByWiql(projectId, teamId, userEmail, iteration)
	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	// Request detailed works
	var newList = make([]int, len(workItemsWiql.WorkItems))
	for index, element := range workItemsWiql.WorkItems {
		newList[index] = element.ID
	}

	workItems, err := interactor.GetWorkItemsDescription(projectId, newList)
	if err != nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, workItems)
}
