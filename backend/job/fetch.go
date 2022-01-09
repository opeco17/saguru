package main

import (
	"context"
	"fmt"
	"strings"

	"opeco17/gitnavi/lib"

	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
)

func FetchGitHubRepositoriesSubset(page int, queries ...string) (*github.RepositoriesSearchResult, string, error) {
	ctx := context.Background()
	client := getGitHubClient(ctx)
	opts := &github.SearchOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: REPOSITORIES_API_RESULTS_PER_PAGE,
		},
	}
	body, resp, _ := client.Search.Repositories(ctx, strings.Join(queries, " "), opts)
	if resp.StatusCode >= 400 {
		return nil, "", fmt.Errorf("bad response status code %d\n%v", resp.StatusCode, body)
	}
	return body, strings.Join(queries, " "), nil
}

func FetchGitHubRepositories(queries ...string) []*github.Repository {
	gitHubRepositories := make([]*github.Repository, 0, REPOSITORIES_API_MAX_RESULTS)
	for page := 0; page < REPOSITORIES_API_MAX_RESULTS/REPOSITORIES_API_RESULTS_PER_PAGE; page++ {
		gitHubRepositoriesResponse, queries, err := FetchGitHubRepositoriesSubset(page, queries...)
		if err != nil {
			logrus.Error(err)
			continue
		}
		gitHubRepositories = append(gitHubRepositories, gitHubRepositoriesResponse.Repositories...)
		if page == 0 {
			logrus.Info("Start fetching repositories.")
			logrus.Info(fmt.Sprintf("Query: %v", queries))
			logrus.Info(fmt.Sprintf("Total count: %v", *gitHubRepositoriesResponse.Total))
		}
	}
	return gitHubRepositories
}

func FetchRepositories(queries ...string) []*lib.Repository {
	gitHubRepositories := FetchGitHubRepositories(queries...)
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
		return nil, fmt.Errorf("bad response status code %d\n%v", resp.StatusCode, body)
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
