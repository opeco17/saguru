package main

import (
	"fmt"
)

type (
	GetRepositoriesInput struct {
		Page           uint   `query:"page"`
		Labels         string `query:"labels"`
		Assigned       *bool  `query:"assigned"`
		Languages      string `query:"languages"`
		StarCountLower *uint  `query:"star_count_lower"`
		StarCountUpper *uint  `query:"star_count_upper"`
		ForkCountLower *uint  `query:"fork_count_lower"`
		ForkCountUpper *uint  `query:"fork_count_upper"`
		License        string `query:"license"`
		Orderby        string `query:"orderby"`
	}

	GetRepositoriesOutputItemIssue struct {
		ID             uint     `json:"id"`
		Title          string   `json:"title"`
		URL            string   `json:"url"`
		AssigneesCount uint     `json:"assigneesCount"`
		Labels         []string `json:"labels"`
	}

	GetRepositoriesOutputItem struct {
		ID             uint                             `json:"id"`
		Name           string                           `json:"name"`
		URL            string                           `json:"url"`
		Description    string                           `json:"description"`
		StarCount      uint                             `json:"starCount"`
		ForkCount      uint                             `json:"forkCount"`
		OpenIssueCount uint                             `json:"openIssueCount"`
		Topics         string                           `json:"topics"`
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
