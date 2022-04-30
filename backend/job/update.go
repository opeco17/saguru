package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func UpdateRepositories() error {
// 	logrus.Info("Start updating repositories.")

// 	// Connect DB
// 	gormDB, sqlDB, err := getDBClient()
// 	if err != nil {
// 		logrus.Error(err)
// 		return fmt.Errorf("error occured when updating repositories")
// 	}
// 	defer sqlDB.Close()

// 	// Fetch and save repositories
// 	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
// 	queries := [...]string{
// 		"stars:30..100",
// 		"stars:100..200",
// 		"stars:200..400",
// 		"stars:400..600",
// 		"stars:600..1000",
// 		"stars:1000..2000",
// 		"stars:2000..4000",
// 		"stars:>4000",
// 	}
// 	for _, eachQuery := range queries {
// 		now := time.Now()
// 		repositories := fetchRepositories(
// 			eachQuery,
// 			"good-first-issues:>0",
// 			fmt.Sprintf("pushed:>%s", threeMonthAgo),
// 		)
// 		gormDB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&repositories)
// 		if restTimeSecond := REPOSITORIES_API_INTERVAL_SECOND - int(time.Since(now).Seconds()); restTimeSecond > 0 {
// 			time.Sleep(time.Second * time.Duration(restTimeSecond))
// 		}
// 	}

// 	// Adjust number of repositories by removing old repositories
// 	var (
// 		repositoryCount    int64
// 		removeRepositories []*lib.Repository
// 	)
// 	gormDB.Model(&lib.Repository{}).Count(&repositoryCount)
// 	if removeRepositoryCount := int(repositoryCount) - MAX_REPOSITORY_RECORES; removeRepositoryCount > 0 {
// 		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))
// 		gormDB.Model(&lib.Repository{}).Order("git_hub_updated_at ASC").Limit(removeRepositoryCount).Find(&removeRepositories)
// 		gormDB.Unscoped().Delete(&removeRepositories)
// 	}
// 	logrus.Info("Successfully finished to update repositories.")
// 	return nil
// }

func UpdateRepositories() error {
	logrus.Info("Start updating repositories.")

	client, err := getMongoDBClient()
	if err != nil {
		message := "Failed to connect to MongoDB"
		logrus.Error(message)
		return fmt.Errorf(message)
	}
	defer client.Disconnect(context.TODO())
	updateOpts := options.Update().SetUpsert(true)
	collection := client.Database("main").Collection("repositories")

	// Fetch and save repositories
	threeMonthAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02T15:04:05+09:00")
	queries := [...]string{
		"stars:30..100",
		"stars:100..200",
		"stars:200..400",
		"stars:400..600",
		"stars:600..1000",
		"stars:1000..2000",
		"stars:2000..4000",
		"stars:>4000",
	}
	for _, eachQuery := range queries {
		now := time.Now()
		repositories := fetchRepositories(
			eachQuery,
			"good-first-issues:>0",
			"pushed:>"+threeMonthAgo,
		)
		for _, repository := range repositories {
			filter := bson.M{"repository_id": repository.RepositoryID}
			update := bson.M{"$set": repository}
			_, err = collection.UpdateOne(context.TODO(), filter, update, updateOpts)
			if err != nil {
				logrus.Warn(err)
			}
		}
		if restTimeSecond := REPOSITORIES_API_INTERVAL_SECOND - int(time.Since(now).Seconds()); restTimeSecond > 0 {
			time.Sleep(time.Second * time.Duration(restTimeSecond))
		}
	}
	// Adjust number of repositories by removing old repositories
	repositoryCount, err := collection.CountDocuments(context.TODO(), bson.M{})
	if removeRepositoryCount := int(repositoryCount) - MAX_REPOSITORY_RECORES; removeRepositoryCount > 0 {
		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))
		// gormDB.Model(&lib.Repository{}).Order("git_hub_updated_at ASC").Limit(removeRepositoryCount).Find(&removeRepositories)
		// gormDB.Unscoped().Delete(&removeRepositories)
	}

	logrus.Info("Successfully finished to update repositories.")
	return nil
}

// func UpdateIssues() error {
// 	logrus.Info("Start updating issues.")

// 	// Connect DB
// 	gormDB, sqlDB, err := getDBClient()
// 	if err != nil {
// 		logrus.Error(err)
// 		return fmt.Errorf("error occured when updating repositories")
// 	}
// 	defer sqlDB.Close()

// 	// Update issues
// 	for i := 0; i < UPDATE_ISSUE_BATCH_SIZE/UPDATE_ISSUE_MINIBATCH_SIZE; i++ {
// 		if i%10 == 0 {
// 			logrus.Info(fmt.Sprintf("Updating issues: %d/%d", i*UPDATE_ISSUE_MINIBATCH_SIZE, UPDATE_ISSUE_BATCH_SIZE))
// 		}
// 		if err := UpdateIssuesMinibach(gormDB, UPDATE_ISSUE_MINIBATCH_SIZE); err != nil {
// 			return err
// 		}
// 	}
// 	logrus.Info("Successfully finished to update issues.")
// 	return nil
// }

// func UpdateIssuesMinibach(gormDB *gorm.DB, updateNum int) error {
// 	// Get target repositories to update issue
// 	var (
// 		notInitializedRepositories []*lib.Repository
// 		initializedRepositories    []*lib.Repository
// 	)
// 	gormDB.Where("issue_initialized = ?", false).Limit(updateNum).Find(&notInitializedRepositories)
// 	if restRepositoryNum := updateNum - len(notInitializedRepositories); restRepositoryNum > 0 {
// 		gormDB.Where("issue_initialized = ?", true).Order("updated_at ASC").Limit(restRepositoryNum).Find(&initializedRepositories)
// 	}
// 	repositories := append(notInitializedRepositories, initializedRepositories...)

// 	// Update issues concurrently
// 	var wg sync.WaitGroup
// 	concurrencyLimitCh := make(chan struct{}, FETCH_ISSUE_CONCURRENCY)
// 	wg.Add(len(repositories))

// 	// Fetch issues
// 	for _, repository := range repositories {
// 		go func(repository *lib.Repository) {
// 			concurrencyLimitCh <- struct{}{}
// 			defer wg.Done()
// 			defer func() { <-concurrencyLimitCh }()
// 			issues := fetchIssues(repository.Name)
// 			repository.Issues = issues
// 			repository.IssueInitialized = true
// 		}(repository)
// 	}
// 	wg.Wait()

// 	gormDB.Clauses(clause.OnConflict{UpdateAll: true}).Save(repositories)
// 	return nil
// }

func UpdateIssues() error {
	return nil
}

// func UpdateFrontLanguages() error {
// 	logrus.Info("Start updating languages.")

// 	// Connect DB
// 	gormDB, sqlDB, err := getDBClient()
// 	if err != nil {
// 		logrus.Error(err)
// 		return fmt.Errorf("error occured when updating front languages")
// 	}
// 	defer sqlDB.Close()

// 	// Fetch languages from other table
// 	var (
// 		languages    []*lib.FrontLanguage
// 		oldLanguages []*lib.FrontLanguage
// 	)
// 	query := gormDB.Model(&lib.Repository{})
// 	query.Select("language AS name, COUNT(language) AS repository_count")
// 	query.Where("language != ?", "")
// 	query.Group("language")
// 	query.Order("repository_count DESC")
// 	query.Find(&languages)

// 	// Update languages inside transaction
// 	gormDB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLanguages).Error; err != nil {
// 			logrus.Error("error occured when deleting old front languages")
// 			return err
// 		}
// 		if err := tx.Create(&languages).Error; err != nil {
// 			logrus.Error("error occured when inserting new front languages")
// 			return err
// 		}
// 		return nil
// 	})

// 	return nil
// }

func UpdateFrontLanguages() error {
	return nil
}

// func UpdateLicenses() error {
// 	logrus.Info("Start updating licenses.")

// 	// Connect DB
// 	gormDB, sqlDB, err := getDBClient()
// 	if err != nil {
// 		logrus.Error(err)
// 		return fmt.Errorf("error occured when updating front languages")
// 	}
// 	defer sqlDB.Close()

// 	// Fetch licenses from other table
// 	var (
// 		licenses    []*lib.FrontLicense
// 		oldLicenses []*lib.FrontLicense
// 	)
// 	query := gormDB.Model(&lib.Repository{})
// 	query.Select("license AS name, COUNT(license) AS repository_count")
// 	query.Where("license != ?", "")
// 	query.Group("license")
// 	query.Order("repository_count DESC")
// 	query.Find(&licenses)

// 	// Update licenses inside transaction
// 	gormDB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLicenses).Error; err != nil {
// 			logrus.Error("error occured when deleting old front licenses")
// 			return err
// 		}
// 		if err := tx.Create(&licenses).Error; err != nil {
// 			logrus.Error("error occured when inserting new front licenses")
// 			return err
// 		}
// 		return nil
// 	})
// 	return nil
// }

func UpdateLicenses() error {
	return nil
}

// func UpdateLabels() error {
// 	logrus.Info("Start updating labels.")

// 	// Connect DB
// 	gormDB, sqlDB, err := getDBClient()
// 	if err != nil {
// 		logrus.Error(err)
// 		return fmt.Errorf("error occured when updating front labels")
// 	}
// 	defer sqlDB.Close()

// 	// Fetch labels from other table
// 	var (
// 		labels    []*lib.FrontLabel
// 		oldLabels []*lib.FrontLabel
// 	)
// 	query := gormDB.Model(&lib.Label{})
// 	query.Select("name, COUNT(name) AS issue_count")
// 	query.Where("name != ?", "")
// 	query.Group("name")
// 	query.Order("issue_count DESC")
// 	query.Find(&labels)

// 	// Update labels inside transaction
// 	gormDB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLabels).Error; err != nil {
// 			logrus.Error("error occured when deleting old front labels")
// 			return err
// 		}
// 		if err := tx.Create(&labels).Error; err != nil {
// 			logrus.Error("error occured when inserting new front labels")
// 			return err
// 		}
// 		return nil
// 	})
// 	return nil
// }

func UpdateLabels() error {
	return nil
}
