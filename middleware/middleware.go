package middleware

import (
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/ishanz23/go-turso-starter-api/config"
)

func AuthReq() func(*fiber.Ctx) {
	config := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(config)
	return err
}
