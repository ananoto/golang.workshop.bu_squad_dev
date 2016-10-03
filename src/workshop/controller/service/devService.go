package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"workshop/controller/dao"
	"workshop/model"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func InsertDev(dev model.Dev, c *gin.Context) {
	//Set
	dev.Active = false
	dev.Create_date = time.Now().Unix()
	dev.Update_date = time.Now().Unix()
	//Insert
	dao.InsertDev(dev)
	fmt.Printf(">>>>", dev)
	c.JSON(200, gin.H{
		"devId":       dev.DevId,
		"dev_name":    dev.Dev_name,
		"active":      dev.Active,
		"create_date": time.Now(),
	})
}

func UpdateDevActive(dev model.Dev) (model.Dev, error){

	dev_name := bson.M{"dev_name":dev.Dev_name}
	dev_active := bson.M{"$set":bson.M{"active":dev.Active}}

	session :=  dao.IntialDB()
	defer session.Close()
	collection := session.DB("workshop").C("dev")
	defer session.Close()
	//Update bu
	err := collection.Update(dev_name,dev_active)
	//Check null
	if err != nil {
		log.Fatal(err)
		return model.Dev{},err
	}else {
		result := dao.FindDevByName(dev.Dev_name)
		return result, nil
	}
}
