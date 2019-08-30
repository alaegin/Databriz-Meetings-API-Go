package db

import (
	"Databriz-Meetings-API-Go/models"
)

var dataChanged bool = false
var request models.ShowRequestBody

func DataHasChanged() bool {
	return dataChanged
}

func StoreData(body models.ShowRequestBody) {
	request = body
	dataChanged = true
}

func GetData() *models.ShowRequestBody {
	dataChanged = false
	return &request
}
