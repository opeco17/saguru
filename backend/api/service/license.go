package service

import (
	"context"
	"opeco17/saguru/api/constant"
	libModel "opeco17/saguru/lib/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCachedLicenses(client *mongo.Client) ([]libModel.CachedItem, error) {
	cacheCollection := client.Database("main").Collection("cached_licenses")
	filter := bson.M{"count": bson.M{"$gte": constant.MINIMUM_COUNT_IN_CACHED_LICENSES}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLicenses []libModel.CachedItem
	if err = cursor.All(context.TODO(), &cachedLicenses); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLicenses, nil
}
