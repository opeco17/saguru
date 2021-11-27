package main

import (
	"database/sql"
	"fmt"
	"opeco17/oss-book/lib"
	"os"

	"gorm.io/gorm"
)

func getDBClient() (*gorm.DB, *sql.DB, error) {
	gormDB, sqlDB, err := lib.GetDBClient(
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	)
	return gormDB, sqlDB, err
}

func orderMetrics() []string {
	return []string{
		"star_count",
		"fork_count",
	}
}

func setOrderQuery(query *gorm.DB, orderby string) {
	order_query := "repositories.id"
	for _, metric := range orderMetrics() {
		if orderby == metric+"_asc" {
			order_query = fmt.Sprintf("repositories.%s ASC", metric)
			break
		} else if orderby == metric+"_desc" {
			order_query = fmt.Sprintf("repositories.%s DESC", metric)
			break
		}
	}
	query.Order(order_query)
}

func setDistinctQuery(query *gorm.DB, orderby string) {
	distinct_query := "repositories.id"
	for _, metric := range orderMetrics() {
		if orderby == metric+"_asc" {
			distinct_query = fmt.Sprintf("repositories.id, repositories.%s", metric)
			break
		} else if orderby == metric+"_desc" {
			distinct_query = fmt.Sprintf("repositories.id, repositories.%s", metric)
			break
		}
	}
	query.Distinct(distinct_query)
}
