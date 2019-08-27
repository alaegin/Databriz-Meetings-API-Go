package services

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	azureAPI = "https://dev.azure.com/"
)

type Client struct {
	Sling    *sling.Sling
	Projects *ProjectsService
}

func NewClient(token string, organization string) *Client {
	base := sling.New().Client(&http.Client{}).Base(azureAPI).SetBasicAuth("", token)
	return &Client{
		Sling:    base,
		Projects: NewProjectsService(base.New(), organization),
	}
}
