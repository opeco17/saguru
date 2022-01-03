package main

import (
	"fmt"
	"opeco17/gitnavi/lib"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdateRepositories() error {
	logrus.Info("Start updating repositories.")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when updating repositories")
	}
	defer sqlDB.Close()

	// Fetch and save repositories
	uniqueQuery := [...]string{
		"stars:30..100",
		// "stars:100..200",
		// "stars:200..400",
		// "stars:400..1000",
		// "stars:1000..3000",
		// "stars:>3000",
	}
	for _, eachQuery := range uniqueQuery {
		now := time.Now()
		repositories := FetchRepositories(eachQuery)
		gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&repositories)
		restTimeSecond := int(REPOSITORIES_API_INTERVAL_SECOND) - int(time.Since(now).Seconds())
		if restTimeSecond > 0 {
			time.Sleep(time.Second * time.Duration(restTimeSecond))
		}
	}

	// Adjust number of repositories by removing old repositories
	var (
		repositoryCount    int64
		removeRepositories []lib.Repository
	)
	gormDB.Model(&lib.Repository{}).Count(&repositoryCount)
	if removeRepositoryCount := int(repositoryCount) - int(MAX_REPOSITORY_RECORES); removeRepositoryCount > 0 {
		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))
		gormDB.Model(&lib.Repository{}).Order("updated_at ASC").Limit(removeRepositoryCount).Find(&removeRepositories)
		gormDB.Unscoped().Delete(&removeRepositories)
	}
	logrus.Info("Successfully finished to update repositories.")
	return nil
}

func UpdateIssues() error {
	logrus.Info("Start updating issues.")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when updating repositories")
	}
	defer sqlDB.Close()

	// Get target repositories to update issue
	var (
		notInitializedRepositories []lib.Repository
		initializedRepositories    []lib.Repository
	)
	gormDB.Where("issue_initialized = ?", false).Limit(int(UPDATE_ISSUE_BATCH_SIZE)).Find(&notInitializedRepositories)
	if restRepositoryNum := int(UPDATE_ISSUE_BATCH_SIZE) - len(notInitializedRepositories); restRepositoryNum > 0 {
		gormDB.Where("issue_initialized = ?", true).Order("updated_at ASC").Limit(restRepositoryNum).Find(&initializedRepositories)
	}
	logrus.Info(fmt.Sprintf("not initialized repository: %d", len(notInitializedRepositories)))
	logrus.Info(fmt.Sprintf("initialized repository: %d", len(initializedRepositories)))

	repositories := append(notInitializedRepositories, initializedRepositories...)

	// Update issues
	var (
		wg    sync.WaitGroup
		mutex = &sync.Mutex{}
	)
	for i := 0; i < len(repositories)/int(UPDATE_ISSUE_MINI_BATCH_SIZE); i++ {
		// Create mini batch of repository
		lower := i * int(UPDATE_ISSUE_MINI_BATCH_SIZE)
		upper := lib.Min((i+1)*int(UPDATE_ISSUE_MINI_BATCH_SIZE), len(repositories)-1)
		miniBatchRepositories := repositories[lower:upper]

		// Fetch issues concurrently for each mini batchn
		modelMiniBatchRepositories := make([]lib.Repository, 0, UPDATE_ISSUE_MINI_BATCH_SIZE)
		for _, repository := range miniBatchRepositories {
			wg.Add(1)
			go func(repository lib.Repository) {
				defer wg.Done()
				issues := FetchIssues(repository.Name)
				repository.Issues = issues
				repository.IssueInitialized = true

				mutex.Lock()
				modelMiniBatchRepositories = append(modelMiniBatchRepositories, repository)
				mutex.Unlock()

			}(repository)
		}
		wg.Wait()
		gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Save(&modelMiniBatchRepositories)
	}
	logrus.Info("Successfully finished to update issues.")
	return nil
}

func UpdateFrontLanguages() error {
	logrus.Info("Start updating languages.")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when updating front languages")
	}
	defer sqlDB.Close()

	// Fetch languages from other table
	var (
		languages    []lib.FrontLanguage
		oldLanguages []lib.FrontLanguage
	)
	query := gormDB.Model(&lib.Repository{})
	query.Select("language AS name, COUNT(language) AS repository_count")
	query.Where("language != ?", "")
	query.Group("language")
	query.Order("repository_count DESC")
	query.Find(&languages)

	// Update languages inside transaction
	gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLanguages).Error; err != nil {
			logrus.Error("error occured when deleting old front languages")
			return err
		}
		if err := tx.Create(&languages).Error; err != nil {
			logrus.Error("error occured when inserting new front languages")
			return err
		}
		return nil
	})

	return nil
}

func UpdateLicenses() error {
	logrus.Info("Start updating licenses.")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when updating front languages")
	}
	defer sqlDB.Close()

	// Fetch licenses from other table
	var (
		licenses    []lib.FrontLicense
		oldLicenses []lib.FrontLicense
	)
	query := gormDB.Model(&lib.Repository{})
	query.Select("license AS name, COUNT(license) AS repository_count")
	query.Where("license != ?", "")
	query.Group("license")
	query.Order("repository_count DESC")
	query.Find(&licenses)

	// Update licenses inside transaction
	gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLicenses).Error; err != nil {
			logrus.Error("error occured when deleting old front licenses")
			return err
		}
		if err := tx.Create(&licenses).Error; err != nil {
			logrus.Error("error occured when inserting new front licenses")
			return err
		}
		return nil
	})
	return nil
}

func UpdateLabels() error {
	logrus.Info("Start updating labels.")

	// Connect DB
	gormDB, sqlDB, err := getDBClient()
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("error occured when updating front labels")
	}
	defer sqlDB.Close()

	// Fetch labels from other table
	var (
		labels    []lib.FrontLabel
		oldLabels []lib.FrontLabel
	)
	query := gormDB.Model(&lib.Label{})
	query.Select("name, COUNT(name) AS issue_count")
	query.Where("name != ?", "")
	query.Group("name")
	query.Order("issue_count DESC")
	query.Find(&labels)

	// Update labels inside transaction
	gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&oldLabels).Error; err != nil {
			logrus.Error("error occured when deleting old front labels")
			return err
		}
		if err := tx.Create(&labels).Error; err != nil {
			logrus.Error("error occured when inserting new front labels")
			return err
		}
		return nil
	})
	return nil
}
