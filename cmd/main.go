package main

import (
	"github.com/joho/godotenv"
	"github.com/rlapenok/dumb_base/internal/app"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Not found .env")
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	app.StartGrpcServer()
}
