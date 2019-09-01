package db

import (
	"Databriz-Meetings-API-Go/models"
	"sync/atomic"
)

type MemoryStorage struct {
	dataRevision int64
	request      *models.ShowRequestBody
}

var storage MemoryStorage

func init() {
	storage = MemoryStorage{}
}

func GetMemoryStorage() *MemoryStorage {
	return &storage
}

func (s *MemoryStorage) GetDataRevision() int64 {
	return s.dataRevision
}

func (s *MemoryStorage) ShouldUpdate(localRevision int64) bool {
	return s.dataRevision > localRevision
}

func (s *MemoryStorage) StoreData(body models.ShowRequestBody) {
	atomic.AddInt64(&s.dataRevision, 1)
	s.request = &body
}

func (s *MemoryStorage) GetData() *models.ShowRequestBody {
	return s.request
}
