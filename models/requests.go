package models

type ShowRequestBody struct {
	UserEmail     string `json:"user_email" binding:"required"`
	ProjectId     string `json:"project_id" binding:"required"`
	TeamId        string `json:"team_id" binding:"required"`
	IterationPath string `json:"iteration_path" binding:"required"`
}
