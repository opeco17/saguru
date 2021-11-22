package main

import (
	"fmt"
	"opeco17/oss-book/lib"
	"sync"

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
	uniqueQuery := [...]string{"stars:30..500", "stars:500..1500", "stars:>1500"}
	for _, eachQuery := range uniqueQuery {
		repositories := FetchRepositories("good-first-issues:>1", eachQuery)
		gormDB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&repositories)
	}

	// Adjust number of repositories
	var (
		repositoryCount    int64
		removeRepositories []lib.Repository
	)
	gormDB.Model(&lib.Repository{}).Count(&repositoryCount)
	removeRepositoryCount := int(repositoryCount) - 3*int(REPOSITORIES_API_MAX_RESULTS)
	if removeRepositoryCount > 0 {
		logrus.Info(fmt.Sprintf("%d repositories will be removed.", removeRepositoryCount))
		gormDB.Model(&lib.Repository{}).Limit(removeRepositoryCount).Find(&removeRepositories)
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

	var (
		repositories []lib.Repository
		wg           sync.WaitGroup
		mutex        = &sync.Mutex{}
	)
	gormDB.Find(&repositories)

	for i := 0; i < len(repositories)/int(MINI_BATCH_SIZE); i++ {
		// Create mini batch of repository
		lower := i * int(MINI_BATCH_SIZE)
		upper := lib.Min((i+1)*int(MINI_BATCH_SIZE), len(repositories)-1)
		miniBatchRepositories := repositories[lower:upper]

		// Fetch issues concurrently for each mini batchn
		modelMiniBatchRepositories := make([]lib.Repository, 0, MINI_BATCH_SIZE)
		for _, repository := range miniBatchRepositories {
			wg.Add(1)
			go func(repository lib.Repository) {
				defer wg.Done()
				issues := FetchIssues(repository.Name)
				repository.Issues = issues

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
