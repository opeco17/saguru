package lib

import (
	"time"
)

type (
	Label struct {
		LabelID int64  `bson:"label_id"`
		Name    string `bson:"name"`
	}

	User struct {
		UserID    int64  `bson:"user_id"`
		Name      string `bson:"name"`
		URL       string `bson:"url"`
		AvatarURL string `bson:"avatar_url"`
	}

	Issue struct {
		IssueID         int64     `bson:"issue_id"`
		Title           string    `bson:"title"`
		URL             string    `bson:"url"`
		PullRequestURL  string    `bson:"pull_request_url"`
		AssigneesCount  *int      `bson:"assignees_count"`
		CommentCount    *int      `bson:"comment_count"`
		GitHubCreatedAt time.Time `bson:"github_created_at"`
		GitHubUpdatedAt time.Time `bson:"github_updated_at"`
		Labels          []*Label  `bson:"labels"`
		Issuer          *User     `bson:"issuer"`
	}

	Repository struct {
		RepositoryID     int64     `bson:"repository_id"`
		Name             string    `bson:"name"`
		URL              string    `bson:"url"`
		Description      string    `bson:"description"`
		StarCount        *int      `bson:"star_count"`
		ForkCount        *int      `bson:"fork_count"`
		OpenIssueCount   *int      `bson:"open_issue_count"`
		Topics           []string  `bson:"topics"`
		License          string    `bson:"license"`
		Language         string    `bson:"language"`
		UpdatedAt        time.Time `bson:"updated_at"`
		IssueInitialized bool      `bson:"issue_initialized"`
		GitHubCreatedAt  time.Time `bson:"github_created_at"`
		GitHubUpdatedAt  time.Time `bson:"github_updated_at"`
		Issues           []*Issue  `bson:"issues"`
	}

	CachedItem struct {
		Name  string `bson:"name"`
		Count int    `bson:"count"`
	}
)
