package main

import (
	"context"
	"fmt"
	"opeco17/gitnavi/lib"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AggregationResult struct {
	ID    string `bson:"_id,omitempty"`
	Count int    `bson:"count,omitempty"`
}

func updateCache() error {
	if err := updateCachedLanguages(); err != nil {
		logrus.Error("Failed to update cached languages.")
		return err
	}
	if err := updateCachedLicenses(); err != nil {
		logrus.Error("Failed to update cached licenses.")
		return err
	}
	if err := updateCachedLabels(); err != nil {
		logrus.Error("Failed to update cached labels.")
		return err
	}
	return nil
}

func updateCachedLanguages() error {
	logrus.Info("Start updating cached languages.")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())

	// Get languages cursor
	repositoryCollection := client.Database("main").Collection("repositories")
	cacheCollection := client.Database("main").Collection("cached_languages")

	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$language", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer cursor.Close(context.TODO())

	// Update cached languages
	if err = updateCachedItems(cacheCollection, cursor); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to update cached languages.")
	return nil
}

func updateCachedLicenses() error {
	logrus.Info("Start updating cached licenses.")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())

	// Get licenses cursor
	repositoryCollection := client.Database("main").Collection("repositories")
	cacheCollection := client.Database("main").Collection("cached_licenses")

	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$license", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer cursor.Close(context.TODO())

	// Update cached licenses
	if err = updateCachedItems(cacheCollection, cursor); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to update cached licenses.")
	return nil
}

func updateCachedLabels() error {
	logrus.Info("Start updating cached labels.")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())

	// Get labels cursor
	repositoryCollection := client.Database("main").Collection("repositories")
	cacheCollection := client.Database("main").Collection("cached_labels")

	unwindStage1 := bson.D{{Key: "$unwind", Value: "$issues"}}
	unwindStage2 := bson.D{{Key: "$unwind", Value: "$issues.labels"}}
	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$issues.labels.name", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{unwindStage1, unwindStage2, groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer cursor.Close(context.TODO())

	// Update cached labels
	if err = updateCachedItems(cacheCollection, cursor); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully finished to update cached labels.")
	return nil
}

func updateCachedItems(collection *mongo.Collection, cursor *mongo.Cursor) error {
	// Get items
	var cachedItems = make([]lib.CachedItem, 0)
	for cursor.Next(context.TODO()) {
		var result AggregationResult
		if err := cursor.Decode(&result); err != nil {
			logrus.Warn(err)
		}
		item := lib.CachedItem{Name: result.ID, Count: result.Count}
		if item.Name == "" {
			continue
		}
		if containsCachedItem(cachedItems, item) {
			continue
		}
		cachedItems = append(cachedItems, item)
	}

	// Update cached items
	var cachedDocs = make([]interface{}, 0)
	for _, cachedItem := range cachedItems {
		cachedDocs = append(cachedDocs, cachedItem)
	}
	collection.Drop(context.TODO())
	_, err := collection.InsertMany(context.TODO(), cachedDocs)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func containsCachedItem(items []lib.CachedItem, target lib.CachedItem) bool {
	for _, item := range items {
		if target.Name == item.Name {
			return true
		}
	}
	return false
}
