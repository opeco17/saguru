package main

import (
	"opeco17/gitnavi/lib"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func getMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")

	client, err := lib.GetMongoDBClient(user, password, host)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}

func orderMetrics() []string {
	return []string{
		"STAR_COUNT",
		"FORK_COUNT",
	}
}
