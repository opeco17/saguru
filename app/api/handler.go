package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func getRepositories(c echo.Context) error {
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
	issueIDs := fetchIssueIDs(gormDB, input)
	repositoryIDs := fetchRepositoryIDs(gormDB, input, issueIDs)
	repositories := fetchRepositoryEntities(gormDB, issueIDs, repositoryIDs)
	getRepositoriesOutput := convertGetRepositoriesOutput(repositories)
	logrus.Info(fmt.Sprintf("Total time to fetch repositories: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getRepositoriesOutput)
}

func getLanguages(c echo.Context) error {
	logrus.Info("Get languages")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLanguages := fetchFrontLanguages(gormDB)
	getLanguagesOutput := convertGetLanguagesOutput(frontLanguages)
	logrus.Info(fmt.Sprintf("Total time to fetch front languages: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLanguagesOutput)
}

func getLicenses(c echo.Context) error {
	logrus.Info("Get licenses")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLicenses := fetchFrontLicenses(gormDB)
	getLicensesOutput := convertGetLicensesOutput(frontLicenses)
	logrus.Info(fmt.Sprintf("Total time to fetch front licenses: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLicensesOutput)
}

func getLabels(c echo.Context) error {
	logrus.Info("Get labels")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLabels := fetchFrontLabels(gormDB)
	getLabelsOutput := convertGetLabelsOutput(frontLabels)
	logrus.Info(fmt.Sprintf("Total time to fetch front labels: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLabelsOutput)
}
