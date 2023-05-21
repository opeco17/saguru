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

func GetLabels(c echo.Context) error {
	logrus.Info("Get labels")

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
	cachedLabels, err := service.GetCachedLabels(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get licenses from database.")
	}
	getLabelsOutput := convertGetLabelsOutput(cachedLabels)
	logrus.Info(fmt.Sprintf("Total time to fetch front labels: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, getLabelsOutput)
}

func convertGetLabelsOutput(cachedLabels []libModel.CachedItem) model.GetLabelsOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	outputItems := make([]string, 0, len(cachedLabels))
	for _, cachedLabel := range cachedLabels {
		outputItems = append(outputItems, cachedLabel.Name)
	}
	return model.GetLabelsOutput{Items: outputItems}
}
