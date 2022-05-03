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
	queries := [...]string{
		"stars:30..100",
		"stars:100..200",
		"stars:200..300",
		"stars:300..400",
		"stars:400..600",
		"stars:600..1000",
		"stars:1000..2000",
		"stars:2000..3000",
		"stars:4000..5000",
		"stars:>5000",
	}
	for _, eachQuery := range queries {
		now := time.Now()
		repositories := fetchRepositories(
			eachQuery,
			"good-first-issues:>0",
			"pushed:>"+threeMonthAgo,
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

		opts := options.Find().SetLimit(int64(removeRepositoryCount)).SetProjection(bson.M{"_id": 1}).SetSort(bson.M{"git_hub_updated_at": 1})
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

	// Connect DB
	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	repositoryCollection := client.Database("main").Collection("repositories")

	// Update issues
	for i := 0; i < UPDATE_ISSUE_BATCH_SIZE/UPDATE_ISSUE_MINIBATCH_SIZE; i++ {
		if i%10 == 0 {
			logrus.Info(fmt.Sprintf("Updating issues: %d/%d", i*UPDATE_ISSUE_MINIBATCH_SIZE, UPDATE_ISSUE_BATCH_SIZE))
		}
		if err := updateIssuesMinibach(repositoryCollection, UPDATE_ISSUE_MINIBATCH_SIZE); err != nil {
			return err
		}
	}

	logrus.Info("Successfully finished to update issues.")
	return nil
}

func updateIssuesMinibach(repositoryCollection *mongo.Collection, updateCount int) error {
	type SimpleRepository struct {
		RepositoryID int    `bson:"repository_id,omitempty"`
		Name         string `bson:"name,omitempty"`
	}
	type RepositoryDiff struct {
		UpdatedAt        time.Time    `bson:"updated_at"`
		IssueInitialized bool         `bson:"issue_initialized"`
		Issues           []*lib.Issue `bson:"issues"`
	}
	simpleRepoFields := bson.M{"repository_id": 1, "name": 1}

	// Fetch not initialized repositories
	var notInitializedRepositories []*SimpleRepository
	opts := options.Find().SetLimit(int64(updateCount)).SetSort(bson.M{"star_count": -1}).SetProjection(simpleRepoFields)
	filter := bson.M{"issue_initialized": false}
	initializedIssueCursor, err := repositoryCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if err = initializedIssueCursor.All(context.TODO(), &notInitializedRepositories); err != nil {
		logrus.Error(err)
		return err
	}

	// Fetch initialized repositories
	var initializedRepositories []*SimpleRepository
	if restRepositoryCount := updateCount - len(notInitializedRepositories); restRepositoryCount > 0 {
		opts = options.Find().SetLimit(int64(updateCount)).SetSort(bson.M{"star_count": -1}).SetProjection(simpleRepoFields)
		filter = bson.M{"issue_initialized": true}
		notInitializedIssueCursor, err := repositoryCollection.Find(context.TODO(), filter, opts)
		if err != nil {
			logrus.Error(err)
			return err
		}
		if err = notInitializedIssueCursor.All(context.TODO(), &initializedRepositories); err != nil {
			logrus.Error(err)
			return err
		}
	}

	// Concat repositories
	repositories := append(notInitializedRepositories, initializedRepositories...)

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

			filter := bson.M{"repository_id": repository.RepositoryID}
			update := bson.M{"$set": diff}
			_, err = repositoryCollection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				logrus.Warn(err)
			}
		}(repository)
	}
	wg.Wait()
	return nil
}
