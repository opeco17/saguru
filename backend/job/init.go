package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

func initDB() error {
	logrus.Info("Start initializing DB.")

	client, err := getMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	client.Database("main").Collection("repositories")
	return nil
}
