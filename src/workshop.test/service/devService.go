package service

import (
	"workshop.test/model"
	"workshop.test/dao"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
	"errors"
)
var (
	insert = dao.Insert
	findOne = dao.FindOne
	update = dao.Update
	findAll = dao.FindAll
)


func InsertDev(dev model.Dev) (model.Dev,error){
	//set query
	query := bson.M{"dev_name": dev.Dev_name}
	//check duplicate
	err:=findOne("dev", query, &dev)
	if err == nil{
		fmt.Println("----->Duplicate na<-----")
		return model.Dev{},errors.New("Duplicate Dev, service Can't insert!!")
	}
	//set
	dev.Active = false
	dev.Create_date = time.Now().Unix()
	dev.Update_date = time.Now().Unix()
	//insert
	insert("dev", &dev)
/*	if err != nil {
		fmt.Println("----->Can't insert na<-----")
		return model.Dev{},err}*/

	//check data
	findOne("dev", query, &dev)
/*	if err != nil {
		fmt.Println("----->Can't find na<-----")
		return model.Dev{},err}*/
	fmt.Println("----->service InserDev Success<-----")
	return dev,nil
}

func FindDevByName(name string) (model.Dev,error) {
	//set query
	var dev model.Dev
	query := bson.M{"dev_name": name}

	//find dev
	err:=findOne("dev",query, &dev)
	if err != nil {
		return model.Dev{},errors.New("service Can't find Dev!!")}
	return dev,nil
}

func FindAllDev() ([]model.Dev,error) {
	//set query
	var dev []model.Dev

	//find dev
	err:=findAll("dev",&dev)
	if err != nil {
		return []model.Dev{},errors.New("service Can't findAll Dev!!")}
	return dev,nil
}

func UpdateDevActiveTrue(dev model.Dev) (model.Dev,error){
	//set query
	query := bson.M{"dev_name": dev.Dev_name}

	//check existing dev
	err:=findOne("dev", query, &dev)

	//set
	dev.Active = true
	dev.Update_date = time.Now().Unix()
	//insert
	err=update("dev", query ,&dev)

	return dev,err
}

func UpdateDevActiveFalse(dev model.Dev) (model.Dev,error){
	//set query
	query := bson.M{"dev_name": dev.Dev_name}

	//check existing dev
	err:=findOne("dev", query, &dev)

	//set
	dev.Active = false
	dev.Update_date = time.Now().Unix()
	//insert
	err=update("dev", query ,&dev)

	return dev,err
}