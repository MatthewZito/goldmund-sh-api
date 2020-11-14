package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MatthewZito/goldmund-sh-api/controllers"
	"github.com/MatthewZito/goldmund-sh-api/db"
	"github.com/MatthewZito/goldmund-sh-api/shared"
	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	shared.FResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}

func main() {
	r := mux.NewRouter()

	s, err := db.InitMongoSession()
	if err != nil {
		log.Fatal(err)
	}

	ec := controllers.InitEntryController(s)
	r.HandleFunc("/", Health)
	r.HandleFunc("/entries", ec.GetAllEntries)
	r.HandleFunc("/entries/{slug}", ec.GetEntryBySlug)

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
