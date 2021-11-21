package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"opeco17/oss-book/lib"

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
	gormDB, sqlDB, err := lib.GetDBClient(
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "test")
	}
	defer sqlDB.Close()

	// Fetch data
	now := time.Now()
	issueIDs := fetchIssueIDs(gormDB, input)
	repositoryIDs := fetchRepositoryIDs(gormDB, input, issueIDs)
	repositories := fetchRepositoryEntities(c, gormDB, issueIDs, repositoryIDs)
	fmt.Printf("Total time: %vms\n", time.Since(now).Milliseconds())

	// Format data

	return c.JSON(http.StatusOK, repositories)
}
