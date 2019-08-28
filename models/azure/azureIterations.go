package azure

import "time"

type IterationsList struct {
	Count     int         `json:"count"`
	Iteration []Iteration `json:"value"`
}

type Iteration struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	Path       string              `json:"path"`
	Attributes IterationAttributes `json:"attributes"`
	URL        string              `json:"url"`
}

type IterationAttributes struct {
	StartDate  time.Time `json:"startDate"`
	FinishDate time.Time `json:"finishDate"`
	TimeFrame  string    `json:"timeFrame"`
}
