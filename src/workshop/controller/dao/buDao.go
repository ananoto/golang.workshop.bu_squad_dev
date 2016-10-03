package dao

import (
	"log"
	"workshop/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
)

func InsertBu(bu model.Bu)  {
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("bu")
	defer session.Close()
	//Insert
	err := collection.Insert(bu)
	//Check null
	if err != nil {
		log.Fatal(err)
		return
	}
}

func FindBuByName(name string) model.Bu {
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("bu")
	defer session.Close()
	var result model.Bu
	err := collection.Find(bson.M{"name":name}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func FindAllBuDao(c *gin.Context) []model.Bu{
	//Database Connecting
	session := IntialDB()
	collection := session.DB("workshop").C("bu")
	defer session.Close()
	var results []model.Bu
	err := collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func UpdateBu() {
	//Database Connecting

}

