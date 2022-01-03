package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"opeco17/gitnavi/lib"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
)

func FetchGitHubRepositoriesSubset(page int, query ...string) (*github.RepositoriesSearchResult, string, error) {
	ctx := context.Background()
	client := getGitHubClient(ctx)

	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
	query = append(query, fmt.Sprintf("pushed:>%s", threeMonthAgo))
	query = append(query, "good-first-issues:>1")
	opts := &github.SearchOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: int(REPOSITORIES_API_RESULTS_PER_PAGE),
		},
	}
	body, resp, _ := client.Search.Repositories(ctx, strings.Join(query, " "), opts)
	if resp.StatusCode >= 400 {
		return nil, "", fmt.Errorf("bad response status code %d\n%s", resp.StatusCode, body)
	}
	return body, strings.Join(query, " "), nil
}

func FetchGitHubRepositories(query ...string) []*github.Repository {
	gitHubRepositories := make([]*github.Repository, 0, REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < int(REPOSITORIES_API_MAX_RESULTS/REPOSITORIES_API_RESULTS_PER_PAGE); page++ {
		gitHubRepositoriesResponse, query, err := FetchGitHubRepositoriesSubset(page, query...)
		if err != nil {
			logrus.Error(err)
			continue
		}
		gitHubRepositories = append(gitHubRepositories, gitHubRepositoriesResponse.Repositories...)
		if page == 0 {
			logrus.Info("Start fetching repositories.")
			logrus.Info(fmt.Sprintf("Query: %v", query))
			logrus.Info(fmt.Sprintf("Total count: %v", *gitHubRepositoriesResponse.Total))
		}
	}
	return gitHubRepositories
}

func FetchRepositories(query ...string) []*lib.Repository {
	gitHubRepositories := FetchGitHubRepositories(query...)
	repositories := make([]*lib.Repository, 0, len(gitHubRepositories))
	for _, gitHubRepository := range gitHubRepositories {
		repositories = append(repositories, convertRepository(gitHubRepository))
	}
	return repositories
}

func FetchGitHubIssues(repositoryName string) ([]*github.Issue, error) {
	ctx := context.Background()
	client := getGitHubClient(ctx)
	repositoryOwner, repositoryName := strings.Split(repositoryName, "/")[0], strings.Split(repositoryName, "/")[1]
	opts := &github.IssueListByRepoOptions{State: "open"}
	body, resp, _ := client.Issues.ListByRepo(ctx, repositoryOwner, repositoryName, opts)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d\n%s", resp.StatusCode, body)
	}
	return body, nil
}

func FetchIssues(RepositoryName string) []*lib.Issue {
	gitHubIssues, err := FetchGitHubIssues(RepositoryName)
	if err != nil {
		logrus.Error(err)
		return []*lib.Issue{}
	}
	issues := make([]*lib.Issue, 0, len(gitHubIssues))
	for _, gitHubIssue := range gitHubIssues {
		issues = append(issues, convertIssue(gitHubIssue))
	}
	return issues
}
