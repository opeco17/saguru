package constant

import "time"

const (
	REPOSITORIES_API_URL              string        = "https://api.github.com/search/repositories"
	MAX_REPOSITORY_RECORES            int           = 20000
	REPOSITORIES_API_MAX_RESULTS      int           = 1000 // Fixed by GitHub repositories API (search)
	REPOSITORIES_API_RESULTS_PER_PAGE int           = 100
	REPOSITORIES_API_INTERVAL         time.Duration = 30 * time.Second // Fixed by GitHub repositories API (search)
	REPOSITORIES_API_TIME_OUT         time.Duration = 180 * time.Second
	ISSUES_API_URL                    string        = "https://api.github.com/repos/%s/issues"
	ISSUES_API_RESULTS_PER_PAGE       int           = 100
	ISSUES_API_TIME_OUT               time.Duration = 30 * time.Second
	UPDATE_ISSUE_SIZE                 int           = 5000 // FIXED by GitHub issues API (core)
	UPDATE_ISSUE_SUBSET_SIZE          int           = 10
)
