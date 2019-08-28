package models

import "Databriz-Meetings-API-Go/src/models/azure"

type WorkItem struct {
	ID     int    `json:"id"`
	Rev    int    `json:"rev"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

func FromAzureWorkItems(azureWorkItemsResponse *azure.WorkItemsList) *[]WorkItem {
	var list = make([]WorkItem, len(azureWorkItemsResponse.Value))

	for index, workItem := range azureWorkItemsResponse.Value {
		list[index] = WorkItem{
			ID:     workItem.ID,
			Rev:    workItem.Rev,
			Type:   workItem.Fields.SystemWorkItemType,
			Reason: workItem.Fields.SystemReason,
			Title:  workItem.Fields.SystemTitle,
			URL:    workItem.URL,
		}
	}

	return &list
}
