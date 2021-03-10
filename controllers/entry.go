package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MatthewZito/goldmund-sh-api/models"
	"github.com/MatthewZito/goldmund-sh-api/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EntryController maintains a pointer to a given db collection
type EntryController struct {
	coll *mongo.Collection
}

// InitEntryController instantiates the entry controllers
func InitEntryController(s *mongo.Collection) *EntryController {
	return &EntryController{s}
}

// GetEntryBySlug fetches from the cluster endpoint a single entry, specified by its unique slug identifier
func (ec EntryController) GetEntryBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.FormValue("slug")

	var entry models.Entry

	err := ec.coll.FindOne(context.Background(), bson.M{"slug": slug}).Decode(&entry)

	if err != nil {
		shared.FError(w, http.StatusBadRequest, "Invalid slug")
		return
	}

	shared.FResponse(w, http.StatusOK, entry)
}

// GetAllEntries fetches all entries from the cluster endpoint
func (ec EntryController) GetAllEntries(w http.ResponseWriter, r *http.Request) {
	lastProcessedID := r.FormValue("last")

	var entries []models.Entry

	options := options.Find()

	options.SetSort(bson.M{"createdAt": -1})

	options.SetLimit(10)

	curs, err := ec.BuildEntryFilter(lastProcessedID, options)

	if err != nil {
		shared.FError(w, http.StatusBadRequest, "Failed to fetch entries")
		return
	}

	// await parsing of all available entries
	defer curs.Close(context.Background())

	for curs.Next(context.Background()) {

		var entry models.Entry
		err := curs.Decode(&entry)

		if err != nil {
			shared.FError(w, http.StatusBadRequest, "Failed to deserialize entries")
			return
		}

		entries = append(entries, entry)
	}

	if err := curs.Err(); err != nil {
		shared.FError(w, http.StatusBadRequest, "Failed to parse entries")
		return
	}

	shared.FResponse(w, http.StatusOK, entries)
}

// BuildEntryFilter constructs a filtered cursor contingent on `lastProcessedID`
func (ec EntryController) BuildEntryFilter(lastProcessedID string, options *options.FindOptions) (curs *mongo.Cursor, err error) {
	if lastProcessedID != "" {
		fmt.Println(lastProcessedID)
		return ec.coll.Find(context.Background(), bson.M{"createdAt": bson.M{"$lt": lastProcessedID}}, options)
	}

	return ec.coll.Find(context.Background(), bson.M{}, options)
}
