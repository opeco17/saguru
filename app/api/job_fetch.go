package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func FetchGitHubRepositoriesSubset(page int, query ...string) (*GitHubRepositoriesResponse, string, error) {
	month_ago := time.Now().AddDate(0, -1, 0).Format("2006-01-02T15:04:05+09:00")
	query = append(query, fmt.Sprintf("pushed:>%s", month_ago))

	client := &http.Client{}
	client.Timeout = time.Second * 15

	request, err := http.NewRequest("GET", REPOSITORIES_API_URL, nil)
	if err != nil {
		return nil, "", err
	}
	request.SetBasicAuth(os.Getenv("GITHUB_API_USER"), os.Getenv("GITHUB_API_TOKEN"))

	params := request.URL.Query()
	params.Add("q", strings.Join(query, " "))
	params.Add("type", "Repositories")
	params.Add("per_page", strconv.Itoa(int(REPOSITORIES_API_RESULTS_PER_PAGE)))
	params.Add("page", strconv.Itoa(page))
	params.Add("sort", "updated")
	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}
	if response.StatusCode >= 400 {
		return nil, "", fmt.Errorf("bad response status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}
	responseBody := new(GitHubRepositoriesResponse)
	json.Unmarshal([]byte(body), responseBody)

	return responseBody, strings.Join(query, " "), nil
}

func FetchGitHubRepositories(query ...string) []GitHubRepository {
	gitHubRepositories := make([]GitHubRepository, 0, REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < int(REPOSITORIES_API_MAX_RESULTS/REPOSITORIES_API_RESULTS_PER_PAGE); page++ {
		gitHubRepositoriesResponse, query, err := FetchGitHubRepositoriesSubset(page, query...)
		if err != nil {
			log.Fatal(err)
			continue
		}
		gitHubRepositories = append(gitHubRepositories, gitHubRepositoriesResponse.Repositories...)
		if page == 0 {
			log.Info("Start fetching repositories")
			log.Info(fmt.Sprintf("Query: %v", query))
			log.Info(fmt.Sprintf("Total count: %v", gitHubRepositoriesResponse.TotalCount))
		}
	}
	return gitHubRepositories
}

func FetchRepositories(query ...string) []Repository {
	gitHubRepositories := FetchGitHubRepositories(query...)
	repositories := make([]Repository, 0, len(gitHubRepositories))
	for _, gitHubRepository := range gitHubRepositories {
		repositories = append(repositories, gitHubRepository.convert())
	}
	return repositories
}

func FetchGitHubIssues(repositoryName string) (GitHubIssuesResponse, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 15

	url := fmt.Sprintf(ISSUES_API_URL, repositoryName)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(os.Getenv("GITHUB_API_USER"), os.Getenv("GITHUB_API_TOKEN"))

	params := request.URL.Query()
	params.Add("state", "open")
	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseBody GitHubIssuesResponse
	json.Unmarshal([]byte(body), &responseBody)

	return responseBody, nil
}

func FetchIssues(RepositoryName string) []Issue {
	gitHubIssues, err := FetchGitHubIssues(RepositoryName)
	if err != nil {
		log.Fatal(err)
		return []Issue{}
	}
	issues := make([]Issue, 0, len(gitHubIssues))
	for _, gitHubIssue := range gitHubIssues {
		issues = append(issues, gitHubIssue.convert())
	}
	return issues
}
