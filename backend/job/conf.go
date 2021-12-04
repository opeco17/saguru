package main

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	REPOSITORIES_API_MAX_RESULTS      uint   = 200
	REPOSITORIES_API_RESULTS_PER_PAGE uint   = 100
	REPOSITORIES_API_INTERVAL_SECOND  uint   = 10 // 30
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	MINI_BATCH_SIZE                   uint   = 10
	UPDATE_ISSUE_NUM_PER_BATCH        uint   = 3000
)
