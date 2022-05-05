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
	if input.Keyword != "" {
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{"$regex": input.Keyword, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": input.Keyword, "$options": "i"}},
		}
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
	issueFilter := bson.M{"assignees_count": bson.M{"$gte": 0}} // To remove empty issues
	if input.IsAssigned != nil && *input.IsAssigned {
		issueFilter["assignees_count"] = bson.M{"$gte": 1}
	} else if input.IsAssigned != nil && !*input.IsAssigned {
		issueFilter["assignees_count"] = 0
	}
	if input.Labels != "" {
		issueFilter["labels.name"] = bson.M{"$in": strings.Split(input.Labels, ",")}
	}
	filter["issues"] = bson.M{"$elemMatch": issueFilter, "$exists": true}

	logrus.Info(fmt.Sprintf("Filter %+v", filter))

	// Set options
	var metric string
	var direction int
	orderBy := input.Orderby
	if orderBy == "" {
		orderBy = "STAR_COUNT_DESC"
	}
	if strings.Contains(orderBy, "STAR_COUNT") {
		metric = "star_count"
	} else if strings.Contains(orderBy, "FORK_COUNT") {
		metric = "fork_count"
	}
	if strings.Contains(orderBy, "DESC") {
		direction = -1
	} else if strings.Contains(orderBy, "ASC") {
		direction = 1
	}
	opts := options.Find().SetLimit(int64(RESULTS_PER_PAGE + 1)).SetSkip(int64(RESULTS_PER_PAGE * input.Page)).SetSort(bson.M{metric: direction})

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

func filterIssuesInRepositories(repositories []lib.Repository, input *GetRepositoriesInput) []lib.Repository {
	filteredRepositories := make([]lib.Repository, 0, len(repositories))
	assigneeFilter := func(assigneesCount int) bool { return true }
	labelFilter := func(labels []*lib.Label) bool { return true }

	// Set filter
	if input.IsAssigned != nil && *input.IsAssigned {
		assigneeFilter = func(assigneesCount int) bool { return assigneesCount > 0 }
	} else if input.IsAssigned != nil && !*input.IsAssigned {
		assigneeFilter = func(assigneesCount int) bool { return assigneesCount == 0 }
	}

	if input.Labels != "" {
		labelFilter = func(labels []*lib.Label) bool {
			inputLabelNames := strings.Split(input.Labels, ",")
			for _, label := range labels {
				for _, inputLabelName := range inputLabelNames {
					if label.Name == inputLabelName {
						return true
					}
				}
			}
			return false
		}
	}

	// Filter issues
	for _, repository := range repositories {
		filteredIssues := make([]*lib.Issue, 0, len(repository.Issues))
		for _, issue := range repository.Issues {
			if assigneeFilter(*issue.AssigneesCount) && labelFilter(issue.Labels) {
				filteredIssues = append(filteredIssues, issue)
			}
		}
		repository.Issues = filteredIssues
		filteredRepositories = append(filteredRepositories, repository)
	}
	return filteredRepositories
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
