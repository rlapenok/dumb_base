package database

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Api interface {
	GetKeys() string
}

type dataBase struct {
	storage []string
}

func New() dataBase {
	bytes_keys, err := os.ReadFile("/home/lprm/my_project/go/github/dumb_base/keys.txt")
	if err != nil {
		logrus.Fatal("Not found keys.txt")
	}
	//From slice bytes to one string
	keys := string(bytes_keys)

	//split string in "\n"
	slice_string := strings.Split(keys, "\n")
	var storage []string
	for _, row_key := range slice_string {
		x := strings.TrimSpace(row_key)
		func(x string) {
			if len(x) != 64 {
				logrus.Warn("Key not supprted")
			} else {
				storage = append(storage, x)
			}

		}(x)
	}

	if len(storage) == 0 {
		logrus.Fatal("Keys not download")
	}
	return dataBase{storage: storage}

}

// impl Api interface
func (db *dataBase) GetKeys() string {

	return strings.Join(db.storage, "")
}
