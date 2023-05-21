package controller

import (
	"context"
	"fmt"
	"net/http"
	"opeco17/saguru/api/constant"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/service"
	"opeco17/saguru/api/util"
	libModel "opeco17/saguru/lib/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetRepositories(c echo.Context) error {
	logrus.Info("Get repositories")

	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	// Validation
	input := new(model.GetRepositoriesInput)
	logrus.Info(fmt.Sprintf("Input %+v", input))
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := input.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Connect DB
	client, err := util.GetMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	repositories, err := service.GetRepositoriesFromDB(client, input)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get repositories from database.")
	}
	repositories = service.FilterIssuesInRepositories(repositories, input)
	output := convertGetRepositoriesOutput(repositories)

	logrus.Info(fmt.Sprintf("Total time to fetch repositories: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func convertGetRepositoriesOutput(repositories []libModel.Repository) model.GetRepositoriesOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	hasNext := len(repositories) > int(constant.RESULTS_PER_PAGE)
	var last int
	if hasNext {
		last = len(repositories) - 1
	} else {
		last = len(repositories)
	}

	getRepositoriesOutputItems := make([]model.GetRepositoriesOutputItem, 0, last)
	for _, repository := range repositories[:last] {
		getRepositoriesOutputItems = append(getRepositoriesOutputItems, convertGetRepositoriesOutputItem(repository))
	}
	GetRepositoriesOutput := model.GetRepositoriesOutput{
		Items:   getRepositoriesOutputItems,
		HasNext: hasNext,
	}
	return GetRepositoriesOutput
}

func convertGetRepositoriesOutputItem(repository libModel.Repository) model.GetRepositoriesOutputItem {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	getRepositoryIssues := make([]model.GetRepositoriesOutputItemIssue, 0, len(repository.Issues)-1)
	for _, issue := range repository.Issues {
		getRepositoryIssues = append(getRepositoryIssues, convertGetRepositoriesOutputItemIssue(*issue))
	}
	getRepositoriesOutputItem := model.GetRepositoriesOutputItem{
		ID:             int(repository.RepositoryID),
		Name:           repository.Name,
		URL:            repository.URL,
		Description:    repository.Description,
		StarCount:      *repository.StarCount,
		ForkCount:      *repository.ForkCount,
		OpenIssueCount: *repository.OpenIssueCount,
		Topics:         repository.Topics,
		License:        repository.License,
		Language:       repository.Language,
		Issues:         getRepositoryIssues,
	}
	return getRepositoriesOutputItem
}

func convertGetRepositoriesOutputItemIssue(issue libModel.Issue) model.GetRepositoriesOutputItemIssue {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	getRepositoryIssueLabels := make([]string, 0, len(issue.Labels))
	for _, label := range issue.Labels {
		getRepositoryIssueLabels = append(getRepositoryIssueLabels, label.Name)
	}
	getRepositoryIssue := model.GetRepositoriesOutputItemIssue{
		ID:                       int(issue.IssueID),
		Title:                    issue.Title,
		URL:                      issue.URL,
		AssigneesCount:           *issue.AssigneesCount,
		CommentCount:             *issue.CommentCount,
		GitHubCreatedAt:          issue.GitHubCreatedAt,
		GitHubCreatedAtFormatted: formatGitHubCreatedAt(issue.GitHubCreatedAt),
		Labels:                   getRepositoryIssueLabels,
	}
	return getRepositoryIssue
}

func formatGitHubCreatedAt(createdAt time.Time) string {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	now := time.Now()
	if createdAt.After(now.AddDate(0, 0, -1)) {
		diffHours := int(now.Sub(createdAt).Hours())
		unit := "hours"
		if diffHours == 1 {
			unit = "hour"
		}
		return fmt.Sprintf("%d %s ago", diffHours, unit)
	}
	if createdAt.After(now.AddDate(0, -1, 0)) {
		diffDays := int(now.Sub(createdAt).Hours() / 24)
		unit := "days"
		if diffDays == 1 {
			unit = "day"
		}
		return fmt.Sprintf("%d %s ago", diffDays, unit)
	}
	if createdAt.After(now.AddDate(-1, 0, 0)) {
		return createdAt.Format("2 Jan")
	}
	return createdAt.Format("2 Jan 2006")
}
