package main

import (
	"time"

	"gorm.io/gorm"
)

type Label struct {
	gorm.Model
	Name    string `gorm:"unique"`
	IssueID uint
}

type User struct {
	gorm.Model
	Name      string
	URL       string
	AvatarURL string
	GitHubID  uint `gorm:"unique"`
	IssueID   uint
}

type Issue struct {
	gorm.Model
	URL             string
	PullRequestURL  string
	GitHubID        uint `gorm:"unique"`
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	Issuer          User
	Assignees       []User
	Labels          []Label
	RepositoryID    uint
}

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
	GitHubID        uint `gorm:"unique"`
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	Issues          []Issue
}
