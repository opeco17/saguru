package service

import (
	"context"
	"opeco17/saguru/api/constant"
	"opeco17/saguru/api/metrics"
	libModel "opeco17/saguru/lib/model"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCachedLabels(client *mongo.Client) ([]libModel.CachedItem, error) {
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	cacheCollection := client.Database("main").Collection("cached_labels")
	filter := bson.M{"count": bson.M{"$gte": constant.MINIMUM_COUNT_IN_CACHED_LABELS}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLabels []libModel.CachedItem
	if err = cursor.All(context.TODO(), &cachedLabels); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLabels, nil
}
