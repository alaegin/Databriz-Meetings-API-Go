package services

import (
	"../httputil"
	"../models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type ProjectsService struct {
	Sling *sling.Sling
}

func newProjectsService(sling *sling.Sling, organization string) *ProjectsService {
	return &ProjectsService{
		Sling: sling.Path(fmt.Sprintf("%s/_apis/projects/", organization)),
	}
}

type ProjectsParams struct {
	ApiVersion string `url:"api-version,omitempty"`
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/projects/list?view=azure-devops-rest-5.1
func (s *ProjectsService) Projects(params *ProjectsParams) (*azure.ProjectsList, *http.Response, error) {
	projects := new(azure.ProjectsList)
	apiError := new(httputil.APIError)
	resp, err := s.Sling.New().Get("").QueryStruct(params).Receive(projects, apiError)
	return projects, resp, httputil.RelevantError(err, apiError)
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get%20teams?view=azure-devops-rest-5.1
func (s *ProjectsService) ProjectTeams(projectId string, params *ProjectsParams) (*azure.TeamsList, *http.Response, error) {
	projectTeams := new(azure.TeamsList)
	apiError := new(httputil.APIError)
	path := fmt.Sprintf("%s/teams", projectId)
	resp, err := s.Sling.New().Get(path).QueryStruct(params).Receive(projectTeams, apiError)
	return projectTeams, resp, httputil.RelevantError(err, apiError)
}
