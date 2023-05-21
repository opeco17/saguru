package controller

import (
	"context"
	"fmt"
	"net/http"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/service"
	"opeco17/saguru/api/util"
	libModel "opeco17/saguru/lib/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetLanguages(c echo.Context) error {
	logrus.Info("Get languages")

	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	// Connect DB
	client, err := util.GetMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	cachedLanguages, err := service.GetCachedLanguagesFromDB(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get languages from database.")
	}
	output := convertGetLanguagesOutput(cachedLanguages)
	logrus.Info(fmt.Sprintf("Total time to get cached languages: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func convertGetLanguagesOutput(cachedLanguages []libModel.CachedItem) model.GetLanguagesOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	outputItems := make([]string, 0, len(cachedLanguages))
	for _, cachedLanguage := range cachedLanguages {
		outputItems = append(outputItems, cachedLanguage.Name)
	}
	return model.GetLanguagesOutput{Items: outputItems}
}
