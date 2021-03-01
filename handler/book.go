package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardoltm/go-book-api-with-fiber/infrastructure/database"
	"github.com/ricardoltm/go-book-api-with-fiber/model"
)


func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var allBooks []model.Book
	
	result := db.Find(&allBooks)
	
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map {
			"message": "Records not found",
		})
	}

	if err := c.JSON(result); err != nil {
		log.Println(err)
	}

	return nil
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	db := database.DBConn
	
	result := db.First(&book, id)

	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map {
			"message": "Record not found",
		})
	}

	if err := c.JSON(result); err != nil {
		log.Println(err)
	}
	
	return nil
}

func NewBook(c *fiber.Ctx) error {
	var bookIn model.Book
	if err := c.BodyParser(&bookIn); err != nil {
		return c.Status(400).JSON(fiber.Map {
			"message": "Error to process data.",
		})
	}

	db := database.DBConn
	result := db.Create(&bookIn)
	if err := c.JSON(result); err != nil {
		return c.Status(400).JSON(fiber.Map {
			"message": "Error to insert a new book.",
		})
	}

	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	checkId := db.First(&book, id)
	if checkId.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map {
			"success": false,
			"message": "Record not found.",
		})
	}

	var bookUp model.Book
	if err := c.BodyParser(&bookUp); err != nil {
		return c.Status(400).JSON(fiber.Map {
			"success": false,
			"message": "Error to process data.",
		})
	}

	db.Model(&book).Update(&bookUp)
	_ = c.JSON(fiber.Map {
			"message": "Record successfully updated.",
		})

	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	checkId := db.First(&model.Book{}, id)
	if checkId.RowsAffected == 0 {
		_ = c.Status(400).JSON(fiber.Map {
			"success": false,
			"message": "Record not found.",
		})
	}

	db.Delete(&model.Book{}, id)
	_ = c.JSON(fiber.Map {
		"message": "Record successfully deleted.",
	})

	return nil
}