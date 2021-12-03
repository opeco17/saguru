package main

import "opeco17/gitnavi/lib"

func convertGetRepositoriesOutputItemIssue(issue lib.Issue) GetRepositoriesOutputItemIssue {
	getRepositoryIssueLabels := make([]string, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		getRepositoryIssueLabels = append(getRepositoryIssueLabels, label.Name)
	}
	getRepositoryIssue := GetRepositoriesOutputItemIssue{
		ID:             issue.ID,
		Title:          issue.Title,
		URL:            issue.URL,
		AssigneesCount: issue.AssigneesCount,
		Labels:         getRepositoryIssueLabels,
	}
	return getRepositoryIssue
}

func convertGetRepositoriesOutputItem(repository lib.Repository) GetRepositoriesOutputItem {
	getRepositoryIssues := make([]GetRepositoriesOutputItemIssue, 0, len(repository.Issues)-1)
	for _, issue := range repository.Issues {
		getRepositoryIssues = append(getRepositoryIssues, convertGetRepositoriesOutputItemIssue(issue))
	}
	getRepositoriesOutputItem := GetRepositoriesOutputItem{
		ID:             repository.ID,
		Name:           repository.Name,
		URL:            repository.URL,
		Description:    repository.Description,
		StarCount:      repository.StarCount,
		ForkCount:      repository.ForkCount,
		OpenIssueCount: repository.OpenIssueCount,
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

func convertGetLanguagesOutput(frontLanguages []lib.FrontLanguage) GetLanguagesOutput {
	GetLanguagesOutputItems := make([]string, 0, len(frontLanguages))
	for _, frontLanguage := range frontLanguages {
		GetLanguagesOutputItems = append(GetLanguagesOutputItems, frontLanguage.Name)
	}
	return GetLanguagesOutput{Items: GetLanguagesOutputItems}
}

func convertGetLicensesOutput(frontLicenses []lib.FrontLicense) GetLicensesOutput {
	getLicensesOutputItems := make([]string, 0, len(frontLicenses))
	for _, frontLicense := range frontLicenses {
		getLicensesOutputItems = append(getLicensesOutputItems, frontLicense.Name)
	}
	return GetLicensesOutput{Items: getLicensesOutputItems}
}

func convertGetLabelsOutput(frontLabels []lib.FrontLabel) GetLabelsOutput {
	getLabelsOutputItems := make([]string, 0, len(frontLabels))
	for _, frontLabel := range frontLabels {
		getLabelsOutputItems = append(getLabelsOutputItems, frontLabel.Name)
	}
	return GetLabelsOutput{Items: getLabelsOutputItems}
}
