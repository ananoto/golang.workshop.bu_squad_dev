package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
	"workshop/controller/dao"
	"workshop/model"
)

func FindSquadByName(name string)  model.Squad {
	result := dao.FindSquadByName(name)
	return result

}

func InsertSquad(squad model.Squad, c *gin.Context) {
	//Set
	squad.Active = false
	squad.Create_date = time.Now().Unix()
	squad.Update_date = time.Now().Unix()
	dao.InsertSquad(squad)
	fmt.Printf(">>>>", squad)
	c.JSON(200, gin.H{
		"name":        squad.Name,
		"active":      squad.Active,
		"create_date": time.Now(),
	})
}

func InsertDevToSquad(devSquad model.DevSquad) (model.Squad, error) {

	squad_result := dao.FindSquadByName(devSquad.Squad_name)
	squad_name := bson.M{"name": squad_result.Name}

	dev_result := dao.FindDevByName(devSquad.Dev_name)
	if dev_result.Active == true {
		return model.Squad{}, nil
	} else {
		//Connect Dev DB
		dev_session := dao.IntialDB()
		defer dev_session.Close()
		dev_collection := dev_session.DB("workshop").C("dev")
		dev_name := bson.M{"dev_name": dev_result.Dev_name}
		//Set Dev
		dev_result.Active = true
		dev_result.Update_date = time.Now().Unix()
		//Update bu
		devErr := dev_collection.Update(dev_name, dev_result)
		if devErr != nil {
			log.Fatal(devErr)
		}
		//Set Squad
		squad_result.Devs = append(squad_result.Devs, dev_result)
		squad_result.Active = true
		squad_result.Update_date = time.Now().Unix()
		//Connect Squad DB
		session := dao.IntialDB()
		defer session.Close()
		collection := session.DB("workshop").C("squad")
		defer session.Close()
		//Update bu
		err := collection.Update(squad_name, squad_result)
		//Check null
		if err != nil {
			log.Fatal(err)
			return model.Squad{}, err
		} else {
			result := dao.FindSquadByName(devSquad.Squad_name)
			return result, nil
		}
	}

}

func SquadDeacivate(squadDeactive model.SquadDeactive) (model.Squad, error) {

	squad_result := dao.FindSquadByName(squadDeactive.Squad_name)
	squad_name := bson.M{"name": squad_result.Name}

	if squad_result.Active == false {
		return model.Squad{}, nil
	} else {
		//Dev DB Connecting
		dev_session := dao.IntialDB()
		defer dev_session.Close()
		dev_collection := dev_session.DB("workshop").C("dev")

		for i := 0; i < len(squad_result.Devs); i++ {
			fmt.Println("<<<Deactive>>>", squad_result.Devs[i])
			dev_result := dao.FindDevByName(squad_result.Devs[i].Dev_name)
			dev_name := bson.M{"dev_name": dev_result.Dev_name}
			dev_result.Active = false
			dev_result.Update_date = time.Now().Unix()
			//Update bu
			err := dev_collection.Update(dev_name, dev_result)
			if err != nil {
				log.Fatal(err)
			}
		}

		squad_result.Devs = nil
		squad_result.Active = false
		squad_result.Update_date = time.Now().Unix()

		session := dao.IntialDB()
		defer session.Close()
		collection := session.DB("workshop").C("squad")
		//Update bu
		err := collection.Update(squad_name, squad_result)
		//Check null
		if err != nil {
			log.Fatal(err)
			return model.Squad{}, err
		} else {
			result := dao.FindSquadByName(squadDeactive.Squad_name)
			return result, nil
		}
	}
}

func SquadDeacivateByName(squad model.Squad){

	squad_result := dao.FindSquadByName(squad.Name)
	squad_name := bson.M{"name": squad_result.Name}

	if squad_result.Active == false {

	} else {
		dev_session := dao.IntialDB()
		defer dev_session.Close()
		dev_collection := dev_session.DB("workshop").C("dev")

		for i := 0; i < len(squad_result.Devs); i++ {
			fmt.Println("<<<Deactive>>>", squad_result.Devs[i])
			fmt.Println("<<<squad Dev Name>>>",squad_result.Devs[i].Dev_name)
			dev_result := dao.FindDevByName(squad_result.Devs[i].Dev_name)
			fmt.Println("<<<dev result>>>",dev_result)
			dev_name := bson.M{"dev_name": dev_result.Dev_name}
			dev_result.Active = false
			dev_result.Update_date = time.Now().Unix()
			//Update bu
			err := dev_collection.Update(dev_name, dev_result)
			if err != nil {
				log.Fatal(err)
			}
		}

		squad_result.Devs = nil
		squad_result.Active = false
		squad_result.Update_date = time.Now().Unix()

		session := dao.IntialDB()
		defer session.Close()
		collection := session.DB("workshop").C("squad")
		//Update bu
		err := collection.Update(squad_name, squad_result)
		//Check null
		if err != nil {
			log.Fatal(err)/*
			return model.Squad{}, err*/
		} else {
			/*result := dao.FindSquadByName(squad.Name)
			return result, nil*/
		}
	}
}
