package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardoltm/go-book-api-with-fiber/model"
)

// Dabase instance
var	DBConn *gorm.DB

func InitDatabase() *gorm.DB {
	var err error
	DBConn, err := gorm.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	DBConn.AutoMigrate(&model.Book{})
	log.Println("AutoMigrate has been completed.")

	// Create
	// DBConn.Create(&model.Book{Title: "Mastering Go", Author: "Mihalis Tsoukalos", Rating: 5, Year: 2019, Publisher: "Packt Publishing"})
	// DBConn.Create(&model.Book{Title: "Building RESTful Web services with Go", Author: "Naren Yellavula", Rating: 4, Year: 2017, Publisher: "Packt Publishing"})
	// DBConn.Create(&model.Book{Title: "Learn Data Structures and Algorithms with Golang", Author: "Bhagvan Kommadi", Rating: 4, Year: 2019, Publisher: "Packt Publishing Ltd."})
	
	return DBConn
}
