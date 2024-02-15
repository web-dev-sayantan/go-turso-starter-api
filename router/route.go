package router

import (
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/handler"
	"github.com/ishanz23/go-turso-starter-api/middleware"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", middleware.AuthReq())

	api.Get("/homestays", handler.HandleGetAllHomestays)
	api.Get("/", func(c *fiber.Ctx) { c.JSON("Hi") })
}
