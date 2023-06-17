package database

import (
	"errors"
	"strings"
	"sync"

	helperfn "github.com/rlapenok/dumb_base/internal/helper_fn"
)

type Api interface {
	GetKeys()
	UpdateKeys(string, chan error)
}

type DataBase struct {
	storage []string
	mutex   sync.Mutex
}

func New() DataBase {

	ptr := helperfn.OpenFile()
	storage := helperfn.ReadFile(ptr)
	return DataBase{storage: storage, mutex: sync.Mutex{}}

}

// impl Api interface
func (db *DataBase) GetKeys(channel chan string) {
	defer db.mutex.Unlock()
	db.mutex.Lock()
	keys := strings.Join(db.storage, " ")
	channel <- keys
}
func (db *DataBase) UpdateKeys(key string, channel chan error) {
	defer db.mutex.Unlock()
	defer close(channel)
	db.mutex.Lock()
	new_key := strings.TrimSpace(key)
	if len(new_key) != 64 {
		err := errors.New(":Key not supported")
		channel <- err
	} else {
		db.storage = append(db.storage, new_key)
		ptr := helperfn.OpenFile()
		if result := helperfn.UpdateFile(ptr, new_key); result != nil {
			channel <- result
		}
		channel <- nil
	}

}
