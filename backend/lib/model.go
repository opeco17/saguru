package lib

import (
	"time"

	"gorm.io/gorm"
)

type (
	Label struct {
		gorm.Model
		Name    string
		IssueID int
	}

	User struct {
		gorm.Model
		Name      string
		URL       string
		AvatarURL string
		IssueID   int
	}

	Issue struct {
		gorm.Model
		GitHubCreatedAt time.Time
		GitHubUpdatedAt time.Time
		Title           string
		URL             string
		PullRequestURL  string
		AssigneesCount  *int
		CommentCount    *int
		Issuer          *User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Labels          []Label `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		RepositoryID    int
	}

	Repository struct {
		gorm.Model
		GitHubCreatedAt  time.Time
		GitHubUpdatedAt  time.Time
		Name             string
		URL              string
		Description      string
		StarCount        *int
		ForkCount        *int
		OpenIssueCount   *int
		Topics           string
		License          string
		Language         string
		IssueInitialized bool
		Issues           []*Issue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}

	FrontLanguage struct {
		gorm.Model
		Name            string
		RepositoryCount int
	}

	FrontLicense struct {
		gorm.Model
		Name            string
		RepositoryCount int
	}

	FrontLabel struct {
		gorm.Model
		Name       string
		IssueCount int
	}
)
