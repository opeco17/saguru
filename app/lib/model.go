package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repository struct {
	gorm.Model
	Name   string
	Issues []Issue
}

type Issue struct {
	gorm.Model
	URL          string
	RepositoryID uint
}

func GetDBClient() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connect database.")
	}
	return db
}

func Init(db *gorm.DB) {
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&Repository{}, &Issue{})
}

func Create(db *gorm.DB) {
	db.Create(&Repository{
		Name: "test/repo",
		Issues: []Issue{
			{URL: "test"},
			{URL: "test"},
		},
	})
}
