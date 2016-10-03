package service

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"errors"
	"time"
	"workshop.test/model"
)

func InsertBu(bu model.Bu) (model.Bu,error){
	//set query
	query := bson.M{"name": bu.Name}
	//check duplicate
	err:=findOne("bu", query, &bu)
	if err == nil{
		fmt.Println("----->Duplicate na<-----")
		return model.Bu{},errors.New("Duplicate Squad, service Can't insert!!")
	}
	//set
	bu.Active = false
	bu.Create_date = time.Now().Unix()
	bu.Update_date = time.Now().Unix()
	//insert
	err=insert("bu", &bu)
	/*	if err != nil {
			fmt.Println("----->Can't insert na<-----")
			return model.Bu{},err}*/

	//check data
	err=findOne("bu", query, &bu)
	/*	if err != nil {
			fmt.Println("----->Can't find na<-----")
			return model.Bu{},err}*/
	return bu,nil
}

func InsertSquadToBu(squad_name string, bu_name string) (model.Bu,error){

	//set query bu
	var bu model.Bu
	bu_query := bson.M{"name": bu_name}

	//check bu exist
	err:=findOne("bu", bu_query, &bu)
	if err != nil{
		fmt.Println("----->service Can't found Bu<-----")
		return model.Bu{},errors.New("Can't found Bu!!")
	}

	//set query squad
	var squad model.Squad
	squad_query := bson.M{"name": squad_name}

	//check squad exist
	err=findOne("squad", squad_query, &squad)
	if err != nil{
		fmt.Println("----->service Can't found Squad<-----")
		return model.Bu{},errors.New("service Can't found Squad!!")
	}
	//set Squad
	squad.Active = true
	squad.Update_date = time.Now().Unix()

	//set bu
	bu.Squads = append(bu.Squads, squad)
	bu.Active = true
	bu.Update_date = time.Now().Unix()

	//update suqad & squad
	update("bu", bu_query ,&bu)
	update("squad",squad_query,&squad)
	return bu, nil
}

func FindBuByName(name string) (model.Bu,error) {
	//set query
	var bu model.Bu
	query := bson.M{"name": name}

	//find bu
	err:=findOne("bu",query, &bu)
	if err != nil {
		return model.Bu{},errors.New("service can't find Bu!!")}
	return bu,nil
}

func FindBuBySquad(squad_name string) (model.Bu,error) {
	bu,err:=FindAllBu()
	for i := 0; i < len(bu); i++{
		for j := 0; j < len(bu[i].Squads); j++{
			if bu[i].Squads[j].Name == squad_name{
				//find bu
				//set query
				var bu_result model.Bu
				query_bu := bson.M{"name": bu[i].Name}
				err=findOne("bu",query_bu, &bu_result)
				if err != nil {
					return model.Bu{},errors.New("service can't find Bu!!")}
				return bu_result,nil
			}
		}
	}

	return model.Bu{},errors.New("service can't find Bu by squad!!")


}


func FindAllBu() ([]model.Bu,error) {
	//set query
	var bu []model.Bu

	//find dev
	err:=findAll("bu",&bu)
	if err != nil {
		return []model.Bu{},errors.New("service Can't findAll Bu!!")}
	return bu,nil
}

func DeactiveBu(bu_name string)  (model.Bu,error){
	//set query bu
	var bu model.Bu
	bu_query := bson.M{"name": bu_name}

	//check bu exist
	err:=findOne("bu", bu_query, &bu)
	if err != nil{
		fmt.Println("----->Can't found Bu<-----")
		return model.Bu{},errors.New("service Can't found Bu!!")
	}

	for i := 0; i < len(bu.Squads); i++ {
		_,err=DeactiveSquad(bu.Squads[i].Name)
		if err != nil{
			fmt.Println("----->Can't found Squad<-----")
			return model.Bu{},errors.New("service Can't found squad!!")
		}
		/*fmt.Println("<<<Deactive>>>", bu.Squads[i])
		squad,err := FindSquadByName(bu.Squads[i].Name)
		if err != nil{
			fmt.Println("----->Can't found Squad<-----")
			return model.Bu{},errors.New("service Can't found squad!!")
		}
		squad_query := bson.M{"squad_name": squad.Name}
		squad.Active = false
		squad.Update_date = time.Now().Unix()
		//Update bu
		update("squad", squad_query, squad)*/
	}

	//set bu
	bu.Squads = nil
	bu.Active = false
	bu.Update_date = time.Now().Unix()

	//update suqad & squad
	update("bu", bu_query ,&bu)
	return bu, nil
}

func UpdateBuActiveTrue(bu model.Bu) (model.Bu,error){
	//set query
	query := bson.M{"name": bu.Name}

	//check existing bu
	err:=findOne("bu", query, &bu)

	//set
	bu.Active = true
	bu.Update_date = time.Now().Unix()
	//insert
	err=update("bu", query ,&bu)

	return bu,err
}

func UpdateBuActiveFalse(bu model.Bu) (model.Bu,error){
	//set query
	query := bson.M{"name": bu.Name}

	//check existing bu
	err:=findOne("bu", query, &bu)

	//set
	bu.Active = false
	bu.Update_date = time.Now().Unix()
	//insert
	err=update("bu", query ,&bu)

	return bu,err
}
