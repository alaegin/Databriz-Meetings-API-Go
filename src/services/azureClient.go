package services

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	azureAPI = "https://dev.azure.com/"
)

type AzureClient struct {
	sling    *sling.Sling
	Projects *ProjectsService
	Teams    *TeamsService
}

func NewAzureClient(token string, organization string) *AzureClient {
	httpClient := &http.Client{}
	base := sling.New().Client(httpClient).Base(azureAPI).SetBasicAuth("", token)

	return &AzureClient{
		sling:    base,
		Projects: newProjectsService(base.New(), organization),
		Teams:    newTeamsService(base.New(), organization),
	}
}
