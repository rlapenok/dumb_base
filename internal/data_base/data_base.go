package database

import (
	"strings"
	"sync"

	helperfn "github.com/rlapenok/dumb_base/internal/helper_fn"
	"github.com/sirupsen/logrus"
)

type Api interface {
	GetKeys(mutex *sync.Mutex) string
	UpdateKeys(string, mutex *sync.Mutex)
}

type DataBase struct {
	storage []string
}

func New() DataBase {
	ptr := helperfn.OpenFile()
	storage := helperfn.ReadFile(ptr)
	return DataBase{storage: storage}

}

// impl Api interface
func (db *DataBase) GetKeys(mutex *sync.Mutex) string {
	defer mutex.Unlock()
	return strings.Join(db.storage, " ")
}
func (db *DataBase) UpdateKeys(key string, mutex *sync.Mutex) {
	defer mutex.Unlock()
	defer logrus.Info("Update keys - success")
	new_key := strings.TrimSpace(key)
	db.storage = append(db.storage, new_key)
	ptr := helperfn.OpenFile()
	helperfn.UpdateFile(ptr, new_key, mutex)

}
