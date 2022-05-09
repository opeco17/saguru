package main

import (
	"fmt"
	"opeco17/gitnavi/lib"
	"time"
)

func convertGetRepositoriesOutputItemIssue(issue lib.Issue) GetRepositoriesOutputItemIssue {
	getRepositoryIssueLabels := make([]string, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		getRepositoryIssueLabels = append(getRepositoryIssueLabels, label.Name)
	}
	getRepositoryIssue := GetRepositoriesOutputItemIssue{
		ID:                       int(issue.IssueID),
		Title:                    issue.Title,
		URL:                      issue.URL,
		AssigneesCount:           *issue.AssigneesCount,
		CommentCount:             *issue.CommentCount,
		GitHubCreatedAt:          issue.GitHubCreatedAt,
		GitHubCreatedAtFormatted: formatGitHubCreatedAt(issue.GitHubCreatedAt),
		Labels:                   getRepositoryIssueLabels,
	}
	return getRepositoryIssue
}

func formatGitHubCreatedAt(createdAt time.Time) string {
	now := time.Now()
	if createdAt.After(now.AddDate(0, 0, -1)) {
		diffHours := int(now.Sub(createdAt).Hours())
		unit := "hours"
		if diffHours == 1 {
			unit = "hour"
		}
		return fmt.Sprintf("%d %s ago", diffHours, unit)
	}
	if createdAt.After(now.AddDate(0, -1, 0)) {
		diffDays := int(now.Sub(createdAt).Hours() / 24)
		unit := "days"
		if diffDays == 1 {
			unit = "day"
		}
		return fmt.Sprintf("%d %s ago", diffDays, unit)
	}
	if createdAt.After(now.AddDate(-1, 0, 0)) {
		return createdAt.Format("2 Jan")
	}
	return createdAt.Format("2 Jan 2006")
}

func convertGetRepositoriesOutputItem(repository lib.Repository) GetRepositoriesOutputItem {
	getRepositoryIssues := make([]GetRepositoriesOutputItemIssue, 0, len(repository.Issues)-1)
	for _, issue := range repository.Issues {
		getRepositoryIssues = append(getRepositoryIssues, convertGetRepositoriesOutputItemIssue(*issue))
	}
	getRepositoriesOutputItem := GetRepositoriesOutputItem{
		ID:             int(repository.RepositoryID),
		Name:           repository.Name,
		URL:            repository.URL,
		Description:    repository.Description,
		StarCount:      *repository.StarCount,
		ForkCount:      *repository.ForkCount,
		OpenIssueCount: *repository.OpenIssueCount,
		Topics:         repository.Topics,
		License:        repository.License,
		Language:       repository.Language,
		Issues:         getRepositoryIssues,
	}
	return getRepositoriesOutputItem
}

func convertGetRepositoriesOutput(repositories []lib.Repository) GetRepositoriesOutput {
	hasNext := len(repositories) > int(RESULTS_PER_PAGE)
	var last int
	if hasNext {
		last = len(repositories) - 1
	} else {
		last = len(repositories)
	}

	getRepositoriesOutputItems := make([]GetRepositoriesOutputItem, 0, last)
	for _, repository := range repositories[:last] {
		getRepositoriesOutputItems = append(getRepositoriesOutputItems, convertGetRepositoriesOutputItem(repository))
	}
	GetRepositoriesOutput := GetRepositoriesOutput{
		Items:   getRepositoriesOutputItems,
		HasNext: hasNext,
	}
	return GetRepositoriesOutput
}

func convertGetLanguagesOutput(cachedLanguages []lib.CachedItem) GetLanguagesOutput {
	outputItems := make([]string, 0, len(cachedLanguages))
	for _, cachedLanguage := range cachedLanguages {
		outputItems = append(outputItems, cachedLanguage.Name)
	}
	return GetLanguagesOutput{Items: outputItems}
}

func convertGetLicensesOutput(cachedLicenses []lib.CachedItem) GetLicensesOutput {
	outputItems := make([]string, 0, len(cachedLicenses))
	for _, cachedLicense := range cachedLicenses {
		outputItems = append(outputItems, cachedLicense.Name)
	}
	return GetLicensesOutput{Items: outputItems}
}

func convertGetLabelsOutput(cachedLabels []lib.CachedItem) GetLabelsOutput {
	outputItems := make([]string, 0, len(cachedLabels))
	for _, cachedLabel := range cachedLabels {
		outputItems = append(outputItems, cachedLabel.Name)
	}
	return GetLabelsOutput{Items: outputItems}
}
