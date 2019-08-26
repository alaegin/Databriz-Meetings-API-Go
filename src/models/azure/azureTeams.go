package azure

type TeamsList struct {
	Teams []Team `json:"value"`
	Count int    `json:"count"`
}

type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	IdentityURL string `json:"identityUrl"`
	ProjectName string `json:"projectName"`
	ProjectID   string `json:"projectId"`
}
