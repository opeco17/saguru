package main

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	REPOSITORIES_API_MAX_RESULTS      uint   = 1000
	REPOSITORIES_API_RESULTS_PER_PAGE uint   = 100
	REPOSITORIES_API_INTERVAL_SECOND  uint   = 30
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	UPDATE_ISSUE_MINI_BATCH_SIZE      uint   = 10
	UPDATE_ISSUE_BATCH_SIZE           uint   = 3000
)
