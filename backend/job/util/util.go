package util

import (
	"context"
	"fmt"
	"opeco17/saguru/lib/database"
	"os"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

func GetMongoDBClient() (*mongo.Client, error) {
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	host := os.Getenv("MONGODB_HOST")

	client, err := database.GetMongoDBClient(user, password, host)
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return nil, fmt.Errorf(message)
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
