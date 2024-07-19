package db

import (
	"database/sql"

	"github.com/ishanz23/go-turso-starter-api/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error
	// url := fmt.Sprintf("%s?authToken=%s", config.Config("TURSO_URL"), config.Config("TURSO_TOKEN"))
	url := config.Config("NEON_URL")

	DB, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}

	if _, err = CreateLocationTable(); err != nil {
		return err
	}
	if _, err = CreateHomestayTable(); err != nil {
		return err
	}
	if _, err = CreateRoomTable(); err != nil {
		return err
	}
	return nil
}
