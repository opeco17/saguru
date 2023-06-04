package service

import (
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/lib/memcached"
	"opeco17/saguru/lib/mongodb"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLabelsFromMemcached(client *memcache.Client) (*memcached.Labels, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	labels, err := memcached.GetLabels(client)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

func GetLabelsFromMongoDB(client *mongo.Client) (*memcached.Labels, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	labels, err := mongodb.AggregateLabels(client)
	if err != nil {
		return nil, err
	}
	return labels.ConvertToMemcachedLabels(), nil
}
