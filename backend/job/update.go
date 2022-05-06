package main

import (
	"context"
	"fmt"
	"opeco17/gitnavi/lib"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SimpleRepository struct {
	Name string `bson:"name,omitempty"`
}

type RepositoryDiff struct {
	UpdatedAt        time.Time    `bson:"updated_at"`
	IssueInitialized bool         `bson:"issue_initialized"`
	Issues           []*lib.Issue `bson:"issues"`
}

func updateRepositories() error {
	logrus.Info("Start updating repositories.")

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	repositoryCollection := client.Database("main").Collection("repositories")

	// Fetch and save repositories
	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
	pushedBeforeThreeMonthAgoQuery := "pushed:>" + threeMonthAgo
	isGoodFirstIssueQuery := "good-first-issues:>0"
	queries := [][]string{
		{"stars:30..50", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:30..50", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:50..100", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:100..150", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:150..200", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:200..300", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:300..400", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:400..600", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:600..1000", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1000..1300", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1300..1500", pushedBeforeThreeMonthAgoQuery, isGoodFirstIssueQuery},
		{"stars:1500..1700", pushedBeforeThreeMonthAgoQuery},
		{"stars:1700..2000", pushedBeforeThreeMonthAgoQuery},
		{"stars:2000..2500", pushedBeforeThreeMonthAgoQuery},
		{"stars:2500..3000", pushedBeforeThreeMonthAgoQuery},
		{"stars:3500..4000", pushedBeforeThreeMonthAgoQuery},
		{"stars:4000..5000", pushedBeforeThreeMonthAgoQuery},
		{"stars:5000..6000", pushedBeforeThreeMonthAgoQuery},
		{"stars:6000..7000", pushedBeforeThreeMonthAgoQuery},
		{"stars:7000..10000", pushedBeforeThreeMonthAgoQuery},
		{"stars:10000..15000", pushedBeforeThreeMonthAgoQuery},
		{"stars:15000..20000", pushedBeforeThreeMonthAgoQuery},
		{"stars:>20000", pushedBeforeThreeMonthAgoQuery},
	}
	for _, eachQuery := range queries {
		now := time.Now()
		repositories := fetchRepositories(
			eachQuery...,
		)
		for _, repository := range repositories {
			repository.UpdatedAt = time.Now()

			// Set existing issues
			var oldRepository lib.Repository
			findFilter := bson.M{"repository_id": repository.RepositoryID}
			repositoryCollection.FindOne(context.TODO(), findFilter).Decode(&oldRepository)
			if oldRepository.RepositoryID != 0 {
				repository.Issues = oldRepository.Issues
			}

			// Update
			updateFilter := bson.M{"repository_id": repository.RepositoryID}
			update := bson.M{"$set": repository}
			_, err = repositoryCollection.UpdateOne(context.TODO(), updateFilter, update, options.Update().SetUpsert(true))
			if err != nil {
				logrus.Warn(err)
			}
		}
		if restTimeSecond := REPOSITORIES_API_INTERVAL_SECOND - int(time.Since(now).Seconds()); restTimeSecond > 0 {
			time.Sleep(time.Second * time.Duration(restTimeSecond))
		}
	}

	// Adjust number of repositories by removing old repositories
	repositoryCount, err := repositoryCollection.CountDocuments(context.TODO(), bson.M{})
	if removeRepositoryCount := int(repositoryCount) - MAX_REPOSITORY_RECORES; removeRepositoryCount > 0 {
		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))

		opts := options.Find()
		opts.SetLimit(int64(removeRepositoryCount))
		opts.SetProjection(bson.M{"_id": 1})
		opts.SetSort(bson.M{"git_hub_updated_at": 1})
		cursor, err := repositoryCollection.Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			logrus.Error(err)
			return err
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var result lib.Repository
			if err := cursor.Decode(&result); err != nil {
				logrus.Error(err)
				return err
			}
			repositoryCollection.DeleteOne(context.TODO(), result)
		}
	}

	logrus.Info("Successfully finished to update repositories.")
	return nil
}

func updateIssues() error {
	logrus.Info("Start updating issues.")

	remainCount := UPDATE_ISSUE_COUNT_PER_BATCH

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	repositoryCollection := client.Database("main").Collection("repositories")

	remainCount = updateIssuesInNotInitializedRepositories(repositoryCollection, remainCount)

	if remainCount > 0 {
		updateIssuesInInitializedRepositories(repositoryCollection, remainCount)
	}

	logrus.Info("Successfully finished to update issues.")
	return nil
}

func updateIssuesInInitializedRepositories(collection *mongo.Collection, remainCount int) int {
	// Fetch not initialized repositories
	opts := options.Find()
	opts.SetLimit(int64(UPDATE_ISSUE_COUNT_PER_BATCH))
	opts.SetSort(bson.M{"updated_at": 1})
	opts.SetProjection(bson.M{"name": 1})
	filter := bson.M{"issue_initialized": false}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		logrus.Error(err)
	}
	defer cursor.Close(context.TODO())

	var repository *SimpleRepository
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(repository); err != nil {
			logrus.Error(err)
		}
		remainCount = updateIssuesInRepository(repository, remainCount)
		if remainCount <= 0 {
			break
		}
	}
	return remainCount
}

func updateIssuesInNotInitializedRepositories(collection *mongo.Collection, remainCount int) int {
	opts := options.Find()
	opts.SetLimit(int64(UPDATE_ISSUE_COUNT_PER_BATCH))
	opts.SetSort(bson.M{"updated_at": 1})
	opts.SetProjection(bson.M{"name": 1})
	filter := bson.M{"issue_initialized": true}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		logrus.Error(err)
	}
	defer cursor.Close(context.TODO())

	var repository *SimpleRepository
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(repository); err != nil {
			logrus.Error(err)
		}
		remainCount = updateIssuesInRepository(repository, remainCount)
		if remainCount <= 0 {
			break
		}
	}
	return remainCount
}

func updateIssuesInRepository(repository *SimpleRepository, remainCount int) int {
	// Update issues concurrently
	var wg sync.WaitGroup
	concurrencyLimitCh := make(chan struct{}, FETCH_ISSUE_CONCURRENCY)
	wg.Add(len(repositories))

	// Fetch and update issues
	for _, repository := range repositories {
		go func(repository *SimpleRepository) {
			concurrencyLimitCh <- struct{}{}
			defer wg.Done()
			defer func() { <-concurrencyLimitCh }()

			issues := fetchIssues(repository.Name)
			diff := RepositoryDiff{
				UpdatedAt:        time.Now(),
				IssueInitialized: true,
				Issues:           issues,
			}

			filter := bson.M{"name": repository.Name}
			update := bson.M{"$set": diff}
			_, err = repositoryCollection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				logrus.Warn(err)
			}
		}(repository)
	}
	wg.Wait()
}
