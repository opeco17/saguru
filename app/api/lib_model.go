package main

import (
	"time"
)

type Label struct {
	Name   string  `gorm:"primaryKey"`
	Issues []Issue `gorm:"many2many:issue_labels;"`
}

type User struct {
	GitHubID  uint `gorm:"primaryKey"`
	Name      string
	URL       string
	AvatarURL string
	IssueID   uint
}

type Issue struct {
	GitHubID        uint `gorm:"primaryKey"`
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	URL             string
	PullRequestURL  string
	AssigneesCount  uint
	Issuer          User
	Labels          []Label `gorm:"many2many:issue_labels;"`
	RepositoryID    uint
}

type Repository struct {
	GitHubID        uint `gorm:"primaryKey"`
	GitHubCreatedAt time.Time
	GitHubUpdatedAt time.Time
	Name            string
	URL             string
	Description     string
	StarCount       uint
	ForkCount       uint
	OpenIssueCount  uint
	Topics          string
	License         string
	Issues          []Issue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
