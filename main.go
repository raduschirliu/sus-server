package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request")
		jsonResponse(w, http.StatusOK, "yay it works")
	})

	port := "8080"

	portEnv, portExists := os.LookupEnv("PORT")

	if portExists {
		port = portEnv
	}

	fmt.Println("Running server on port " + port)

	http.ListenAndServe(":"+port, r)
	fmt.Println("Stopping server")
}
