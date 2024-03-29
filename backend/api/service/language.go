package service

import (
	"opeco17/saguru/api/metrics"
	errorsutil "opeco17/saguru/lib/errors"
	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLanguagesFromMemcached(client *memcache.Client) (*memcached.Languages, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	languages, err := memcached.GetLanguages(client)
	if err != nil {
		return nil, errorsutil.Wrap(err, "Failed to get languages from Memcached")
	}
	return languages, nil
}

func GetLanguagesFromMongoDB(client *mongo.Client) (*memcached.Languages, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	languages, err := mongodb.AggregateLanguages(client)
	if err != nil {
		return nil, errorsutil.Wrap(err, "Failed to get languages from MongoDB")
	}
	return languages.ConvertToMemcachedLanguages(), nil
}
