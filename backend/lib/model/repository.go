package model

import "time"

type (
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
)
