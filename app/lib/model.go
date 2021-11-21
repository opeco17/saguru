package lib

import (
	"time"

	"gorm.io/gorm"
)

type (
	Label struct {
		gorm.Model
		Name    string
		IssueID uint
	}

	User struct {
		gorm.Model
		Name      string
		URL       string
		AvatarURL string
		IssueID   uint
	}

	Issue struct {
		gorm.Model
		GitHubCreatedAt time.Time
		GitHubUpdatedAt time.Time
		URL             string
		PullRequestURL  string
		AssigneesCount  uint
		Issuer          User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Labels          []Label `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		RepositoryID    uint
	}

	Repository struct {
		gorm.Model
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
		Language        string
		Issues          []Issue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
