package util

import (
	"context"
	"os"
	"strconv"

	errorsutil "opeco17/saguru/lib/errors"
	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/go-github/v41/github"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

func GetMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")
	port, err := strconv.Atoi(os.Getenv("MONGODB_PORT"))
	if err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	client, err := mongodb.GetMongoDBClient(user, password, host, port)
	if err != nil {
		return nil, errorsutil.Wrap(err, "Failed to connect to MongoDB")
	}
	return client, nil
}

func GetMemcachedClient() (*memcache.Client, error) {
	host := os.Getenv("MEMCACHED_HOST")
	port, err := strconv.Atoi(os.Getenv("MEMCACHED_PORT"))
	if err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	client, err := memcached.GetMemcachedClient(host, port)
	if err != nil {
		return nil, errorsutil.Wrap(err, "Failed to connect to Memcached")
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
