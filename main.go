package main

import (
	"database/sql"
	"hash/crc32"
	"io/ioutil"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func hash(link string) string {
	num = crc32.ChecksumIEEE([]byte(link))
	return fmt.Sprintf("%x", num)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runScript(db *sql.DB, file string) {
	content, err := ioutil.ReadFile(file)
	checkError(err)

	contentStr := string(content)
	_, err = db.Exec(contentStr)
	checkError(err)
}

func main() {
	err := godotenv.Load()
	checkError(err)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request")
		jsonResponse(w, http.StatusOK, "yay it works")
	})

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	checkError(err)
	defer db.Close()

	// runScript(db, "sql/init.sql")

	fmt.Println("Running server on port " + port)

	http.ListenAndServe(":"+port, r)
	fmt.Println("Stopping server")
}
