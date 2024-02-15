package handler

import (
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/data"
	"github.com/ishanz23/go-turso-starter-api/db"
)

func HandleGetAllHomestays(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT name, address, locationId from homestay")

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	defer rows.Close()
	homestays := []data.Homestay{}
	for rows.Next() {
		homestay := data.Homestay{}
		if err := rows.Scan(&homestay.Name, &homestay.Address, &homestay.LocationId); err != nil {
			c.Status(500).JSON(&fiber.Map{"success": false, "error": err})
			return
		}
		homestays = append(homestays, homestay)
	}
	if err := c.JSON(&fiber.Map{"success": true, "homestays": homestays}); err != nil {
		c.Status(500).JSON(&fiber.Map{"success": false, "error": err})
		return
	}
}
