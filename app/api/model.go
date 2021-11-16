package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repository struct {
	gorm.Model
	Name            string
	URL             string
	Description     string
	StarCount       uint
	ForkCount       uint
	OpenIssueCount  uint
	License         string
	Topics          string
	GitHubID        uint
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	Owner           User
	Issues          []Issue
}

type User struct {
	gorm.Model
	Name      string
	URL       string
	AvatarURL string
	GitHubID  uint
}

type Issue struct {
	gorm.Model
	URL             string
	PullRequestURL  string
	GitHubID        uint
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	Issuer          User
	Assignees       []User
	Labels          []Label
}

type Label struct {
	gorm.Model
	Name  string
	Color string
}
