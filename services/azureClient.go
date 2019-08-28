package services

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	azureAPI   = "https://dev.azure.com/"
	apiVersion = "5.1"
)

type clientParams struct {
	ApiVersion string `url:"api-version,omitempty"`
}

type AzureClient struct {
	sling     *sling.Sling
	Projects  *ProjectsService
	Teams     *TeamsService
	WorkItems *WorkItemsService
}

func NewAzureClient(token string, organization string) *AzureClient {
	httpClient := &http.Client{}
	base := sling.
		New().
		Client(httpClient).
		Base(azureAPI).
		SetBasicAuth("", token).
		QueryStruct(clientParams{ApiVersion: apiVersion})

	return &AzureClient{
		sling:     base,
		Projects:  newProjectsService(base.New(), organization),
		Teams:     newTeamsService(base.New(), organization),
		WorkItems: newWorkItemsService(base.New(), organization),
	}
}
