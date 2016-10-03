package model

import "gopkg.in/mgo.v2/bson"

type Squad struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Devs        []Dev         `json:"devs" bson:"devs"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}
