package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func GetRepositories(c echo.Context) error {
	logrus.Info("Get repositories")

	// Validation
	input := new(GetRepositoriesInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := input.validator(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	useIssueIDs := false
	issueIDs := []uint{}
	logrus.Info(input.Labels)
	if input.Labels != "" || input.Assigned != nil {
		useIssueIDs = true
		issueIDs = FetchIssueIDs(gormDB, input)
	}
	repositoryIDs := FetchRepositoryIDs(gormDB, input, useIssueIDs, issueIDs)
	repositories := FetchRepositoryEntities(gormDB, input, useIssueIDs, issueIDs, repositoryIDs)
	getRepositoriesOutput := convertGetRepositoriesOutput(repositories)
	logrus.Info(fmt.Sprintf("Total time to fetch repositories: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getRepositoriesOutput)
}

func GetLanguages(c echo.Context) error {
	logrus.Info("Get languages")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLanguages := FetchFrontLanguages(gormDB)
	getLanguagesOutput := convertGetLanguagesOutput(frontLanguages)
	logrus.Info(fmt.Sprintf("Total time to fetch front languages: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLanguagesOutput)
}

func GetLicenses(c echo.Context) error {
	logrus.Info("Get licenses")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLicenses := FetchFrontLicenses(gormDB)
	getLicensesOutput := convertGetLicensesOutput(frontLicenses)
	logrus.Info(fmt.Sprintf("Total time to fetch front licenses: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLicensesOutput)
}

func GetLabels(c echo.Context) error {
	logrus.Info("Get labels")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLabels := FetchFrontLabels(gormDB)
	getLabelsOutput := convertGetLabelsOutput(frontLabels)
	logrus.Info(fmt.Sprintf("Total time to fetch front labels: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLabelsOutput)
}

func GetOrderMetrics(c echo.Context) error {
	logrus.Info("Get order metrics")
	metrics := []string{}
	for _, metric := range orderMetrics() {
		metrics = append(metrics, metric+"_desc", metric+"_asc")
	}
	getOrderMetricsOutput := GetOrderMetricsOutput{Items: metrics}
	return c.JSON(http.StatusOK, getOrderMetricsOutput)
}
