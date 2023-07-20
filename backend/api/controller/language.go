package controller

import (
	"net/http"
	"opeco17/saguru/api/constant"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/service"
	"opeco17/saguru/api/util"
	errorsutil "opeco17/saguru/lib/errors"
	"opeco17/saguru/lib/memcached"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetLanguages(c echo.Context) error {
	logrus.Info("Get languages")

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
	languages := new(memcached.Languages)
	if connectedToMemcached {
		languages, err = service.GetLanguagesFromMemcached(memcachedClient)
		if err != nil {
			hitCache = false
		}
		metrics.M.CountCacheAccess(memcached.LANGUAGES_CACHE_KEY, hitCache)
	}

	if !(connectedToMemcached && hitCache) {
		mongoDBClient, err := util.GetMongoDBClient()
		if err != nil {
			logrus.Error("Failed to connect to MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get languages")
		}

		languages, err = service.GetLanguagesFromMongoDB(mongoDBClient)
		if err != nil {
			logrus.Error("Failed to get labels from MongoDB")
			logrus.Errorf("%#v", err)
			return c.String(http.StatusServiceUnavailable, "Failed to get languages")
		}
	}

	if connectedToMemcached && !hitCache {
		languages.Save(memcachedClient)
	}

	output := convertGetLanguagesOutput(languages)
	if err := c.JSON(http.StatusOK, output); err != nil {
		logrus.Errorf("%#v", errorsutil.Wrap(err, err.Error()))
		return c.String(http.StatusServiceUnavailable, "Something wrong happend")
	}
	return nil
}

func convertGetLanguagesOutput(languages *memcached.Languages) model.GetLanguagesOutput {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	items := languages.Items
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	outputItems := make([]string, 0, len(languages.Items))
	for _, language := range items {
		if language.Name == "" {
			continue
		}
		if language.Count <= constant.MINIMUM_COUNT_IN_CACHED_LANGUAGES {
			continue
		}
		outputItems = append(outputItems, language.Name)
	}
	return model.GetLanguagesOutput{Items: outputItems}
}
