package main

import (
	"context"
	"fmt"
	"opeco17/gitnavi/lib"
	"os"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

func getMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")

	client, err := lib.GetMongoDBClient(user, password, host)
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
	}
	return client, nil
}

func getGitHubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
