package models

import "./azure"

type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FromAzureTeamsList(azureList *azure.TeamsList) *[]Team {
	var list = make([]Team, len(azureList.Teams))
	for index, team := range azureList.Teams {
		list[index] = Team{
			ID:          team.ID,
			Name:        team.Name,
			Description: team.Description,
		}
	}

	return &list
}
