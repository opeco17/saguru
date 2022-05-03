package main

import (
	"context"
	"opeco17/gitnavi/lib"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
