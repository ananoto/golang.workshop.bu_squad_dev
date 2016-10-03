package model

import "gopkg.in/mgo.v2/bson"

type Dev struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	DevId       string        `json:"devId" bson:"devId"`
	Dev_name    string        `json:"dev_name" bson:"dev_name"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}
