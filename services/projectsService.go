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
	resp, err := s.sling.New().Get("").ReceiveSuccess(projects)
	return projects, resp, httputil.RelevantError(err, resp)
}

type ProjectTeamsParams struct {
	ProjectId string
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get%20teams?view=azure-devops-rest-5.1
func (s *ProjectsService) ProjectTeams(params *ProjectTeamsParams) (*azure.TeamsList, *http.Response, error) {
	projectTeams := new(azure.TeamsList)
	path := fmt.Sprintf("%s/teams", params.ProjectId)
	resp, err := s.sling.New().Get(path).ReceiveSuccess(projectTeams)
	return projectTeams, resp, httputil.RelevantError(err, resp)
}
