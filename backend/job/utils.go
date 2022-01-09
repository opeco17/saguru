package main

import (
	"context"
	"database/sql"
	"opeco17/gitnavi/lib"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func getDBClient() (*gorm.DB, *sql.DB, error) {
	gormDB, sqlDB, err := lib.GetDBClient(
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	)
	return gormDB, sqlDB, err
}

func getGitHubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
