package controllers

import (
	"net/http"

	"github.com/MatthewZito/goldmund-sh-api/models"
	"github.com/MatthewZito/goldmund-sh-api/shared"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const coll = "entries"

type EntryController struct {
	session *mgo.Session
}

func InitEntryController(s *mgo.Session) *EntryController {
	return &EntryController{s}
}

func (ec EntryController) GetEntryBySlug(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	slug := p["slug"]
	var entry models.Entry

	if err := ec.session.DB("").C(coll).Find(bson.M{"slug": slug}).One(&entry); err != nil {
		shared.FError(w, http.StatusBadRequest, "Invalid slug")
		return
	}

	shared.FResponse(w, http.StatusOK, entry)
}

func (ec EntryController) GetAllEntries(w http.ResponseWriter, r *http.Request) {

	coll := ec.session.DB("").C(coll)

	var entries []*models.Entry

	err := coll.Find(nil).All(&entries)
	if err != nil {
		shared.FError(w, http.StatusBadRequest, "Invalid slug")
		return
	}
	shared.FResponse(w, http.StatusOK, entries)
}
