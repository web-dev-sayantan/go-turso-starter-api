package common

import "database/sql"

type CustomContext struct {
	db *sql.DB
}
