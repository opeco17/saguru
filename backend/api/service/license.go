package service

import (
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLicensesFromMemcached(client *memcache.Client) (*memcached.Licenses, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	licenses, err := memcached.GetLicenses(client)
	if err != nil {
		return nil, err
	}
	return licenses, nil
}

func GetLicensesFromMongoDB(client *mongo.Client) (*memcached.Licenses, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	licenses, err := mongodb.AggregateLicenses(client)
	if err != nil {
		return nil, err
	}
	return licenses.ConvertToMemcachedLicenses(), nil
}
