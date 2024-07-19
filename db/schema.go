package db

import (
	"database/sql"
)

func CreateLocationTable() (*sql.Rows, error) {
	return DB.Query(`CREATE TABLE IF NOT EXISTS location (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		lat FLOAT,
		long FLOAT,
		state VARCHAR(40) NOT NULL,
		altitude INTEGER,
		description TEXT,
		coverUrl TEXT
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
		id SERIAL PRIMARY KEY,
		name VARCHAR(120) NOT NULL,
		address VARCHAR(255),
		locationId INT NOT NULL,
		CONSTRAINT fk_location
			FOREIGN KEY(locationId) REFERENCES location(id) 
	)`)
}
func CreateRoomTable() (*sql.Rows, error) {
	return DB.Query(`CREATE TABLE IF NOT EXISTS room (
		id SERIAL PRIMARY KEY,
		name VARCHAR(120) NOT NULL,
		category VARCHAR(25) NOT NULL,
		baseOccupancy INTEGER NOT NULL,
		extraOccupancy INTEGER DEFAULT 0,
		toiletAttached BOOLEAN DEFAULT true,
		balconyAttached BOOLEAN DEFAULT false,
		kitchenAttached BOOLEAN DEFAULT false,
		airConditioned BOOLEAN DEFAULT false,
		recommended BOOLEAN DEFAULT false,
		isDorm BOOLEAN DEFAULT false,
		homestayId INT NOT NULL,
		CONSTRAINT fk_homestay FOREIGN KEY (homestayId) REFERENCES homestay(id) 
	)`)
}
