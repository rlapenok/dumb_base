package main

import (
	"github.com/rlapenok/dumb_base/internal/app"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	app.StartGrpcServer()
}
