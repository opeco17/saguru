package update

import (
	"context"
	"opeco17/saguru/lib/mongodb"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateIndices(client *mongo.Client) error {
	logrus.Info("Start updating indices")

	collection := client.Database(mongodb.DATABASE_NAME).Collection(mongodb.REPOSITORY_COLLECTION_NAME)
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
	if _, err := collection.Indexes().CreateMany(context.Background(), indexes, opts); err != nil {
		return err
	}

	logrus.Info("Finished updating indices.")
	return nil
}
