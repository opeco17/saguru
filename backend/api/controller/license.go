package controller

import (
	"context"
	"fmt"
	"net/http"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/service"
	"opeco17/saguru/api/util"
	libModel "opeco17/saguru/lib/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetLicenses(c echo.Context) error {
	logrus.Info("Get licenses")

	// Connect DB
	client, err := util.GetMongoDBClient()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to connect database.")
	}
	defer client.Disconnect(context.TODO())

	// Get data
	now := time.Now()
	cachedLicenses, err := service.GetCachedLicenses(client)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "Failed to get licenses from database.")
	}
	output := convertGetLicensesOutput(cachedLicenses)
	logrus.Info(fmt.Sprintf("Total time to get cached licenses: %vms\n", time.Since(now).Milliseconds()))

	return c.JSON(http.StatusOK, output)
}

func convertGetLicensesOutput(cachedLicenses []libModel.CachedItem) model.GetLicensesOutput {
	outputItems := make([]string, 0, len(cachedLicenses))
	for _, cachedLicense := range cachedLicenses {
		outputItems = append(outputItems, cachedLicense.Name)
	}
	return model.GetLicensesOutput{Items: outputItems}
}
