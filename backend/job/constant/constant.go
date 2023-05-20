package constant

const (
	REPOSITORIES_API_URL              string = "https://api.github.com/search/repositories"
	MAX_REPOSITORY_RECORES            int    = 20000
	REPOSITORIES_API_MAX_RESULTS      int    = 1000 // Fixed by GitHub repositories API (search)
	REPOSITORIES_API_RESULTS_PER_PAGE int    = 100
	REPOSITORIES_API_INTERVAL_SECOND  int    = 30 // Fixed by GitHub repositories API (search)
	ISSUES_API_URL                    string = "https://api.github.com/repos/%s/issues"
	ISSUES_API_RESULTS_PER_PAGE       int    = 100
	UPDATE_ISSUE_BATCH_SIZE           int    = 5000 // FIXED by GitHub issues API (core)
	UPDATE_ISSUE_MINIBATCH_SIZE       int    = 50
	FETCH_ISSUE_CONCURRENCY           int    = 5
)
