package main

import (
	"fmt"
	"workshop.test/model"
	"gopkg.in/mgo.v2/bson"
	"workshop.test/dao"
	"workshop.test/service"
	"workshop.test/route"
)

func main() {
	//DaoInsertDev()
	route.Route()

	//service.InsertSquadToBu("squad","bu")

	//var result model.Dev

	//>>>>Dao Insert Dev<<<<
	///DaoInsertDev

	//>>>>Dao Find One Dev<<<<
	//DaoFindOneDev

	//>>>>Dao Update Dev<<<<
	//DaoUpdateDev

	//>>>>Dao Find All Dev<<<<
	//DaoFindAllDev

	//1
	//>>>>Service Insert Dev Service<<<<
	//ServiceInsert2Devs()

	//2
	//>>>>Service FindByName Dev<<<<
	//ServiceFindByNameDev()

	//>>>>>Dao Squad service insert<<<<<<
	//DaoInsertSquad()

/*	squad,err:=service.InsertDevToSquad("dev","squad")
	fmt.Println("suqad insert>>>>>>",squad,"\nerror>>>>>>>",err)*/
}

func DaoInsertDev()  {
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Create_date = 1475088791
	dev.Update_date = 1475088851
	dev.Active = false

	dao.Insert("dev", &dev)
	fmt.Printf(">>>>>", dev)
}

func DaoFindOneDev()  {
	var dev model.Dev
	query := bson.M{"dev_name": "dev"}
	err := dao.FindOne("dev", query, &dev)
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>>>>", dev)
}

func DaoUpdateDev()  {
	var dev_find model.Dev
	query_find := bson.M{"dev_name": "dev"}
	err := dao.FindOne("dev", query_find, &dev_find)
	if err != nil {
		panic(err)
	}

	dev_find.Active = true
	query_update := bson.M{"dev_name": "dev"}
	err = dao.Update("dev", query_update, &dev_find)
	if err != nil {	panic(err)}
	fmt.Println(">>>>>", dev_find)

}

/*func DaoFindAllDev()  {
	var dev_findAll []model.Dev
	_:=dao.FindAll("dev",dev_findAll)
	fmt.Println(">>>>>FindAll",dev_findAll)
}*/

func ServiceInsert2Devs()  {
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"

	dev_result,err:=service.InsertDev(dev)
	fmt.Println("dev>>>>>", dev_result,err)

	var dev2 model.Dev
	dev2.Id = bson.ObjectIdHex("56ec1197123f02deb2e91e0d")
	dev2.Dev_name = "dev2"
	dev2.DevId = "devid2"

	dev_result2,err:=service.InsertDev(dev2)
	fmt.Println("dev2>>>>>", dev_result2,err)
}

func ServiceFindByNameDev()  {
	name := "dev"
	dev_result3,err := service.FindDevByName(name)
	fmt.Println(">>>>>Find by name",dev_result3,err)

}

//>>>>>>>>>>>>>>>Squad<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func DaoInsertSquad()  {
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
/*	dev.Create_date = 1475088791
	dev.Update_date = 1475088851*/
	squad.Active = false

	dao.Insert("squad", &squad)
	fmt.Printf(">>>>>", squad)
}

func DaoFindOneSquad()  {
/*	var dev model.Dev
	query := bson.M{"dev_name": "dev"}
	err := dao.FindOne("dev", query, &dev)
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>>>>", dev)*/
}

func DaoUpdateSquad()  {
/*	var dev_find model.Dev
	query_find := bson.M{"dev_name": "dev"}
	err := dao.FindOne("dev", query_find, &dev_find)
	if err != nil {
		panic(err)
	}

	dev_find.Active = true
	query_update := bson.M{"dev_name": "dev"}
	err = dao.Update("dev", query_update, &dev_find)
	if err != nil {	panic(err)}
	fmt.Println(">>>>>", dev_find)*/

}

/*func DaoFindAllDev()  {
	var dev_findAll []model.Dev
	_:=dao.FindAll("dev",dev_findAll)
	fmt.Println(">>>>>FindAll",dev_findAll)
}*/

func ServiceInsert2Squads()  {
	/*var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"

	dev_result,err:=service.InsertDev(dev)
	fmt.Println("dev>>>>>", dev_result,err)

	var dev2 model.Dev
	dev2.Id = bson.ObjectIdHex("56ec1197123f02deb2e91e0d")
	dev2.Dev_name = "dev2"
	dev2.DevId = "devid2"

	dev_result2,err:=service.InsertDev(dev2)
	fmt.Println("dev2>>>>>", dev_result2,err)*/
}

func ServiceFindByNameSquad()  {
	//name := "dev"
	//dev_result3,err := service.FindDevByName(name)
	//fmt.Println(">>>>>Find by name",dev_result3,err)

}
