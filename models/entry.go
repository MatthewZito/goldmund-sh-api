package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	Title     string             `json:"title" bson:"title" `
	Subtitle  string             `json:"subtitle" bson:"subtitle"`
	Imgsrc    string             `json:"imgSrc" bson:"imgsrc"`
	Content   string             `json:"content" bson:"content"`
	ID        primitive.ObjectID `json:"objectId" bson:"_id"`
	Tags      []string           `json:"tags" bson:"tags"`
	Slug      string             `json:"slug" bson:"slug"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
