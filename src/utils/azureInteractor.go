package utils

import (
	"Databriz-Meetings-API-Go/src/models/azure"
	"encoding/json"
	"errors"
	"strings"
)

const (
	organizationPlaceholder string = "{organization}"
	projectIdPlaceholder    string = "{projectId}"
	teamIdPlaceholder       string = "{teamId}"
	userEmailPlaceholder    string = "{email}"
	iterationPlaceholder    string = "{iteration}"

	azureProjectsUrl      = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects?api-version=5.1"
	azureProjectTeamsUrl  = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams?api-version=5.0"
	azureTeamMembers      = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams/" + teamIdPlaceholder + "/members?api-version=5.1"
	azureTeamIterations   = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/" + teamIdPlaceholder + "/_apis/work/teamsettings/iterations?api-version=5.1"
	azureMemberTasks      = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/" + teamIdPlaceholder + "/_apis/wit/wiql?api-version=5.1"
	azureMemberTasksBatch = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/_apis/wit/workitemsbatch?api-version=5.1"

	memberTasksWiql = "select [System.Id] from WorkItems where [System.TeamProject] = @project and [System.AssignedTo] = '" + userEmailPlaceholder + "' and [System.IterationPath] = '" + iterationPlaceholder + "' order by [System.ChangedDate] desc"
)

type AzureInteractor struct {
	OrganizationName string
	AuthToken        string
}

type WorkItemsBatchRequestBody struct {
	Ids    []int    `json:"ids"`
	Fields []string `json:"fields"`
}

func (i *AzureInteractor) GetWorkItemsDescription(projectId string, workItemsIds []int) (response *azure.WorkItemsResponse, err error) {
	requestBody, err := json.Marshal(WorkItemsBatchRequestBody{
		Ids: workItemsIds,
		Fields: []string{
			"System.Title",
			"System.WorkItemType",
			"System.Reason",
		},
	})
	if err != nil {
		return nil, err
	}

	// Create url for calling azure api
	url := strings.Replace(azureMemberTasksBatch, organizationPlaceholder, i.OrganizationName, -1)
	url = strings.Replace(url, projectIdPlaceholder, projectId, -1)

	// Call Azure API
	PostToAzure(url, i.AuthToken, requestBody, &response)

	if response == nil {
		return nil, errors.New("error while receiving data from azure")
	}

	return
}
