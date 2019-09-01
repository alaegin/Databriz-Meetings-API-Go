package db

import (
	"Databriz-Meetings-API-Go/models"
)

type MemoryStorage struct {
	dataRevision int
	request      *models.ShowRequestBody
}

var storage MemoryStorage

func init() {
	storage = MemoryStorage{}
}

func GetMemoryStorage() *MemoryStorage {
	return &storage
}

func (storage *MemoryStorage) GetDataRevision() int {
	return storage.dataRevision
}

func (storage *MemoryStorage) ShouldUpdate(localRevision int) bool {
	return storage.dataRevision > localRevision
}

func (storage *MemoryStorage) StoreData(body models.ShowRequestBody) {
	storage.dataRevision++
	storage.request = &body
}

func (storage *MemoryStorage) GetData() *models.ShowRequestBody {
	return storage.request
}
