package main

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	MAX_REPOSITORY_RECORES            int    = 10000
	REPOSITORIES_API_MAX_RESULTS      int    = 1000
	REPOSITORIES_API_RESULTS_PER_PAGE int    = 100
	REPOSITORIES_API_INTERVAL_SECOND  int    = 30
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	UPDATE_ISSUE_BATCH_SIZE           int    = 5000
	UPDATE_ISSUE_MINIBATCH_SIZE       int    = 50
	FETCH_ISSUE_CONCURRENCY           int    = 5
)
