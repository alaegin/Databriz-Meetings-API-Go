package azure

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
	SystemWorkItemType string `json:"System.WorkItemType"`
	SystemReason       string `json:"System.Reason"`
	SystemTitle        string `json:"System.Title"`
}
