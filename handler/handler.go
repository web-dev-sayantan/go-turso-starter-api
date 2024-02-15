package handler

import (
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/db"
)

func HandleGetAllHomestays(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT * from homestay")

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	defer rows.Close()
}
