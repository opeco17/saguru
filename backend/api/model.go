package main

import (
	"fmt"
)

type (
	GetRepositoriesInput struct {
		Page           int    `query:"page"`
		Labels         string `query:"labels"`
		Assigned       *bool  `query:"assigned"`
		Languages      string `query:"languages"`
		StarCountLower *int   `query:"star_count_lower"`
		StarCountUpper *int   `query:"star_count_upper"`
		ForkCountLower *int   `query:"fork_count_lower"`
		ForkCountUpper *int   `query:"fork_count_upper"`
		License        string `query:"license"`
		Orderby        string `query:"orderby"`
	}

	GetRepositoriesOutputItemIssue struct {
		ID             int      `json:"id"`
		Title          string   `json:"title"`
		URL            string   `json:"url"`
		AssigneesCount int      `json:"assigneesCount"`
		Labels         []string `json:"labels"`
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

	GetLanguagesOutput struct {
		Items []string `json:"items"`
	}

	GetLicensesOutput struct {
		Items []string `json:"items"`
	}

	GetLabelsOutput struct {
		Items []string `json:"items"`
	}

	GetOrderMetricsOutput struct {
		Items []string `json:"items"`
	}
)

func (input *GetRepositoriesInput) validator() error {
	if input.StarCountLower != nil && input.StarCountUpper != nil && *input.StarCountLower > *input.StarCountUpper {
		return fmt.Errorf("star_count_lower should be less than star_count_upper")
	}
	if input.ForkCountLower != nil && input.ForkCountUpper != nil && *input.ForkCountLower > *input.ForkCountUpper {
		return fmt.Errorf("fork_count_lower should be less than fork_count_upper")
	}
	return nil
}
