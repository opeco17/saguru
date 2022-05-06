package main

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	MAX_REPOSITORY_RECORES            int    = 20000
	REPOSITORIES_API_MAX_RESULTS      int    = 1000 // Fixed by GitHub repositories API (search)
	REPOSITORIES_API_RESULTS_PER_PAGE int    = 100
	REPOSITORIES_API_INTERVAL_SECOND  int    = 30 // Fixed by GitHub repositories API (search)
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	UPDATE_ISSUE_COUNT_PER_BATCH      int    = 5000 // FIXED by GitHub issues API (core)
	FETCH_ISSUE_CONCURRENCY           int    = 5
)
