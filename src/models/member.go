package models

import "./azure"

type Member struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromAzureMembersList(azureList *azure.MembersList) *[]Member {
	var list = make([]Member, len(azureList.Members))
	for index, member := range azureList.Members {
		identity := member.Identity
		list[index] = Member{identity.ID, identity.DisplayName, identity.UniqueName}
	}

	return &list
}
