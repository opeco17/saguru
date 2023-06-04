package util

import (
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"
	"os"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDBClient() (*mongo.Client, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")
	port, err := strconv.Atoi(os.Getenv("MONGODB_PORT"))
	if err != nil {
		logrus.Error("MONGODB_PORT should be integer")
		return nil, err
	}

	client, err := mongodb.GetMongoDBClient(user, password, host, port)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return client, nil
}

func GetMemcachedClient() (*memcache.Client, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	host := os.Getenv("MEMCACHED_HOST")
	port, err := strconv.Atoi(os.Getenv("MEMCACHED_PORT"))
	if err != nil {
		logrus.Error("MEMCACHED_PORT should be integer")
		return nil, err
	}

	client, err := memcached.GetMemcachedClient(host, port)
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
