package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/raduschirliu/sus-server/db"
	"github.com/raduschirliu/sus-server/util"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func getLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request on GET")
	jsonResponse(w, http.StatusOK, "GET works")
}

func postLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request on POST")
	jsonResponse(w, http.StatusOK, "POST works")
}

func main() {
	err := godotenv.Load()
	util.CheckError(err)

	r := mux.NewRouter()
	r.HandleFunc("/", getLink).Methods("GET")
	r.HandleFunc("/", postLink).Methods("POST")

	port := os.Getenv("PORT")
	db.Init()
	defer db.Close()

	fmt.Println("Running server on port " + port)

	http.ListenAndServe(":"+port, r)
	fmt.Println("Stopping server")
}
