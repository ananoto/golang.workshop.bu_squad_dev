package dao

import (
	"workshop/model"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func InsertSquad(squad model.Squad)  {
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("squad")
	defer session.Close()
	//Insert
	err := collection.Insert(squad)
	//Check null
	if err != nil {
		log.Fatal(err)
		return
	}
}

func FindSquadByName(name string) model.Squad {
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("squad")
	defer session.Close()
	var result model.Squad
	err := collection.Find(bson.M{"name":name}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}
