package database

import (
	. "cdn/api/util"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	// Connect to database
	var err error
	db, err = sqlx.Open("postgres", Config.DbUrl+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	setupSchema()
}

func setupSchema() {
	// Create 'uploads' table if not exists
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS uploads (
	file_name VARCHAR(40) NOT NULL,
	checksum bytea NOT NULL
)`)
	if err != nil {
		log.Fatal(err)
	}
}

func AddImage(filename string, checksum []byte) (string, bool) {
	// Check if image already exists in the database
	var existingFileName string
	err := db.Get(&existingFileName, "SELECT file_name FROM uploads WHERE checksum = $1", string(checksum[:]))
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if len(existingFileName) != 0 {
		return existingFileName, true
	}

	// Add image to the database
	_, err = db.Exec("INSERT INTO uploads (file_name, checksum) VALUES ($1, $2)", filename, string(checksum[:]))
	if err != nil {
		log.Fatal(err)
	}
	return filename, false
}