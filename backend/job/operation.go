package main

import (
	"context"
	"opeco17/gitnavi/lib"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() error {
	logrus.Info("Start initializing DB.")

	client, err := getMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	client.Database("main").CreateCollection(context.TODO(), "repositories")

	var validator = bson.M{
		"$jsonSchema": lib.MongoSchema,
	}
	command := bson.D{{Key: "collMod", Value: "repositories"}, {Key: "validator", Value: validator}}
	err = client.Database("main").RunCommand(context.TODO(), command).Err()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Finished to initialize DB.")
	return nil
}

func createIndex() error {
	logrus.Info("Start creating index.")

	client, err := getMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("main").Collection("repositories")
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"language": 1},
		},
		{
			Keys: bson.M{"issues.labels.name": 1},
		},
		{
			Keys: bson.M{"star_count": -1},
		},
	}
	opts := options.CreateIndexes().SetMaxTime(100 * time.Second)
	_, err = collection.Indexes().CreateMany(context.TODO(), indexes, opts)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Finished to create DB.")
	return nil
}
