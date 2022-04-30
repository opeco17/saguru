package lib

import (
	"time"
)

type (
	Label struct {
		LabelID uint   `bson:"label_id"`
		Name    string `bson:"name"`
	}

	User struct {
		UserID    uint   `bson:"user_id"`
		Name      string `bson:"name"`
		URL       string `bson:"url"`
		AvatarURL string `bson:"avatar_url"`
	}

	Issue struct {
		IssueID         uint      `bson:"issue_id"`
		Title           string    `bson:"title"`
		URL             string    `bson:"url"`
		PullRequestURL  string    `bson:"pull_request_url"`
		AssigneesCount  *int      `bson:"assignees_count"`
		CommentCount    *int      `bson:"comment_count"`
		Issuer          *User     `bson:"issuer"`
		Labels          []*Label  `bson:"labels"`
		GitHubCreatedAt time.Time `bson:"github_created_at"`
		GitHubUpdatedAt time.Time `bson:"github_updated_at"`
	}

	Repository struct {
		RepositoryID     uint      `bson:"repository_id"`
		Name             string    `bson:"name"`
		URL              string    `bson:"url"`
		Description      string    `bson:"description"`
		StarCount        *int      `bson:"star_count"`
		ForkCount        *int      `bson:"fork_count"`
		OpenIssueCount   *int      `bson:"open_issue_count"`
		Topics           string    `bson:"topics"`
		License          string    `bson:"license"`
		Language         string    `bson:"language"`
		GitHubCreatedAt  time.Time `bson:"github_created_at"`
		GitHubUpdatedAt  time.Time `bson:"github_updated_at"`
		IssueInitialized bool      `bson:"issue_initialized"`
		Issues           []*Issue  `bson:"issues"`
	}

	FrontLanguage struct {
		Name            string `bson:"name"`
		RepositoryCount int    `bson:"repository_count"`
	}

	FrontLicense struct {
		Name            string `bson:"name"`
		RepositoryCount int    `bson:"repository_count"`
	}

	FrontLabel struct {
		Name       string `bson:"name"`
		IssueCount int    `bson:"issue_count"`
	}
)
