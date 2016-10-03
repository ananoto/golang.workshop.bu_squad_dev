package dao

import (
	"workshop/model"
	"gopkg.in/mgo.v2/bson"
)

func InsertDev(dev model.Dev)  error{
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("dev")
	defer session.Close()
	//Insert
	err := collection.Insert(dev)
	//Check null
	if err != nil {
		return err
	}
	return nil
}

func FindDevByName(name string) model.Dev {
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("dev")
	defer session.Close()
	var result model.Dev
	err := collection.Find(bson.M{"dev_name":name}).One(&result)
	if err != nil {	panic(err)}
	return result
}
