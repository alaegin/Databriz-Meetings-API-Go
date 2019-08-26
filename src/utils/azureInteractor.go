package utils

const (
	organizationPlaceholder string = "{organization}"
	projectIdPlaceholder    string = "{projectId}"
	teamIdPlaceholder       string = "{teamId}"
	userEmailPlaceholder    string = "{email}"
	iterationPlaceholder    string = "{iteration}"

	azureProjectsUrl     = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects?api-version=5.1"
	azureProjectTeamsUrl = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams?api-version=5.0"
	azureTeamMembers     = "https://dev.azure.com/" + organizationPlaceholder + "/_apis/projects/" + projectIdPlaceholder + "/teams/" + teamIdPlaceholder + "/members?api-version=5.1"
	azureTeamIterations  = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/" + teamIdPlaceholder + "/_apis/work/teamsettings/iterations?api-version=5.1"
	azureMemberTasks     = "https://dev.azure.com/" + organizationPlaceholder + "/" + projectIdPlaceholder + "/" + teamIdPlaceholder + "/_apis/wit/wiql?api-version=5.1"

	memberTasksWiql = "select [System.Id] from WorkItems where [System.TeamProject] = @project and [System.AssignedTo] = '" + userEmailPlaceholder + "' and [System.IterationPath] = '" + iterationPlaceholder + "' order by [System.ChangedDate] desc"
)

/*type AzureInteractor struct {

}

type WorkItemsBatchRequestBody struct {
	Ids []int `json:"ids"`
}

func GetWorkItemsDescription(workItemsIds []int) {
	requestBody, err := json.Marshal(WorkItemsBatchRequestBody{Ids: workItemsIds})
	if err != nil {

	}

}*/
