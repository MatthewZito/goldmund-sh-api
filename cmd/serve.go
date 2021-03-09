package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MatthewZito/goldmund-sh-api/controllers"
	"github.com/MatthewZito/goldmund-sh-api/db"
	"github.com/MatthewZito/goldmund-sh-api/shared"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS", "HEAD", "POST", "PUT"},
	})

	s, err := db.InitMongoSession()
	if err != nil {
		log.Fatal(err)
	}

	ec := controllers.InitEntryController(s)
	r.Path("/").HandlerFunc(Health)
	r.Path("/entries").Queries("slug", "{*?}").HandlerFunc(ec.GetEntryBySlug)
	r.Path("/entries").HandlerFunc(ec.GetAllEntries)
	r.Path("/entries").Queries("last", "{*?}").HandlerFunc(ec.GetAllEntries)

	if err := http.ListenAndServe(":5000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
