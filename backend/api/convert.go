package main

import "opeco17/oss-book/lib"

func convertGetRepositoryIssue(issue lib.Issue) GetRepositoryIssue {
	getRepositoryIssueLabels := make([]string, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		getRepositoryIssueLabels = append(getRepositoryIssueLabels, label.Name)
	}
	getRepositoryIssue := GetRepositoryIssue{
		URL:            issue.URL,
		AssigneesCount: issue.AssigneesCount,
		Labels:         getRepositoryIssueLabels,
	}
	return getRepositoryIssue
}

func convertGetRepository(repository lib.Repository) GetRepository {
	getRepositoryIssues := make([]GetRepositoryIssue, 0, len(repository.Issues))
	for _, issue := range repository.Issues {
		getRepositoryIssues = append(getRepositoryIssues, convertGetRepositoryIssue(issue))
	}
	getRepository := GetRepository{
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
	return getRepository
}

func convertGetRepositoriesOutput(repositories []lib.Repository) GetRepositoriesOutput {
	getRepositoriesOutput := make(GetRepositoriesOutput, 0, len(repositories))
	for _, repository := range repositories {
		getRepositoriesOutput = append(getRepositoriesOutput, convertGetRepository(repository))
	}
	return getRepositoriesOutput
}

func convertGetLanguagesOutput(frontLanguages []lib.FrontLanguage) GetLanguagesOutput {
	getLanguagesOutput := make(GetLanguagesOutput, 0, len(frontLanguages))
	for _, frontLanguage := range frontLanguages {
		getLanguagesOutput = append(getLanguagesOutput, frontLanguage.Name)
	}
	return getLanguagesOutput
}

func convertGetLicensesOutput(frontLicenses []lib.FrontLicense) GetLicensesOutput {
	getLicensesOutput := make(GetLicensesOutput, 0, len(frontLicenses))
	for _, frontLicense := range frontLicenses {
		getLicensesOutput = append(getLicensesOutput, frontLicense.Name)
	}
	return getLicensesOutput
}

func convertGetLabelsOutput(frontLabels []lib.FrontLabel) GetLabelsOutput {
	getLabelsOutput := make(GetLabelsOutput, 0, len(frontLabels))
	for _, frontLabel := range frontLabels {
		getLabelsOutput = append(getLabelsOutput, frontLabel.Name)
	}
	return getLabelsOutput
}
