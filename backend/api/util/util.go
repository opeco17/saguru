package util

import (
	"opeco17/saguru/lib/database"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")

	client, err := database.GetMongoDBClient(user, password, host)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}

func OrderMetrics() []string {
	return []string{
		"STAR_COUNT",
		"FORK_COUNT",
	}
}
