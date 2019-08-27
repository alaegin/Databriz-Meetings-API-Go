package services

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	AzureAPI = "https://dev.azure.com/"
)

type Client struct {
	sling *sling.Sling

	Projects *ProjectsService
}

func NewClient(token string, organization string) *Client {
	base := sling.New().Client(&http.Client{}).Base(AzureAPI).SetBasicAuth("", token)
	return &Client{
		sling:    base,
		Projects: NewProjectsService(base.New(), organization),
	}
}
