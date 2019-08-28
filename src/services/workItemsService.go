package services

import (
	"Databriz-Meetings-API-Go/src/httputil"
	"Databriz-Meetings-API-Go/src/models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

const (
	workItemsListQuery = "select [System.Id] from WorkItems " +
		"where [System.TeamProject] = @project " +
		"and [System.AssignedTo] = '%s' " +
		"and [System.IterationPath] = '%s' " +
		"order by [System.ChangedDate] desc"
)

var fields = []string{
	"System.Title",
	"System.WorkItemType",
	"System.State",
	"System.Reason",
	"System.CreatedDate",
	"Microsoft.VSTS.Scheduling.OriginalEstimate",
	"Microsoft.VSTS.Scheduling.CompletedWork",
	"Microsoft.VSTS.Common.Priority",
}

type WorkItemsService struct {
	sling *sling.Sling
}

type WorkItemsParams struct {
	ApiVersion string `url:"api-version,omitempty"`
}

func newWorkItemsService(sling *sling.Sling, organization string) *WorkItemsService {
	return &WorkItemsService{
		sling: sling.Path(fmt.Sprintf("%s/", organization)),
	}
}

func (s *WorkItemsService) MemberWorkItems(projectId, teamId, iteration, userEmail string, params *WorkItemsParams) (*azure.WorkItemsList, error) {
	workItemsList, _, err := s.workItemsList(projectId, teamId, iteration, userEmail, params)
	if err != nil {
		return nil, err
	}

	workItems, _, err := s.workItemsDescription(projectId, workItemIds(workItemsList), params)
	if err != nil {
		return nil, err
	}

	return workItems, nil
}

func workItemIds(workItemsList *azure.ShortWorkItemsList) []int {
	var ids = make([]int, len(workItemsList.ShortWorkItems))
	for index, element := range workItemsList.ShortWorkItems {
		ids[index] = element.ID
	}
	return ids
}

type workItemsListRequestBody struct {
	Query string `json:"query,omitempty"`
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/wiql/query%20by%20wiql?view=azure-devops-rest-5.1
func (s *WorkItemsService) workItemsList(projectId, teamId, iteration, userEmail string, params *WorkItemsParams) (*azure.ShortWorkItemsList, *http.Response, error) {
	workItemsList := new(azure.ShortWorkItemsList)
	apiError := new(httputil.APIError)

	path := fmt.Sprintf("%s/%s/_apis/wit/wiql", projectId, teamId)

	query := fmt.Sprintf(workItemsListQuery, userEmail, iteration)
	body := &workItemsListRequestBody{Query: query}

	resp, err := s.sling.New().Post(path).BodyJSON(body).QueryStruct(params).Receive(workItemsList, apiError)
	return workItemsList, resp, httputil.RelevantError(err, apiError)
}

type workItemsBatchRequestBody struct {
	Ids    []int    `json:"ids,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/wit/work%20items/get%20work%20items%20batch?view=azure-devops-rest-5.1
func (s *WorkItemsService) workItemsDescription(projectId string, workItemIds []int, params *WorkItemsParams) (*azure.WorkItemsList, *http.Response, error) {
	workItems := new(azure.WorkItemsList)
	apiError := new(httputil.APIError)

	path := fmt.Sprintf("%s/_apis/wit/workitemsbatch", projectId)

	body := &workItemsBatchRequestBody{Ids: workItemIds, Fields: fields}

	resp, err := s.sling.New().Post(path).BodyJSON(body).QueryStruct(params).Receive(workItems, apiError)
	return workItems, resp, httputil.RelevantError(err, apiError)
}
