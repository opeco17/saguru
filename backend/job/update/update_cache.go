package update

import (
	"opeco17/saguru/lib/mongodb"
	"sync"

	errorsutil "opeco17/saguru/lib/errors"

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

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := cacheLanguages(mongoDBClient, memcachedClient); err != nil {
			logrus.Error("Failed to cache languages")
			logrus.Errorf("%#v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := cacheLicenses(mongoDBClient, memcachedClient); err != nil {
			logrus.Error("Failed to cache licenses")
			logrus.Errorf("%#v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := cacheLabels(mongoDBClient, memcachedClient); err != nil {
			logrus.Error("Failed to cache labels")
			logrus.Errorf("%#v", err)
		}
	}()

	wg.Wait()

	logrus.Info("Finished updating caches")

	return nil
}

func cacheLanguages(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache languages")

	// Get languages
	languages, err := mongodb.AggregateLanguages(mongoDBClient)
	if err != nil {
		return errorsutil.Wrap(err, "Failed to get languages from MongoDB")
	}

	// Cache languages
	memcachedLanguages := languages.ConvertToMemcachedLanguages()
	if err := memcachedLanguages.Save(memcachedClient); err != nil {
		return errorsutil.Wrap(err, "Failed to cache languages to Memcached")
	}

	logrus.Info("Successfully finished to cache languages.")
	return nil
}

func cacheLicenses(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache licenses")

	// Get licenses
	licenses, err := mongodb.AggregateLicenses(mongoDBClient)
	if err != nil {
		return errorsutil.Wrap(err, "Failed to get licenses from MongoDB")
	}

	// Cache licenses
	memcachedLicenses := licenses.ConvertToMemcachedLicenses()
	if err := memcachedLicenses.Save(memcachedClient); err != nil {
		return errorsutil.Wrap(err, "Failed to cache licenses to Memcached")
	}

	logrus.Info("Successfully finished to cache licenses.")
	return nil
}

func cacheLabels(mongoDBClient *mongo.Client, memcachedClient *memcache.Client) error {
	logrus.Info("Start to cache labels")

	// Get labels
	labels, err := mongodb.AggregateLabels(mongoDBClient)
	if err != nil {
		return errorsutil.Wrap(err, "Failed to get labels from MongoDB")
	}

	// Cache labels
	memcachedLabels := labels.ConvertToMemcachedLabels()
	if err := memcachedLabels.Save(memcachedClient); err != nil {
		return errorsutil.Wrap(err, "Failed to cache labels to Memcached")
	}

	logrus.Info("Successfully finished to cache labels")
	return nil
}
