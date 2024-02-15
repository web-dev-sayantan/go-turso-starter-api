package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/data"
	"github.com/ishanz23/go-turso-starter-api/db"
)

func HandleGetAllHomestays(c *fiber.Ctx) {
	rows, err := db.DB.Query("SELECT homestay.id, homestay.name, homestay.address, location.name from homestay INNER JOIN location ON homestay.locationId=location.id")

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()
	homestays := []data.Homestay{}
	for rows.Next() {
		homestay := data.Homestay{}
		if err := rows.Scan(&homestay.ID, &homestay.Name, &homestay.Address, &homestay.LocationName); err != nil {
			c.Status(500).JSON(&fiber.Map{"success": false, "message": err.Error()})
			return
		}
		homestays = append(homestays, homestay)
	}
	if err := c.JSON(&fiber.Map{"success": true, "homestays": homestays}); err != nil {
		c.Status(500).JSON(&fiber.Map{"success": false, "message": err.Error()})
		return
	}
}

func HandleGetHomestayById(c *fiber.Ctx) {
	id := c.Params("id")
	homestay := data.Homestay{}
	row := db.DB.QueryRow(`SELECT homestay.id, homestay.name, homestay.address, location.name from homestay INNER JOIN location ON homestay.locationId=location.id WHERE homestay.id = ?`, id)

	switch err := row.Scan(&homestay.ID, &homestay.Name, &homestay.Address, &homestay.LocationName); err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "No rows returned",
		})
	case nil:
		if err := c.JSON(&fiber.Map{
			"success":  true,
			"homestay": homestay,
		}); err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": "JSONify failed",
				"error":   err.Error(),
			})
		}
	default:
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
}
