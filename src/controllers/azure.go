package controllers

import (
	"../config"
	"../httputil"
	"../models"
	"../models/azure"
	"../utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	organizationPlaceholder string = "{organization}"
	projectIdPlaceholder    string = "{projectId}"
	teamIdPlaceholder       string = "{teamId}"

	azureProjectsUrl     = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects?api-version=5.1"
	azureProjectTeamsUrl = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams?api-version=5.0"
	azureTeamMembers     = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams/" + teamIdPlaceholder + "/members?api-version=5.1"
	azureTeamIterations  = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/" + teamIdPlaceholder + "/_apis/work/teamsettings/iterations?api-version=5.1"
)

// TODO Remove interacting with azure from controller
var orgName string
var token string

var interactor utils.AzureInteractor

type AzureController struct{}

func NewAzureController() *AzureController {

	orgName = viper.GetString(config.AzureOrganization)
	token = viper.GetString(config.AzureToken)

	interactor = utils.AzureInteractor{
		OrganizationName: viper.GetString(config.AzureOrganization),
		AuthToken:        viper.GetString(config.AzureToken),
	}

	return &AzureController{}
}

// @Summary Список проектов
// @Description Возвращает список проектов организации
// @Tags Azure
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} httputil.HTTPError "When failed to receive data from Azure"
// @Router /api/v1/azure/getProjects [get]
func (c *AzureController) GetProjectsList(ctx *gin.Context) {
	// Create url for calling azure api
	url := strings.Replace(azureProjectsUrl, organizationPlaceholder, orgName, -1)

	// Call Azure API
	var projects *azure.ProjectsList
	utils.GetFromAzure(url, token, &projects)

	if projects == nil {
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
func (c *AzureController) GetProjectTeams(ctx *gin.Context) {
	projectId := ctx.Param("projectId")

	if projectId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId must be provided")
		return
	}

	// Create url for calling azure api
	url := strings.Replace(azureProjectTeamsUrl, organizationPlaceholder, orgName, -1)
	url = strings.Replace(url, projectIdPlaceholder, projectId, -1)

	// Call Azure API
	var teams *azure.TeamsList
	utils.GetFromAzure(url, token, &teams)

	if teams == nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureTeamsList(teams))
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
func (c *AzureController) GetTeamIterations(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	teamId := ctx.Param("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	// Create url for calling azure api
	url := strings.Replace(azureTeamIterations, organizationPlaceholder, orgName, -1)
	url = strings.Replace(url, projectIdPlaceholder, projectId, -1)
	url = strings.Replace(url, teamIdPlaceholder, teamId, -1)

	// Call Azure API
	var iterations *azure.IterationsList
	utils.GetFromAzure(url, token, &iterations)

	if iterations == nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.FromAzureIterations(iterations))
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
func (c *AzureController) GetTeamMembers(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	teamId := ctx.Param("teamId")

	if projectId == "" || teamId == "" {
		httputil.NewError(ctx, http.StatusBadRequest, "projectId and teamId must be provided")
		return
	}

	// Create url for calling azure api
	url := strings.Replace(azureTeamMembers, organizationPlaceholder, orgName, -1)
	url = strings.Replace(url, projectIdPlaceholder, projectId, -1)
	url = strings.Replace(url, teamIdPlaceholder, teamId, -1)

	// Call Azure API
	var members *azure.MembersList
	utils.GetFromAzure(url, token, &members)

	if members == nil {
		httputil.NewInternalAzureError(ctx)
		return
	}

	// Convert Azure response to own format
	result := models.FromAzureMembersList(members)

	ctx.JSON(http.StatusOK, result)
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
func (c *AzureController) GetMemberWorkItems(ctx *gin.Context) {
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
