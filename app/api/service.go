package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

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
	db.AutoMigrate(&Repository{}, &User{}, &Issue{}, &Label{})
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
