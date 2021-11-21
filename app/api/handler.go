package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func getRepositories(c echo.Context) error {
	// Validation
	input := new(getRepositoriesInput)
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
	fmt.Printf("Total time to fetch repositories: %vms\n", time.Since(now).Milliseconds())

	// Format data

	return c.JSON(http.StatusOK, repositories)
}

func getLanguages(c echo.Context) error {
	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	frontLanguages := fetchFrontLanguages(gormDB)
	fmt.Printf("Total time to fetch front languages: %vms\n", time.Since(now).Milliseconds())

	// Format data

	return c.JSON(http.StatusOK, frontLanguages)
}
