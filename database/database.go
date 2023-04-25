package database

import (
	. "cdn/api/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Connect to database
	var err error
	db, err = sql.Open("postgres", Config.DbUrl + "?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	setupSchema()
}

func runQuery(query string, args ...any) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatalf("ERROR EXECUTING QUERY:\n\n%s\n\nERROR: %s", query, err.Error())
	}

	return rows
}

func setupSchema() {
	runQuery(`
CREATE TABLE IF NOT EXISTS uploads (
	file_name VARCHAR(40) NOT NULL,
	checksum bytea NOT NULL
)`)
}

func AddImage(filename string, checksum []byte) (string, bool) {
	// Check image exists
	var existingFileName string
	db.QueryRow("SELECT file_name FROM uploads WHERE checksum = $1", string(checksum[:])).Scan(&existingFileName)
	if len(existingFileName) != 0 {
		return existingFileName, true
	}

	// Add image
	runQuery("INSERT INTO uploads (file_name, checksum) VALUES ($1, $2)", filename, string(checksum[:]))
	return filename, false
}