package controllers

import (
	"Databriz-Meetings-API-Go/configs"
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/models"
	"Databriz-Meetings-API-Go/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var client *services.AzureClient

type AzureController struct{}

func NewAzureController() *AzureController {

	client = services.NewAzureClient(
		viper.GetString(configs.AzureToken),
		viper.GetString(configs.AzureOrganization),
	)

	return &AzureController{}
}

func (c *AzureController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("projects/list", c.getProjectsList)

	router.GET("teams/list", c.getProjectTeams)              // ?projectId
	router.GET("teams/members/list", c.getTeamMembers)       // ?projectId, teamId
	router.GET("teams/iterations/list", c.getTeamIterations) // ?projectId, teamId

	router.GET("members/:memberId/workItems/list", c.getMemberWorkItems) // ?projectId, teamId, iteration
}

// @Summary Список проектов
// @Description Возвращает список проектов организации
// @Tags Azure
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /v1/azure/projects/list [get]
func (AzureController) getProjectsList(ctx *gin.Context) {
	projects, _, err := client.Projects.Projects(
		&services.ProjectsParams{},
	)
	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureProjectsList(projects))
}

// @Summary Список команд
// @Description Возвращает список команд проекта
// @Tags Azure
// @Produce json
// @Param projectId query string true "Project Id"
// @Success 200 {array} models.Team
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /v1/azure/teams/list [get]
func (AzureController) getProjectTeams(ctx *gin.Context) {
	projectId := ctx.Query("projectId")

	if projectId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId must be provided")
		return
	}

	projectTeams, _, err := client.Projects.ProjectTeams(
		&services.ProjectTeamsParams{ProjectId: projectId},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureTeamsList(projectTeams))
}

// @Summary Список участников команды
// @Description Возвращает список участников команды проекта
// @Tags Azure
// @Produce json
// @Param projectId query string true "Project Id"
// @Param teamId query string true "Team Id"
// @Success 200 {array} models.Member
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /v1/azure/teams/members/list [get]
func (AzureController) getTeamMembers(ctx *gin.Context) {
	projectId := ctx.Query("projectId")
	teamId := ctx.Query("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	teamMembers, _, err := client.Teams.TeamMembers(
		&services.TeamMembersParams{ProjectId: projectId,
			TeamId: teamId},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureMembersList(teamMembers))
}

// @Summary Список спринтов команды
// @Description Возвращает список спринтов команды
// @Tags Azure
// @Produce json
// @Param projectId query string true "Project Id"
// @Param teamId query string true "Team Id"
// @Success 200 {array} models.Iteration
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /v1/azure/teams/iterations/list [get]
func (AzureController) getTeamIterations(ctx *gin.Context) {
	projectId := ctx.Query("projectId")
	teamId := ctx.Query("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	teamIterations, _, err := client.Teams.TeamIterations(
		&services.TeamIterationsParams{ProjectId: projectId,
			TeamId: teamId},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureIterations(teamIterations))
}

// @Summary Задачи определенного участника команды
// @Tags Azure
// @Produce json
// @Param memberId path string true "User Email"
// @Param projectId query string true "Project Id"
// @Param teamId query string true "Team Id"
// @Param iteration query string true "Iteration Name"
// @Success 200 {array} models.WorkItem
// @Failure 400 {object} httputil.HTTPError "When user has not provided projectId or teamId parameter"
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /v1/azure/members/{memberId}/workItems/list [get]
func (AzureController) getMemberWorkItems(ctx *gin.Context) {
	userEmail := ctx.Param("memberId")
	projectId := ctx.Query("projectId")
	teamId := ctx.Query("teamId")
	iteration := ctx.Query("iteration")

	if projectId == "" || teamId == "" || userEmail == "" || iteration == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId, teamId, userEmail, iteration must be provided")
		return
	}

	workItems, err := client.WorkItems.MemberWorkItems(
		&services.MemberWorkItemsParams{ProjectId: projectId,
			TeamId:    teamId,
			Iteration: iteration,
			UserEmail: userEmail},
	)

	if err != nil {
		httputil.NewInternalAzureError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureWorkItems(workItems))
}
