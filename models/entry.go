package models

import "github.com/globalsign/mgo/bson"

type Entry struct {
	Title    string        `json:"title" bson:"title" `
	Subtitle string        `json:"subtitle" bson:"subtitle"`
	Imgsrc   string        `json:"img_src" bson:"imgsrc"`
	Content  string        `json:"content" bson:"content"`
	Id       bson.ObjectId `json:"object_id" bson:"_id"`
	Tags     []string      `json:"tags" bson:"tags"`
	Slug     string        `json: "slug" bson: "slug"`
}
