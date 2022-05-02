package main

import (
	"context"
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
	logrus.Info(fmt.Sprintf("Input %+v", input))
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := input.validator(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	repositories, err := getRepositoriesFromDB(client, input)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get repositories from database.")
	}
	repositories = filterIssuesInRepositories(repositories, input)
	output := convertGetRepositoriesOutput(repositories)

	logrus.Info(fmt.Sprintf("Total time to fetch repositories: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func getLanguages(c echo.Context) error {
	logrus.Info("Get languages")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	cachedLanguages, err := getCachedLanguagesFromDB(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get languages from database.")
	}
	output := convertGetLanguagesOutput(cachedLanguages)
	logrus.Info(fmt.Sprintf("Total time to get cached languages: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func getLicenses(c echo.Context) error {
	logrus.Info("Get licenses")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	cachedLicenses, err := getCachedLicenses(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get licenses from database.")
	}
	output := convertGetLicensesOutput(cachedLicenses)
	logrus.Info(fmt.Sprintf("Total time to get cached licenses: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func getLabels(c echo.Context) error {
	logrus.Info("Get labels")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	cachedLabels, err := getCachedLabels(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get licenses from database.")
	}
	getLabelsOutput := convertGetLabelsOutput(cachedLabels)
	logrus.Info(fmt.Sprintf("Total time to fetch front labels: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLabelsOutput)
}

func getOrderMetrics(c echo.Context) error {
	logrus.Info("Get order metrics")

	metrics := []string{}
	for _, metric := range orderMetrics() {
		metrics = append(metrics, metric+"_desc", metric+"_asc")
	}
	getOrderMetricsOutput := GetOrderMetricsOutput{Items: metrics}
	return c.JSON(http.StatusOK, getOrderMetricsOutput)
}
