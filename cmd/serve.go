package main

import (
	"log"
	"net/http"

	"github.com/MatthewZito/goldmund-sh-api/controllers"
	"github.com/MatthewZito/goldmund-sh-api/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

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
	r.Path("/").HandlerFunc(controllers.Health)
	r.Path("/entries").Queries("slug", "{*?}").HandlerFunc(ec.GetEntryBySlug)
	r.Path("/entries").HandlerFunc(ec.GetAllEntries)
	r.Path("/entries").Queries("last", "{*?}").HandlerFunc(ec.GetAllEntries)

	if err := http.ListenAndServe(":5000", c.Handler(r)); err != nil {
		log.Fatal(err)
	}
}
