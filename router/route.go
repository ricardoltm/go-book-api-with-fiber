package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoltm/go-book-api-with-fiber/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	
	book := v1.Group("/book")
	book.Get("/", handler.GetBooks)
	book.Get("/:id", handler.GetBook)
	book.Post("/", handler.NewBook)
	book.Put("/:id", handler.UpdateBook)
	book.Delete("/:id", handler.DeleteBook)
}
