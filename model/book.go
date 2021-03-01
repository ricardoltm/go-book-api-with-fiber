package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int 	`json:"rating"`
	Year int `json:"year"`
	Publisher string `json:"publisher"`
}

type Books struct {
	Books []Book `json:"books"`
}