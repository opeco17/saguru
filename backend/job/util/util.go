package util

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

func GetMongoDBClient() (*mongo.Client, error) {
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
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
	}
	return client, nil
}

func GetMemcachedClient() (*memcache.Client, error) {
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

func GetGitHubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
