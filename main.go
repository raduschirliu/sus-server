package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/raduschirliu/sus-server/db"
	"github.com/raduschirliu/sus-server/util"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	var response []byte

	if data != nil {
		response, _ = json.Marshal(data)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func decodeJSON(dest interface{}, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	util.CheckError(err)
	json.Unmarshal(bytes, dest)
}

func getLink(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET: %s\n", r.URL)

	type query struct {
		ID *string
	}

	var q query
	decodeJSON(&q, r)

	if q.ID == nil {
		jsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	fmt.Println("Getting: ", q.ID)
	res := db.Query(*q.ID)

	jsonResponse(w, http.StatusOK, res)
}

func postLink(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST: %s\n", r.URL)

	type query struct {
		Link *string
	}

	var q query
	decodeJSON(&q, r)

	if q.Link == nil {
		jsonResponse(w, http.StatusBadRequest, nil)
		return
	}

	id := util.Hash(*q.Link)
	fmt.Println("Inserting: ", id, " ", q.Link)
	res := db.Insert(id, *q.Link)

	jsonResponse(w, http.StatusOK, res)
}

func main() {
	godotenv.Load()

	r := mux.NewRouter()
	r.HandleFunc("/link", getLink).Methods("GET")
	r.HandleFunc("/link", postLink).Methods("POST")

	port := os.Getenv("PORT")
	db.Init()
	defer db.Close()

	fmt.Println("Running server on port " + port)

	err := http.ListenAndServe(":"+port, r)
	util.CheckError(err)
}
