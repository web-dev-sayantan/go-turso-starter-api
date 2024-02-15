package handler

import (
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/data"
	"github.com/ishanz23/go-turso-starter-api/db"
)

func HandleGetAllLocations(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT id, name, state, description, lat, long, altitude, coverUrl from location")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	defer rows.Close()
	locations := []data.Location{}
	for rows.Next() {
		location := data.Location{}
		if err := rows.Scan(&location.ID, &location.Name, &location.State, &location.Description, &location.Lat, &location.Long, &location.Altitude, &location.CoverUrl); err != nil {
			c.Status(500).JSON(&fiber.Map{"success": false, "error": err})
			return
		}
		locations = append(locations, location)
	}
	if err := c.JSON(&fiber.Map{"success": true, "locations": locations}); err != nil {
		c.Status(500).JSON(&fiber.Map{"success": false, "error": err})
		return
	}
}
