package models

import (
	"Databriz-Meetings-API-Go/models/azure"
	"time"
)

type WorkItem struct {
	ID               int       `json:"id"`
	Rev              int       `json:"rev"`
	Type             string    `json:"type"`
	State            string    `json:"state"`
	Reason           string    `json:"reason"`
	CreationDate     time.Time `json:"creation_date"`
	Title            string    `json:"title"`
	OriginalEstimate float64   `json:"original_estimate"`
	CompletedTime    float64   `json:"completed_time"`
	Priority         int       `json:"priority"`
	URL              string    `json:"url"`
}

func FromAzureWorkItems(azureWorkItemsResponse *azure.WorkItemsList) *[]WorkItem {
	var list = make([]WorkItem, len(azureWorkItemsResponse.Value))

	for index, workItem := range azureWorkItemsResponse.Value {
		list[index] = WorkItem{
			ID:               workItem.ID,
			Rev:              workItem.Rev,
			Type:             workItem.Fields.SystemWorkItemType,
			State:            workItem.Fields.SystemState,
			CreationDate:     workItem.Fields.SystemCreatedDate,
			Reason:           workItem.Fields.SystemReason,
			Title:            workItem.Fields.SystemTitle,
			OriginalEstimate: workItem.Fields.MicrosoftVSTSSchedulingOriginalEstimate,
			CompletedTime:    workItem.Fields.MicrosoftVSTSSchedulingCompletedWork,
			Priority:         workItem.Fields.MicrosoftVSTSCommonPriority,
			URL:              workItem.URL,
		}
	}

	return &list
}
