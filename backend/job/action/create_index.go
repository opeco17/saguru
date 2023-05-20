package action

import (
	"context"
	"opeco17/saguru/job/util"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndex() error {
	logrus.Info("Start creating index.")

	client, err := util.GetMongoDBClient()
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
