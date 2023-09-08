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

	errorsutil "opeco17/saguru/lib/errors"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetLabels(c echo.Context) error {
	logrus.Info("Get labels")

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
	labels := new(memcached.Labels)
	if connectedToMemcached {
		labels, err = service.GetLabelsFromMemcached(memcachedClient)
		if err != nil {
			hitCache = false
		}
		metrics.M.CountCacheAccess(memcached.LABELS_CACHE_KEY, hitCache)
	}

	if !(connectedToMemcached && hitCache) {
		mongoDBClient, err := util.GetMongoDBClient()
		if err != nil {
			logrus.Error("Failed to connect to MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get labels")
		}

		labels, err = service.GetLabelsFromMongoDB(mongoDBClient)
		if err != nil {
			logrus.Error("Failed to get labels from MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get labels")
		}
	}

	if connectedToMemcached && !hitCache {
		if err := labels.Save(memcachedClient); err != nil {
			logrus.Warn("Failed to save labels into Memcached")
			logrus.Warnf("%#v", err)
		}
	}

	output := convertGetLabelsOutput(labels)
	if err := c.JSON(http.StatusOK, output); err != nil {
		logrus.Errorf("%#v", errorsutil.Wrap(err, err.Error()))
		return c.String(http.StatusServiceUnavailable, "Something wrong happend")
	}
	return nil
}

func convertGetLabelsOutput(labels *memcached.Labels) model.GetLabelsOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	items := labels.Items
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	outputItems := make([]string, 0, len(labels.Items))
	for _, label := range items {
		if label.Name == "" {
			continue
		}
		if label.Count <= constant.MINIMUM_COUNT_IN_CACHED_LABELS {
			continue
		}
		outputItems = append(outputItems, label.Name)
	}
	return model.GetLabelsOutput{Items: outputItems}
}
