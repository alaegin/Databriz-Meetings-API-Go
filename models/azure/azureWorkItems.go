package azure

import "time"

type ShortWorkItemsList struct {
	ShortWorkItems []ShortWorkItem `json:"workItems"`
}

type ShortWorkItem struct {
	ID int `json:"id"`
}

type WorkItemsList struct {
	Count int        `json:"count"`
	Value []WorkItem `json:"value"`
}

type WorkItem struct {
	ID     int    `json:"id"`
	Rev    int    `json:"rev"`
	Fields Fields `json:"fields"`
	URL    string `json:"url"`
}

type Fields struct {
	SystemWorkItemType                      string    `json:"System.WorkItemType"`
	SystemState                             string    `json:"System.State"`
	SystemReason                            string    `json:"System.Reason"`
	SystemCreatedDate                       time.Time `json:"System.CreatedDate"`
	SystemTitle                             string    `json:"System.Title"`
	MicrosoftVSTSSchedulingOriginalEstimate float64   `json:"Microsoft.VSTS.Scheduling.OriginalEstimate"`
	MicrosoftVSTSSchedulingCompletedWork    float64   `json:"Microsoft.VSTS.Scheduling.CompletedWork"`
	MicrosoftVSTSCommonPriority             int       `json:"Microsoft.VSTS.Common.Priority"`
}
