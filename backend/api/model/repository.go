package model

import (
	"fmt"
	"time"
)

type (
	GetRepositoriesInput struct {
		Page           int    `query:"page"`
		Labels         string `query:"labels"`
		IsAssigned     *bool  `query:"isAssigned"`
		Languages      string `query:"languages"`
		StarCountLower *int   `query:"starCountLower"`
		StarCountUpper *int   `query:"starCountUpper"`
		ForkCountLower *int   `query:"forkCountLower"`
		ForkCountUpper *int   `query:"forkCountUpper"`
		License        string `query:"license"`
		Keyword        string `query:"keyword"`
		Orderby        string `query:"orderby"`
	}

	GetRepositoriesOutputItemIssue struct {
		ID                       int       `json:"id"`
		Title                    string    `json:"title"`
		URL                      string    `json:"url"`
		AssigneesCount           int       `json:"assigneesCount"`
		CommentCount             int       `json:"commentCount"`
		GitHubCreatedAt          time.Time `json:"gitHubCreatedAt"`
		GitHubCreatedAtFormatted string    `json:"gitHubCreatedAtFormatted"`
		Labels                   []string  `json:"labels"`
	}

	GetRepositoriesOutputItem struct {
		ID             int                              `json:"id"`
		Name           string                           `json:"name"`
		URL            string                           `json:"url"`
		Description    string                           `json:"description"`
		StarCount      int                              `json:"starCount"`
		ForkCount      int                              `json:"forkCount"`
		OpenIssueCount int                              `json:"openIssueCount"`
		Topics         []string                         `json:"topics"`
		License        string                           `json:"license"`
		Language       string                           `json:"language"`
		Issues         []GetRepositoriesOutputItemIssue `json:"issues"`
	}

	GetRepositoriesOutput struct {
		Items   []GetRepositoriesOutputItem `json:"items"`
		HasNext bool                        `json:"hasNext"`
	}
)

func (input *GetRepositoriesInput) Validate() error {
	if input.StarCountLower != nil && input.StarCountUpper != nil && *input.StarCountLower > *input.StarCountUpper {
		return fmt.Errorf("star_count_lower should be less than star_count_upper")
	}
	if input.ForkCountLower != nil && input.ForkCountUpper != nil && *input.ForkCountLower > *input.ForkCountUpper {
		return fmt.Errorf("fork_count_lower should be less than fork_count_upper")
	}
	return nil
}
