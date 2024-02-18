package router

import (
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/handler"
	"github.com/ishanz23/go-turso-starter-api/middleware"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", middleware.AuthReq())

	api.Get("/", func(c *fiber.Ctx) { c.JSON("Hi") })
	api.Get("/homestays", handler.HandleGetAllHomestays)
	api.Get("/homestays/:id", handler.HandleGetHomestayById)
	api.Get("/locations", handler.HandleGetAllLocations)
	api.Get("/homestays/:id/rooms", handler.HandleGetRoomsByHomestayId)
	api.Get("/homestays/:homestayId/rooms/:roomId", handler.HandleGetSingleRoomFromHomestayById)
}
