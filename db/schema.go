package db

import (
	"database/sql"
	"fmt"
)

func CreateLocationTable() (*sql.Rows, error) {
	fmt.Println("Creating location table")
	return DB.Query(`CREATE TABLE IF NOT EXISTS location (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR UNIQUE NOT NULL,
		lat FLOAT,
		long FLOAT,
		state VARCHAR NOT NULL,
		altitude INTEGER,
		description VARCHAR,
		coverUrl VARCHAR
	)`)
}

func DeleteLocationTable() (*sql.Rows, error) {
	return DB.Query(`DROP TABLE IF EXISTS location`)
}

func DeleteHomestayTable() (*sql.Rows, error) {
	return DB.Query(`DROP TABLE IF EXISTS homestay`)
}

func DeleteRoomTable() (*sql.Rows, error) {
	return DB.Query(`DROP TABLE IF EXISTS room`)
}
func CreateHomestayTable() (*sql.Rows, error) {
	return DB.Query(`CREATE TABLE IF NOT EXISTS homestay (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		address VARCHAR,
		locationId INTEGER,
		FOREIGN KEY (locationId) REFERENCES location(id) 
	)`)
}
func CreateRoomTable() (*sql.Rows, error) {
	return DB.Query(`CREATE TABLE IF NOT EXISTS room (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		category VARCHAR NOT NULL,
		baseOccupancy INTEGER NOT NULL,
		extraOccupancy INTEGER DEFAULT 0,
		toiletAttached BOOLEAN DEFAULT true,
		balconyAttached BOOLEAN DEFAULT false,
		kitchenAttached BOOLEAN DEFAULT false,
		airConditioned BOOLEAN DEFAULT false,
		recommended BOOLEAN DEFAULT false,
		isDorm BOOLEAN DEFAULT false,
		homestayId INTEGER,
		FOREIGN KEY (homestayId) REFERENCES homestay(id) 
	)`)
}
