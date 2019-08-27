package services

import (
	"../httputil"
	"../models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type ProjectsService struct {
	sling *sling.Sling
}

func NewProjectsService(sling *sling.Sling, organization string) *ProjectsService {
	return &ProjectsService{
		sling: sling.Path(fmt.Sprintf("%s/_apis/projects/", organization)),
	}
}

type ProjectsParams struct {
	ApiVersion string `url:"api-version,omitempty"`
}

func (s *ProjectsService) Projects(params *ProjectsParams) (*azure.ProjectsList, *http.Response, error) {
	projects := new(azure.ProjectsList)
	apiError := new(httputil.APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(projects, apiError)
	return projects, resp, httputil.RelevantError(err, apiError)
}

func (s *ProjectsService) ProjectTeams(projectId string, params *ProjectsParams) (*azure.TeamsList, *http.Response, error) {
	projectTeams := new(azure.TeamsList)
	apiError := new(httputil.APIError)
	path := fmt.Sprintf("%s/teams", projectId)
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(projectTeams, apiError)
	return projectTeams, resp, httputil.RelevantError(err, apiError)
}
