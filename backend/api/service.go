package main

import (
	"context"
	"fmt"
	"opeco17/gitnavi/lib"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func FetchIssueIDs(gormDB *gorm.DB, input *GetRepositoriesInput) []uint {
// 	var (
// 		issues   []*lib.Issue
// 		issueIDs []uint
// 	)
// 	query := gormDB.Model(&lib.Issue{})
// 	query.Joins("INNER JOIN labels ON labels.issue_id = issues.id")
// 	if input.Labels != "" {
// 		query.Where("labels.name IN ?", strings.Split(input.Labels, ","))
// 	}
// 	if input.Assigned != nil && *input.Assigned {
// 		query.Where("issues.assignees_count > ?", 0)
// 	} else if input.Assigned != nil && !*input.Assigned {
// 		query.Where("issues.assignees_count = ?", 0)
// 	}
// 	query.Distinct("issues.id")
// 	query.Find(&issues)
// 	for _, issue := range issues {
// 		issueIDs = append(issueIDs, issue.ID)
// 	}
// 	logrus.Info(fmt.Sprintf("Issue: %d record\n", len(issues)))
// 	return issueIDs
// }

// func FetchRepositoryIDs(gormDB *gorm.DB, input *GetRepositoriesInput, useIssueIDs bool, issueIDs []uint) []uint {
// 	var (
// 		repositories  []*lib.Repository
// 		repositoryIDs []uint
// 	)
// 	query := gormDB.Model(&lib.Repository{})
// 	query.Joins("INNER JOIN issues ON issues.repository_id = repositories.id")
// 	if input.Languages != "" {
// 		query.Where("repositories.language IN ?", strings.Split(input.Languages, ","))
// 	}
// 	if input.License != "" {
// 		query.Where("repositories.license = ?", input.License)
// 	}
// 	if input.StarCountLower != nil {
// 		query.Where("repositories.star_count > ?", *input.StarCountLower)
// 	}
// 	if input.StarCountUpper != nil {
// 		query.Where("repositories.star_count < ?", *input.StarCountUpper)
// 	}
// 	if input.ForkCountLower != nil {
// 		query.Where("repositories.fork_count > ?", *input.ForkCountLower)
// 	}
// 	if input.ForkCountUpper != nil {
// 		query.Where("repositories.fork_count < ?", *input.ForkCountUpper)
// 	}
// 	if useIssueIDs {
// 		query.Where("issues.id IN ?", issueIDs)
// 	}
// 	setOrderQuery(query, input.Orderby)
// 	setDistinctQuery(query, input.Orderby)
// 	query.Offset(int(input.Page) * int(RESULTS_PER_PAGE))
// 	query.Limit(int(RESULTS_PER_PAGE) + 1)
// 	query.Find(&repositories)

// 	for _, repository := range repositories {
// 		repositoryIDs = append(repositoryIDs, repository.ID)
// 	}
// 	logrus.Info(fmt.Sprintf("Repository: %d record\n", len(repositories)))
// 	return repositoryIDs
// }

// func FetchRepositoryEntities(gormDB *gorm.DB, input *GetRepositoriesInput, useIssueIDs bool, issueIDs []uint, repositoryIDs []uint) []lib.Repository {
// 	var repositories []lib.Repository

// 	query := gormDB.Model(&repositories)
// 	if useIssueIDs {
// 		query.Preload("Issues", "id IN ?", issueIDs)
// 	} else {
// 		query.Preload("Issues")
// 	}
// 	query.Preload("Issues.Labels")
// 	query.Preload("Issues.Issuer")
// 	query.Where("id IN ?", repositoryIDs)
// 	setOrderQuery(query, input.Orderby)
// 	query.Find(&repositories)

// 	for _, repository := range repositories {
// 		logrus.Info(fmt.Sprintf("%v: %v", repository.Name, len(repository.Issues)))
// 	}
// 	return repositories
// }

func getRepositoriesFromDB(client *mongo.Client, input *GetRepositoriesInput) ([]lib.Repository, error) {
	repositoryCollection := client.Database("main").Collection("repositories")
	filter := bson.M{}

	// Filter about repositories
	if input.Languages != "" {
		filter["language"] = bson.M{"$in": strings.Split(input.Languages, ",")}
	}
	if input.License != "" {
		filter["license"] = input.License
	}
	if input.StarCountLower != nil || input.StarCountUpper != nil {
		starCountFilter := bson.M{}
		if input.StarCountLower != nil {
			starCountFilter["$gte"] = input.StarCountLower
		}
		if input.StarCountUpper != nil {
			starCountFilter["$lte"] = input.StarCountUpper
		}
		filter["star_count"] = starCountFilter
	}
	if input.ForkCountLower != nil || input.ForkCountUpper != nil {
		forkCountFilter := bson.M{}
		if input.ForkCountLower != nil {
			forkCountFilter["$gte"] = input.ForkCountLower
		}
		if input.StarCountUpper != nil {
			forkCountFilter["$lte"] = input.ForkCountUpper
		}
		filter["fork_count"] = forkCountFilter
	}

	// Filter about issues
	if input.Assigned != nil && *input.Assigned {
		filter["issues.assignees_count"] = bson.M{"$gte": 1}
	} else if input.Assigned != nil && !*input.Assigned {
		filter["issues.assignees_count"] = 0
	}

	// Filter about labels
	if input.Labels != "" {
		filter["issues.labels.name"] = bson.M{"$in": strings.Split(input.Labels, ",")}
	}

	logrus.Info(fmt.Sprintf("Filter %+v", filter))

	// Set options
	opts := options.Find().SetLimit(int64(RESULTS_PER_PAGE)).SetSkip(int64(RESULTS_PER_PAGE * input.Page))
	if input.Orderby != "" {
		// TODO
	}

	cursor, err := repositoryCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var repositories []lib.Repository
	if err = cursor.All(context.TODO(), &repositories); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return repositories, nil
}

func filterRepositories(repositories []lib.Repository, input *GetRepositoriesInput) []lib.Repository {
	// TODO
	return repositories
}

func getCachedLanguagesFromDB(client *mongo.Client) ([]lib.CachedItem, error) {
	cacheCollection := client.Database("main").Collection("cached_languages")
	filter := bson.M{"count": bson.M{"$gte": MINIMUM_COUNT_IN_CACHED_LANGUAGES}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLanguages []lib.CachedItem
	if err = cursor.All(context.TODO(), &cachedLanguages); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLanguages, nil
}

func getCachedLicenses(client *mongo.Client) ([]lib.CachedItem, error) {
	cacheCollection := client.Database("main").Collection("cached_licenses")
	filter := bson.M{"count": bson.M{"$gte": MINIMUM_COUNT_IN_CACHED_LICENSES}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLicenses []lib.CachedItem
	if err = cursor.All(context.TODO(), &cachedLicenses); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLicenses, nil
}

func getCachedLabels(client *mongo.Client) ([]lib.CachedItem, error) {
	cacheCollection := client.Database("main").Collection("cached_labels")
	filter := bson.M{"count": bson.M{"$gte": MINIMUM_COUNT_IN_CACHED_LABELS}}
	cursor, err := cacheCollection.Find(context.TODO(), filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	var cachedLabels []lib.CachedItem
	if err = cursor.All(context.TODO(), &cachedLabels); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cachedLabels, nil
}
