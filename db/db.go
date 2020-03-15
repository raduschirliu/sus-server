package db

import (
	"database/sql"
	"io/ioutil"
	"os"

	// Import the PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/raduschirliu/sus-server/util"
)

var db *sql.DB

// Init database connection
func Init() {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	util.CheckError(err)
	defer db.Close()
}

// Close database connection and cleanup
func Close() {
	db.Close()
}

// RunScript will run a local PSQL script file
func RunScript(file string) {
	content, err := ioutil.ReadFile(file)
	util.CheckError(err)

	contentStr := string(content)
	_, err = db.Exec(contentStr)
	util.CheckError(err)
}
