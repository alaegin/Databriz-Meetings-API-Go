package models

import (
	"./azure"
	"time"
)

type Iteration struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	StartDate  time.Time `json:"startDate"`
	FinishDate time.Time `json:"finishDate"`
	TimeFrame  string    `json:"timeFrame"`
}

func FromAzureIterations(azureList *azure.IterationsList) *[]Iteration {
	var list = make([]Iteration, len(azureList.Iteration))
	for index, iteration := range azureList.Iteration {
		list[index] = Iteration{
			ID:         iteration.ID,
			Name:       iteration.Name,
			Path:       iteration.Path,
			StartDate:  iteration.Attributes.StartDate,
			FinishDate: iteration.Attributes.FinishDate,
			TimeFrame:  iteration.Attributes.TimeFrame,
		}
	}

	return &list
}
