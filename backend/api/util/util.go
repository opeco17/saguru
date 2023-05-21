package util

import (
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/lib/database"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDBClient() (*mongo.Client, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

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
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	return []string{
		"STAR_COUNT",
		"FORK_COUNT",
	}
}
