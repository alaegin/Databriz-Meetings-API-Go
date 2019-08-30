package db

import (
	"Databriz-Meetings-API-Go/models"
)

type memoryStorage struct {
	dataRevision int
	request      models.ShowRequestBody
}

var storage memoryStorage

func init() {
	storage = memoryStorage{}
}

func GetMemoryStorage() *memoryStorage {
	return &storage
}

func (storage *memoryStorage) GetDataRevision() int {
	return storage.dataRevision
}

func (storage *memoryStorage) ShouldUpdate(localRevision int) bool {
	return storage.dataRevision > localRevision
}

func (storage *memoryStorage) StoreData(body models.ShowRequestBody) {
	storage.dataRevision++
	storage.request = body
}

func (storage *memoryStorage) GetData() *models.ShowRequestBody {
	return &storage.request
}
