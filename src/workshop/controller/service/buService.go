package service

import (
	"fmt"
	"log"
	"time"
	"workshop/model"
	"github.com/gin-gonic/gin"
	"workshop/controller/dao"
	"gopkg.in/mgo.v2/bson"
)

func SaveBu(bu model.Bu, c *gin.Context) {
	//Set
	bu.Active = false
	bu.Create_date = time.Now().Unix()
	bu.Update_date = time.Now().Unix()
	dao.InsertBu(bu)
	fmt.Printf(">>>>", bu)
	c.JSON(200, gin.H{
		"name": bu.Name,
		"active": bu.Active,
		"create_date": time.Now(),
	})
}

func FindBuByName(name string,c *gin.Context)  {

	result := dao.FindBuByName(name)
	fmt.Printf(">>>",result)
	c.JSON(200, gin.H{
		"bu": result ,
	})

}

func FindAllBu(c *gin.Context)  {

	results :=dao.FindAllBuDao(c)
	c.JSON(200, gin.H{
		"bu": results ,
	})

}

func UpdateBuActive(bu model.Bu) (model.Bu, error){

	bu_name := bson.M{"name":bu.Name}
	bu_active := bson.M{"$set":bson.M{"active":bu.Active}}

	session :=  dao.IntialDB()
	defer session.Close()
	collection := session.DB("workshop").C("bu")
	defer session.Close()
	//Update bu
	err := collection.Update(bu_name,bu_active)
	//Check null
	if err != nil {
		log.Fatal(err)
		return model.Bu{},err
	}else {
		result := dao.FindBuByName(bu.Name)
		return result, nil
	}
}

func InsertSquadToBu(squadBu  model.SquadBu) (model.Bu,error){
	bu_result := dao.FindBuByName(squadBu.Bu_name)
	if bu_result.Active == true {
		return model.Bu{}, nil
	} else {
		bu_name := bson.M{"name":bu_result.Name}

		squad_result := dao.FindSquadByName(squadBu.Squad_name)
		//squad := bson.M{"$set":bson.M{"squads":squad_result}}
		//Set Bu
		bu_result.Squads = append(bu_result.Squads, squad_result)
		bu_result.Update_date = time.Now().Unix()
		bu_result.Active = true
		fmt.Printf(">>>>>>", squad_result)
		//Connect Bu DB
		session := dao.IntialDB()
		defer session.Close()
		collection := session.DB("workshop").C("bu")
		defer session.Close()
		//Update bu
		err := collection.Update(bu_name, bu_result)
		//Check null
		if err != nil {
			log.Fatal(err)
			return model.Bu{}, err
		} else {
			result := dao.FindBuByName(squadBu.Bu_name)
			return result, nil
		}
	}
}

func BuDeacivate(buDeactive model.BuDeactive) (model.Bu, error) {

	bu_result := dao.FindBuByName(buDeactive.Bu_name)
	bu_name := bson.M{"name": bu_result.Name}
	if bu_result.Active == false {
		return model.Bu{}, nil
	} else {

		for i := 0; i < len(bu_result.Squads); i++ {
			fmt.Println("<<<Deactive>>>", bu_result.Squads[i])
			fmt.Println("squad name>>>>>",bu_result.Squads[i].Name)
			squad_result := dao.FindSquadByName(bu_result.Squads[i].Name)
			fmt.Println("squad name result>>>>>",squad_result.Name)
			//Deactivate Squad
			SquadDeacivateByName(squad_result)
		}

		bu_result.Squads = nil
		bu_result.Active = false
		bu_result.Update_date = time.Now().Unix()

		session := dao.IntialDB()
		defer session.Close()
		collection := session.DB("workshop").C("bu")
		//Update bu
		err := collection.Update(bu_name, bu_result)
		//Check null
		if err != nil {
			log.Fatal(err)
			return model.Bu{}, err
		} else {
			result := dao.FindBuByName(buDeactive.Bu_name)
			return result, nil
		}
	}
}