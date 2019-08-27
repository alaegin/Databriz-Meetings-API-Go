package azure

import "time"

// Model for serializing response json from azure api for getting projects list for organization
type ProjectsList struct {
	Count    int       `json:"count"`
	Projects []Project `json:"value"`
}

type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Revision       int       `json:"revision"`
	Visibility     string    `json:"visibility"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}

/*func (list *ProjectsList) ToOwnProjectList() *models.ProjectList {
	var newList = make([]models.Project, len(list.Projects))
	for index, project := range list.Projects {
		newList[index] = models.Project{
			ID:   project.ID,
			Name: project.Name,
		}
	}

	return &models.ProjectList{Projects: newList}
}*/
