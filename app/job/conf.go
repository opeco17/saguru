package main

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	REPOSITORIES_API_MAX_RESULTS      uint   = 1000 // Number of records in repositories is this value * 3
	REPOSITORIES_API_RESULTS_PER_PAGE uint   = 100
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	MINI_BATCH_SIZE                   uint   = 10
)
