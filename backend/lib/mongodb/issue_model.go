package mongodb

import "time"

type (
	Issue struct {
		IssueID         int64     `bson:"issue_id"`
		Title           string    `bson:"title"`
		URL             string    `bson:"url"`
		AssigneesCount  *int      `bson:"assignees_count"`
		CommentCount    *int      `bson:"comment_count"`
		GitHubCreatedAt time.Time `bson:"github_created_at"`
		GitHubUpdatedAt time.Time `bson:"github_updated_at"`
		Labels          []*Label  `bson:"labels"`
		Issuer          *User     `bson:"issuer"`
	}
)
