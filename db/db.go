package db

import (
	"database/sql"
	"fmt"

	"github.com/ishanz23/go-turso-starter-api/config"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func Connect() error {
	var err error
	url := fmt.Sprintf("%s?authToken=%s", config.Config("TURSO_URL"), config.Config("TURSO_TOKEN"))

	DB, err = sql.Open("libsql", url)
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
