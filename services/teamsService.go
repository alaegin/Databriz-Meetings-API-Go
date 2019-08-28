package services

import (
	"Databriz-Meetings-API-Go/httputil"
	"Databriz-Meetings-API-Go/models/azure"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type TeamsService struct {
	sling *sling.Sling
}

func newTeamsService(sling *sling.Sling, organization string) *TeamsService {
	return &TeamsService{
		sling: sling.Path(fmt.Sprintf("%s/", organization)),
	}
}

type TeamMembersParams struct {
	ProjectId string
	TeamId    string
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/core/teams/get%20team%20members%20with%20extended%20properties?view=azure-devops-rest-5.1
func (s *TeamsService) TeamMembers(params *TeamMembersParams) (*azure.MembersList, *http.Response, error) {
	members := new(azure.MembersList)
	path := fmt.Sprintf("_apis/projects/%s/teams/%s/members", params.ProjectId, params.TeamId)
	resp, err := s.sling.New().Get(path).ReceiveSuccess(members)
	return members, resp, httputil.RelevantError(err, resp)
}

type TeamIterationsParams struct {
	ProjectId string
	TeamId    string
}

// Api reference - https://docs.microsoft.com/en-us/rest/api/azure/devops/work/iterations/list?view=azure-devops-rest-5.1
func (s *TeamsService) TeamIterations(params *TeamIterationsParams) (*azure.IterationsList, *http.Response, error) {
	iterations := new(azure.IterationsList)
	path := fmt.Sprintf("%s/%s//_apis/work/teamsettings/iterations", params.ProjectId, params.TeamId)
	resp, err := s.sling.New().Get(path).ReceiveSuccess(iterations)
	return iterations, resp, httputil.RelevantError(err, resp)
}
