package service

import (
	"context"
	"opeco17/saguru/api/constant"
	libModel "opeco17/saguru/lib/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCachedLanguagesFromDB(client *mongo.Client) ([]libModel.CachedItem, error) {
	cacheCollection := client.Database("main").Collection("cached_languages")
	filter := bson.M{"count": bson.M{"$gte": constant.MINIMUM_COUNT_IN_CACHED_LANGUAGES}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLanguages []libModel.CachedItem
	if err = cursor.All(context.TODO(), &cachedLanguages); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLanguages, nil
}
