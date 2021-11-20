package main

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBUSER     string = "root"
	DBPASSWORD string = "root"
	DBHOST     string = "localhost"
	DBPORT     string = "3306"
	DBNAME     string = "testdb"
)

func GetDBClient() (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUSER, DBPASSWORD, DBHOST, DBPORT, DBNAME,
	)
	gormLogger := logger.New(
		stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags),
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
		log.Fatal("Error occurred when connecting database")
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Successfully connect database.")
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return gormDB, sqlDB
}

func Init() {
	gormDB, sqlDB := GetDBClient()
	defer sqlDB.Close()
	err := gormDB.AutoMigrate(&Repository{}, &User{}, &Issue{}, &Label{})
	if err != nil {
		log.Fatal("Error occurred when migrating database")
		log.Fatal(err)
		os.Exit(1)
	}
}
