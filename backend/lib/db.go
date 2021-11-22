package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBClient(user string, password string, host string) (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, DBPORT, DBNAME,
	)
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 3 * time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
		Logger:          gormLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error occurred when connecting database\n%s", err)
	}
	logrus.Info("Successfully connect database.")
	logrus.Info(fmt.Sprintf("User: %s, Host: %s, Database: %s", user, host, DBNAME))
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, nil, err
	}
	return gormDB, sqlDB, nil
}
