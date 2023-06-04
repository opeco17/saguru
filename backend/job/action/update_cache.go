package action

import (
	"context"
	"fmt"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/mongodb"

	"github.com/sirupsen/logrus"
)

type AggregationResult struct {
	ID    string `bson:"_id,omitempty"`
	Count int    `bson:"count,omitempty"`
}

func UpdateCache() error {
	if err := cacheLanguages(); err != nil {
		logrus.Error("Failed to cache languages.")
		return err
	}
	if err := cacheLicenses(); err != nil {
		logrus.Error("Failed to cache licenses.")
		return err
	}
	if err := cacheLabels(); err != nil {
		logrus.Error("Failed to cache labels.")
		return err
	}
	return nil
}

func cacheLanguages() error {
	logrus.Info("Start to cache languages.")

	// Connect MongoDB
	mongoDBClient, err := util.GetMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer mongoDBClient.Disconnect(context.TODO())

	// Connect Memcached
	memcachedClinet, err := util.GetMemcachedClient()
	if err != nil {
		message := "Failed to connect to Memcached"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer memcachedClinet.Close()

	// Get languages
	languages, err := mongodb.AggregateLanguages(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate languages"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache languages
	memcachedLanguages := languages.ConvertToMemcachedLanguages()
	if err := memcachedLanguages.Save(memcachedClinet); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache languages.")
	return nil
}

func cacheLicenses() error {
	logrus.Info("Start to cache licenses.")

	// Connect MongoDB
	mongoDBClient, err := util.GetMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer mongoDBClient.Disconnect(context.TODO())

	// Connect Memcached
	memcachedClinet, err := util.GetMemcachedClient()
	if err != nil {
		message := "Failed to connect to Memcached"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer memcachedClinet.Close()

	// Get licenses
	licenses, err := mongodb.AggregateLicenses(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate licenses"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache licenses
	memcachedLicenses := licenses.ConvertToMemcachedLicenses()
	if err := memcachedLicenses.Save(memcachedClinet); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache licenses.")
	return nil
}

func cacheLabels() error {
	logrus.Info("Start to cache labels.")

	// Connect MongoDB
	mongoDBClient, err := util.GetMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer mongoDBClient.Disconnect(context.TODO())

	// Connect Memcached
	memcachedClinet, err := util.GetMemcachedClient()
	if err != nil {
		message := "Failed to connect to Memcached"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer memcachedClinet.Close()

	// Get labels
	labels, err := mongodb.AggregateLabels(mongoDBClient)
	if err != nil {
		message := "Failed to aggregate labels"
		logrus.Error(message)
		return fmt.Errorf(message)
	}

	// Cache labels
	memcachedLabels := labels.ConvertToMemcachedLabels()
	if err := memcachedLabels.Save(memcachedClinet); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to cache labels.")
	return nil
}
