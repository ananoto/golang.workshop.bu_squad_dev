package service

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"workshop.test/model"
	"time"
	"errors"
)


func InsertSquad(squad model.Squad) (model.Squad,error){
	//set query
	query := bson.M{"name": squad.Name}
	//check duplicate
	err:=findOne("squad", query, &squad)
	if err == nil{
		fmt.Println("----->Duplicate na<-----")
		return model.Squad{},errors.New("Duplicate Dev, service Can't insert!!")
	}
	//set
	squad.Active = false
	squad.Create_date = time.Now().Unix()
	squad.Update_date = time.Now().Unix()
	//insert
	err=insert("squad", &squad)
	/*	if err != nil {
			fmt.Println("----->Can't insert na<-----")
			return model.Squad{},err}*/

	//check data
	err=findOne("squad", query, &squad)
	/*	if err != nil {
			fmt.Println("----->Can't find na<-----")
			return model.Squad{},err}*/
	fmt.Println("----->InserDev Success<-----")
	return squad,nil
}

func InsertDevToSquad(dev_name string, squad_name string) (model.Squad,error){

	//set query squad
	var squad model.Squad
	squad_query := bson.M{"name": squad_name}

	//check squad exist
	err:=findOne("squad", squad_query, &squad)
	if err != nil{
		fmt.Println("----->service Can't found Squad<-----")
		return model.Squad{},errors.New("service Can't found Squad!!")
	}

	//set query dev
	var dev model.Dev
	dev_query := bson.M{"dev_name": dev_name}

	//check dev exist
	err=findOne("dev", dev_query, &dev)
	if err != nil{
		fmt.Println("----->service Can't found Dev<-----")
		return model.Squad{},errors.New("service Can't found Dev!!")
	}

	//check dev active
	if dev.Active == true {
		fmt.Println("----->Duplicate Dev, Can't insert!!<-----")
		return model.Squad{},errors.New("Duplicate Dev, Can't insert!!")
	}
	//set dev
	dev.Active = true
	dev.Update_date = time.Now().Unix()

	//set squad
	squad.Devs = append(squad.Devs, dev)
	squad.Active = true
	squad.Update_date = time.Now().Unix()

	//update suqad & dev
	update("squad", squad_query ,&squad)
	update("dev",dev_query,&dev)
	return squad, nil
}

func FindSquadByName(name string) (model.Squad,error) {
	//set query
	var squad model.Squad
	query := bson.M{"name": name}

	//find squad
	err:=findOne("squad",query, &squad)
	if err != nil {
		return model.Squad{},errors.New("service can't find Squad!!")}
	return squad,nil
}


func FindAllSquad() ([]model.Squad,error) {
	//set query
	var squad []model.Squad

	//find dev
	err:=findAll("squad",&squad)
	if err != nil {
		return []model.Squad{},errors.New("service Can't findAll Squad!!")}
	return squad,nil
}


func DeactiveSquad(squad_name string)  (model.Squad,error){
	//set query squad
	var squad model.Squad
	squad_query := bson.M{"name": squad_name}

	//check squad exist
	err:=findOne("squad", squad_query, &squad)
	if err != nil{
		fmt.Println("----->service Can't found Squad<-----")
		return model.Squad{},errors.New("service Can't found Squad!!")
	}

	for i := 0; i < len(squad.Devs); i++ {
		fmt.Println("<<<Deactive>>>", squad.Devs[i])
		dev,err := FindDevByName(squad.Devs[i].Dev_name)
		if err != nil{
			fmt.Println("----->service Can't found Dev<-----")
			return model.Squad{},errors.New("service Can't found dev!!")
		}
		dev_query := bson.M{"dev_name": dev.Dev_name}
		dev.Active = false
		dev.Update_date = time.Now().Unix()
		//Update bu
		update("dev", dev_query, dev)
	}

	//set squad
	squad.Devs = nil
	squad.Active = false
	squad.Update_date = time.Now().Unix()

	//update suqad & dev
	update("squad", squad_query ,&squad)
	return squad, nil
}

func UpdateSquadActiveTrue(squad model.Squad) (model.Squad,error){
	//set query
	query := bson.M{"name": squad.Name}

	//check existing squad
	err:=findOne("squad", query, &squad)

	//set
	squad.Active = true
	squad.Update_date = time.Now().Unix()
	//insert
	err=update("squad", query ,&squad)

	return squad,err
}

func UpdateSquadActiveFalse(squad model.Squad) (model.Squad,error){
	//set query
	query := bson.M{"name": squad.Name}

	//check existing squad
	err:=findOne("squad", query, &squad)

	//set
	squad.Active = false
	squad.Update_date = time.Now().Unix()
	//insert
	err=update("squad", query ,&squad)

	return squad,err
}
