package service

import (
	"opeco17/saguru/api/metrics"
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
		return nil, err
	}
	return languages, nil
}

func GetLanguagesFromMongoDB(client *mongo.Client) (*memcached.Languages, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	languages, err := mongodb.AggregateLanguages(client)
	if err != nil {
		return nil, err
	}
	return languages.ConvertToMemcachedLanguages(), nil
}
