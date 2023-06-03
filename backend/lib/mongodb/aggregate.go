package mongodb

import (
	"context"
	"opeco17/saguru/lib/memcached"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type aggregationResult struct {
	ID    string `bson:"_id,omitempty"`
	Count int    `bson:"count,omitempty"`
}

type AggregatedLanguage struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type AggregatedLanguages struct {
	Items []AggregatedLanguage `json:"items"`
}

type AggregatedLicense struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type AggregatedLicenses struct {
	Items []AggregatedLicense `json:"items"`
}

type AggregatedLabel struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type AggregatedLabels struct {
	Items []AggregatedLabel `json:"items"`
}

func (languages *AggregatedLanguages) ConvertToMemcachedLanguages() *memcached.Languages {
	memcachedLanguages := new(memcached.Languages)
	for _, language := range languages.Items {
		memcachedLanguages.Items = append(memcachedLanguages.Items, memcached.Language(language))
	}
	return memcachedLanguages
}

func (licenses *AggregatedLicenses) ConvertToMemcachedLicenses() *memcached.Licenses {
	memcachedLicenses := new(memcached.Licenses)
	for _, license := range licenses.Items {
		memcachedLicenses.Items = append(memcachedLicenses.Items, memcached.License(license))
	}
	return memcachedLicenses
}

func (labels *AggregatedLabels) ConvertToMemcachedLabels() *memcached.Labels {
	memcachedLabels := new(memcached.Labels)
	for _, label := range labels.Items {
		memcachedLabels.Items = append(memcachedLabels.Items, memcached.Label(label))
	}
	return memcachedLabels
}

func AggregateLanguages(client *mongo.Client) (*AggregatedLanguages, error) {
	repositoryCollection := client.Database(DATABASE_NAME).Collection(REPOSITORY_COLLECTION_NAME)

	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$language", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	languages := new(AggregatedLanguages)
	var result aggregationResult
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			logrus.Error(err)
			return nil, err
		}
		languages.Items = append(languages.Items, AggregatedLanguage{Name: result.ID, Count: result.Count})
	}
	return languages, nil
}

func AggregateLicenses(client *mongo.Client) (*AggregatedLicenses, error) {
	repositoryCollection := client.Database(DATABASE_NAME).Collection(REPOSITORY_COLLECTION_NAME)

	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$license", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	licenses := new(AggregatedLicenses)
	var result aggregationResult
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			logrus.Error(err)
			return nil, err
		}
		licenses.Items = append(licenses.Items, AggregatedLicense{Name: result.ID, Count: result.Count})
	}
	return licenses, nil
}

func AggregateLabels(client *mongo.Client) (*AggregatedLabels, error) {
	repositoryCollection := client.Database(DATABASE_NAME).Collection(REPOSITORY_COLLECTION_NAME)

	unwindStage1 := bson.D{{Key: "$unwind", Value: "$issues"}}
	unwindStage2 := bson.D{{Key: "$unwind", Value: "$issues.labels"}}
	groupStage := bson.D{{Key: "$group", Value: bson.M{"_id": "$issues.labels.name", "count": bson.M{"$sum": 1}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.M{"count": -1}}}
	cursor, err := repositoryCollection.Aggregate(context.TODO(), mongo.Pipeline{unwindStage1, unwindStage2, groupStage, sortStage})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	labels := new(AggregatedLabels)
	var result aggregationResult
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			logrus.Error(err)
			return nil, err
		}
		labels.Items = append(labels.Items, AggregatedLabel{Name: result.ID, Count: result.Count})
	}
	return labels, nil
}
