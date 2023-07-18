package update

import (
	"fmt"
	"opeco17/saguru/lib/mongodb"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type AggregationResult struct {
	ID    string `bson:"_id,omitempty"`
	Count int    `bson:"count,omitempty"`
}

func UpdateCaches(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start updating caches")

	if err := cacheLanguages(mongoDBClient, memcachedClient); err != nil {
		logrus.Error("Failed to cache languages")
		return err
	}
	if err := cacheLicenses(mongoDBClient, memcachedClient); err != nil {
		logrus.Error("Failed to cache licenses")
		return err
	}
	if err := cacheLabels(mongoDBClient, memcachedClient); err != nil {
		logrus.Error("Failed to cache labels")
		return err
	}

	logrus.Info("Finished updating caches")

	return nil
}

func cacheLanguages(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache languages")

	// Get languages
	languages, err := mongodb.AggregateLanguages(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate languages"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache languages
	memcachedLanguages := languages.ConvertToMemcachedLanguages()
	if err := memcachedLanguages.Save(memcachedClient); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache languages.")
	return nil
}

func cacheLicenses(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache licenses")

	// Get licenses
	licenses, err := mongodb.AggregateLicenses(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate licenses"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache licenses
	memcachedLicenses := licenses.ConvertToMemcachedLicenses()
	if err := memcachedLicenses.Save(memcachedClient); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache licenses.")
	return nil
}

func cacheLabels(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache labels")

	// Get labels
	labels, err := mongodb.AggregateLabels(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate labels"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache labels
	memcachedLabels := labels.ConvertToMemcachedLabels()
	if err := memcachedLabels.Save(memcachedClient); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache labels")
	return nil
}
