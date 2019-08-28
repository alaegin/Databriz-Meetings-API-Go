package services

import (
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type ProjectsService struct {
	sling *sling.Sling
}

func newProjectsService(sling *sling.Sling, organization string) *ProjectsService {
	return &ProjectsService{
		sling: sling.Path(fmt.Sprintf("%s/_apis/projects/", organization)),
	}
}

type ProjectsParams struct{}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/list?view=azure-devops-rest-5.1
func (s *ProjectsService) Projects(params *ProjectsParams) (*azure.ProjectsList, *http.Response, error) {
	projects := new(azure.ProjectsList)
	apiError := new(httputil.APIError)
	resp, err := s.sling.New().Get("").Receive(projects, apiError)
	return projects, resp, httputil.RelevantError(err, apiError)
}

type ProjectTeamsParams struct {
	ProjectId string
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get%20teams?view=azure-devops-rest-5.1
func (s *ProjectsService) ProjectTeams(params *ProjectTeamsParams) (*azure.TeamsList, *http.Response, error) {
	projectTeams := new(azure.TeamsList)
	apiError := new(httputil.APIError)
	path := fmt.Sprintf("%s/teams", params.ProjectId)
	resp, err := s.sling.New().Get(path).Receive(projectTeams, apiError)
	return projectTeams, resp, httputil.RelevantError(err, apiError)
}
