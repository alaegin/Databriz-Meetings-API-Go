package models

import "./azure"

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FromAzureProjectsList(azureList *azure.ProjectsList) *[]Project {
	var list = make([]Project, len(azureList.Projects))
	for index, project := range azureList.Projects {
		list[index] = Project{
			ID:   project.ID,
			Name: project.Name,
		}
	}

	return &list
}
