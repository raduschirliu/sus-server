package db

import (
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	// Import the PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/raduschirliu/sus-server/util"
)

var db *sqlx.DB

// Init database connection
func Init() {
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	db, err = sqlx.Open("postgres", dbURL)
	util.CheckError(err)

	RunScript("db/sql/init.sql")
}

// Close database connection and cleanup
func Close() {
	db.Close()
}

// Query database for a shortened link with given id
func Query(id string) util.Result {
	var res util.Result
	err := db.Get(&res.Link, "SELECT * FROM links WHERE id=$1", id)

	if err != nil {
		res.Error = util.StringPtr("No results found")
	}

	return res
}

// Insert a new link into the database
func Insert(id string, link string) util.Result {
	var res util.Result
	res.Link = &util.Link{ID: id, Link: link}

	_, err := db.Exec("INSERT INTO links (id, link) VALUES ($1, $2)", id, link)

	if err != nil {
		res.Error = util.StringPtr("Key already exists")
	}

	return res
}

// RunScript will run a local PSQL script file
func RunScript(file string) {
	content, err := ioutil.ReadFile(file)
	util.CheckError(err)

	contentStr := string(content)
	_, err = db.Exec(contentStr)
	util.CheckError(err)
}
