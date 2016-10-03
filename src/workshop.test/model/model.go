package model

import "gopkg.in/mgo.v2/bson"

//Entities
type Bu struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Squads      []Squad       `json:"squads" bson:"squads"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}


type Squad struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Devs        []Dev         `json:"devs" bson:"devs"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}

type Dev struct {
	Id         bson.ObjectId  `json:"_id" bson:"_id,omitempty"`
	DevId       string        `json:"devId" bson:"devId"`
	Dev_name    string        `json:"dev_name" bson:"dev_name"`
	Create_date int64         `json:"create_date" bson:"create_date"`
	Update_date int64         `json:"update_date" bson:"update_date"`
	Active      bool          `json:"active" bson:"active"`
}

//Dump
type SquadBu struct {
	Bu_name    string `json:"bu_name" bson:"bu_name"`
	Squad_name string `json:"squad_name" bson:"squad_name"`
}

type DevSquad struct {
	Dev_name   string `json:"dev_name" bson:"dev_name"`
	Squad_name string `json:"squad_name" bson:"squad_name"`
}

type SquadDeactive struct {
	Squad_name string `json:"squad_name" bson:"squad_name"`
}

type BuDeactive struct {
	Bu_name string `json:"bu_name" bson:"bu_name"`
}


