package mongodb

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitMongoDB(client *mongo.Client) error {
	client.Database("main").CreateCollection(context.TODO(), REPOSITORY_COLLECTION_NAME)

	var validator = bson.M{
		"$jsonSchema": RepositorySchema,
	}
	command := bson.D{{Key: "collMod", Value: REPOSITORY_COLLECTION_NAME}, {Key: "validator", Value: validator}}
	err := client.Database("main").RunCommand(context.TODO(), command).Err()
	if err != nil {
		logrus.Error(err)
	}
	return nil
}
