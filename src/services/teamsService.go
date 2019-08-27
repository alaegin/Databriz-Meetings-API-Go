package services

import (
	"Databriz-Meetings-API-Go/src/httputil"
	"Databriz-Meetings-API-Go/src/models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type TeamsService struct {
	Sling *sling.Sling
}

type TeamsParams struct {
	ApiVersion string `url:"api-version,omitempty"`
}

func newTeamsService(sling *sling.Sling, organization string) *TeamsService {
	return &TeamsService{
		Sling: sling.Path(fmt.Sprintf("%s/", organization)),
	}
}

func (s *TeamsService) TeamMembers(projectId, teamId string, params *TeamsParams) (*azure.MembersList, *http.Response, error) {
	members := new(azure.MembersList)
	apiError := new(httputil.APIError)
	path := fmt.Sprintf("_apis/projects/%s/teams/%s/members", projectId, teamId)
	resp, err := s.Sling.New().Get(path).QueryStruct(params).Receive(members, apiError)
	return members, resp, httputil.RelevantError(err, apiError)
}

func (s *TeamsService) TeamIterations(projectId, teamId string, params *TeamsParams) (*azure.IterationsList, *http.Response, error) {
	iterations := new(azure.IterationsList)
	apiError := new(httputil.APIError)
	path := fmt.Sprintf("%s/%s//_apis/work/teamsettings/iterations", projectId, teamId)
	resp, err := s.Sling.New().Get(path).QueryStruct(params).Receive(iterations, apiError)
	return iterations, resp, httputil.RelevantError(err, apiError)
}
