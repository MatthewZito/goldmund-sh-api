package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/MatthewZito/goldmund-sh-api/models"
	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	FResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}

func FindFlightEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entry := models.Entry{
		Title:    "Test title",
		Subtitle: "test subtitle",
	}

	FResponse(w, http.StatusOK, entry)
}

func FResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func FError(w http.ResponseWriter, code int, msg string) {
	FResponse(w, code, map[string]string{"error": msg})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Health)
	r.HandleFunc("/entries", GetAllEntries)
	r.HandleFunc("/entries/{id}", GetAllEntries)

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
