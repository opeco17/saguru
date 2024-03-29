package controller

import (
	"net/http"
	"opeco17/saguru/api/constant"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/service"
	"opeco17/saguru/api/util"
	"opeco17/saguru/lib/memcached"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetLicenses(c echo.Context) error {
	logrus.Info("Get licenses")

	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	connectedToMemcached := true
	memcachedClient, err := util.GetMemcachedClient()
	if err != nil {
		logrus.Warn("Failed to connect to Memcached")
		logrus.Warnf("%#v", err)
		connectedToMemcached = false
	}

	hitCache := true
	licenses := new(memcached.Licenses)
	if connectedToMemcached {
		licenses, err = service.GetLicensesFromMemcached(memcachedClient)
		if err != nil {
			hitCache = false
		}
		metrics.M.CountCacheAccess(memcached.LICENSES_CACHE_KEY, hitCache)
	}

	if !(connectedToMemcached && hitCache) {
		mongoDBClient, err := util.GetMongoDBClient()
		if err != nil {
			logrus.Error("Failed to connect to MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get licenses")
		}

		licenses, err = service.GetLicensesFromMongoDB(mongoDBClient)
		if err != nil {
			logrus.Error("Failed to get labels from MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get licenses")
		}
	}

	if connectedToMemcached && !hitCache {
		if err := licenses.Save(memcachedClient); err != nil {
			logrus.Warn("Failed to save licenses into Memcached")
			logrus.Warnf("%#v", err)
		}
	}

	output := convertGetLicensesOutput(licenses)
	if err := c.JSON(http.StatusOK, output); err != nil {
		return c.String(http.StatusServiceUnavailable, "Something wrong happend")
	}
	return nil
}

func convertGetLicensesOutput(licenses *memcached.Licenses) model.GetLicensesOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	items := licenses.Items
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	outputItems := make([]string, 0, len(licenses.Items))
	for _, license := range items {
		if license.Name == "" {
			continue
		}
		if license.Count <= constant.MINIMUM_COUNT_IN_CACHED_LICENSES {
			continue
		}
		outputItems = append(outputItems, license.Name)
	}
	return model.GetLicensesOutput{Items: outputItems}
}
