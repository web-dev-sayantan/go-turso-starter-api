package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/ishanz23/go-turso-starter-api/db"
	"github.com/ishanz23/go-turso-starter-api/router"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(8000)
}
