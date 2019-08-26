package azure

import "time"

type WiqlWorkItemsResponse struct {
	QueryType       string    `json:"queryType"`
	QueryResultType string    `json:"queryResultType"`
	AsOf            time.Time `json:"asOf"`
	//Columns         []Columns     `json:"columns"`
	//SortColumns     []SortColumns `json:"sortColumns"`
	WorkItems []WorkItems `json:"workItems"`
}

/*type Columns struct {
	ReferenceName string `json:"referenceName"`
	Name          string `json:"name"`
	URL           string `json:"url"`
}*/
/*type Field struct {
	ReferenceName string `json:"referenceName"`
	Name          string `json:"name"`
	URL           string `json:"url"`
}*/
/*type SortColumns struct {
	Field      Field `json:"field"`
	Descending bool  `json:"descending"`
}*/
type WorkItems struct {
	ID int `json:"id"`
	//URL string `json:"url"`
}

type WorkItemsResponse struct {
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
