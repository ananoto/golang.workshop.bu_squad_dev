package model

import "gopkg.in/mgo.v2/bson"

type Bu struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Squads      []Squad       `json:"squads" bson:"squads"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}
