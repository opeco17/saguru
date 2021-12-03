package main

import (
	"database/sql"
	"opeco17/gitnavi/lib"
	"os"

	"gorm.io/gorm"
)

func getDBClient() (*gorm.DB, *sql.DB, error) {
	gormDB, sqlDB, err := lib.GetDBClient(
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	)
	return gormDB, sqlDB, err
}
