package main

import "opeco17/gitnavi/lib"

func convertGetRepositoriesOutputItemIssue(issue lib.Issue) GetRepositoriesOutputItemIssue {
	getRepositoryIssueLabels := make([]string, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		getRepositoryIssueLabels = append(getRepositoryIssueLabels, label.Name)
	}
	getRepositoryIssue := GetRepositoriesOutputItemIssue{
		ID:             int(issue.ID),
		Title:          issue.Title,
		URL:            issue.URL,
		AssigneesCount: *issue.AssigneesCount,
		Labels:         getRepositoryIssueLabels,
	}
	return getRepositoryIssue
}

func convertGetRepositoriesOutputItem(repository lib.Repository) GetRepositoriesOutputItem {
	getRepositoryIssues := make([]GetRepositoriesOutputItemIssue, 0, len(repository.Issues)-1)
	for _, issue := range repository.Issues {
		getRepositoryIssues = append(getRepositoryIssues, convertGetRepositoriesOutputItemIssue(*issue))
	}
	getRepositoriesOutputItem := GetRepositoriesOutputItem{
		ID:             int(repository.ID),
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
