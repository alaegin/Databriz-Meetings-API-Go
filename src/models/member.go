package models

import "../models/azure"

type MembersList struct {
	Members []Member `json:"members"`
}

type Member struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromAzureMembersList(list *azure.MembersList) *MembersList {
	var newList = make([]Member, len(list.Members))
	for index, element := range list.Members {
		identity := element.Identity
		newList[index] = Member{identity.ID, identity.DisplayName, identity.UniqueName}
	}

	return &MembersList{Members: newList}
}
