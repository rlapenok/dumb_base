package app

import database "github.com/rlapenok/dumb_base/internal/data_base"

func InitDataBase() database.DataBase {

	return database.New()

}
