package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/ishanz23/go-turso-starter-api/db"
	"github.com/ishanz23/go-turso-starter-api/router"
)

func main() {
	start := time.Now()
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(middleware.Logger())

	router.SetupRoutes(app)
	fmt.Println("Time taken: ", time.Since(start))
	app.Listen(8000)
}
