package action

import (
	"context"
	"opeco17/saguru/job/util"
	"opeco17/saguru/lib/schema"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func InitDB() error {
	logrus.Info("Start initializing DB.")

	client, err := util.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	client.Database("main").CreateCollection(context.TODO(), "repositories")

	var validator = bson.M{
		"$jsonSchema": schema.RepositorySchema,
	}
	command := bson.D{{Key: "collMod", Value: "repositories"}, {Key: "validator", Value: validator}}
	err = client.Database("main").RunCommand(context.TODO(), command).Err()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Finished to initialize DB.")
	return nil
}
