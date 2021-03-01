package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoltm/go-book-api-with-fiber/infrastructure/database"
	"github.com/ricardoltm/go-book-api-with-fiber/router"
)

func main() {
	app := fiber.New()

	db := database.InitDatabase()
	database.DBConn = db
	defer database.DBConn.Close()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Book API with Go Fiber and Gorm")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
