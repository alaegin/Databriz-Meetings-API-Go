package models

import "Databriz-Meetings-API-Go/src/models/azure"

type Member struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromAzureMembersList(azureList *azure.MembersList) *[]Member {
	var list = make([]Member, len(azureList.Members))
	for index, member := range azureList.Members {
		identity := member.Identity
		list[index] = Member{
			ID:    identity.ID,
			Name:  identity.DisplayName,
			Email: identity.UniqueName}
	}

	return &list
}
