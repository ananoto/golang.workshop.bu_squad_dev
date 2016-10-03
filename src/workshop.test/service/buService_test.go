package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"gopkg.in/mgo.v2/bson"
	"workshop.test/model"
	"workshop.test/dao"
)

func Test_buService_InsertBu_Success(t *testing.T){
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	_,err := InsertBu(bu)
	assert := assert.New(t)
	assert.Equal(nil, err)

	/*	query := bson.M{"Bu_Name": bu_result.name}
		dao.FindOne("bu", query, &bu_result)
		assert.Equal(bu_result, bu)*/
}

func Test_buService_InsertBu_Duplicate(t *testing.T){
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	_,err := InsertBu(bu)

	assert := assert.New(t)
	assert.Error(err)
}



/*func Test_buService_DaoInsertBu_fail(t *testing.T){
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	bu_result,err := InsertBu(bu)

	assert := assert.New(t)
	assert.Equal(nil, err)

	query := bson.M{"bu_name": bu_result.name}
	dao.FindOne("bu", query, &bu_result)

	assert.Equal(bu_result, bu)
}*/

func Test_buService_FindByName_success(t *testing.T)  {
	//mock
	name := "bu"
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	bu.Active = false

	//call service
	_,err := FindBuByName(name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	/*assert.Equal(bu,bu_result)*/
}

func Test_buService_FindByName_NotSuccess(t *testing.T)  {
	//mock
	name := "bu_notFound"
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	bu.Active = false

	//call service
	bu_result,err := FindBuByName(name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.NotEqual(bu,bu_result)
}

func Test_buService_FindAll_success(t *testing.T)  {
	//mock
	name := "bu"
	var bu model.Bu
	bu.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	bu.Name = "bu"
	bu.Active = false

	var bu2 model.Bu
	bu2.Id = bson.ObjectIdHex("52ec1197123f02deb2e91e0d")
	bu2.Name = "bu2"
	InsertBu(bu2)

	//call service
	var bu_result []model.Bu
	err := findAll(name,&bu_result)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(bu.Name,bu_result[0].Name)
	assert.Equal(bu.Id,bu_result[0].Id)
	assert.Equal(bu.Active,bu_result[0].Active)
	assert.Equal(bu2.Name,bu_result[1].Name)
	assert.Equal(bu2.Id,bu_result[1].Id)
	assert.Equal(bu2.Active,bu_result[1].Active)
}

func Test_buService_FindAll_Notsuccess(t *testing.T)  {
	//mock
	name := "budNotFound"

	//call service
	var bu_result []model.Bu
	err := findAll(name,&bu_result)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal([]model.Bu(nil),bu_result)
}


func Test_buService_InsertSquadToBu_Success(t *testing.T) {
	//mock squad
	var squad_insert model.Squad
	squad_insert.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad_insert.Name = "squad"
	_,err := InsertSquad(squad_insert)

	//mock
	squad,err:=FindSquadByName("squad")
	squad.Active = true
	squad_name := "squad"
	bu_name := "bu"

	//call service
	bu,err:=InsertSquadToBu(squad_name,bu_name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(squad.Active,bu.Squads[0].Active)
	assert.Equal(squad.Name,bu.Squads[0].Name)
	assert.Equal(squad.Id,bu.Squads[0].Id)
	assert.Equal(squad.Create_date,bu.Squads[0].Create_date)
}

func Test_buService_FindBuBySquad_Success(t *testing.T) {
	//mock
	squad_name := "squad"
	bu_dump,_ := FindBuByName("bu")
	//call
	bu,err := FindBuBySquad(squad_name)

	//test
	assert := assert.New(t)
	assert.Equal(bu_dump,bu)
	assert.Equal(nil,err)

}

func Test_buService_FindBuBySquad_NotSuccess(t *testing.T) {
	//mock
	squad_name := "squad"
	bu_dump,_ := FindBuByName("bu")
	//call
	bu,err := FindBuBySquad(squad_name)

	//test
	assert := assert.New(t)
	assert.Equal(bu_dump,bu)
	assert.Equal(nil,err)

}

func Test_buService_InsertDevToSquadToBu_Success(t *testing.T) {
	assert := assert.New(t)
	//mock dev
	var dev_insert model.Dev
	dev_insert.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_insert.Dev_name = "dev"
	dev_insert.DevId = "devid"
	dev,err := InsertDev(dev_insert)
	dev,err=FindDevByName(dev.Dev_name)
	assert.Nil(err)

	bu,err:=FindBuByName("bu")
	assert.Nil(err)
	assert.Equal(true,bu.Active)
	assert.Equal(true,bu.Squads[0].Active)

	bu,err=DeactiveBu("bu")
	assert.Equal(false,bu.Active)

	squad,err:=FindSquadByName("squad")
	assert.Equal(false,squad.Active)

	squad,err=InsertDevToSquad(dev.Dev_name,squad.Name)
	assert.Equal(true,squad.Active)
	assert.Equal(true,squad.Devs[0].Active)

	//mock
	dev,err=FindDevByName(dev.Dev_name)
	squad,err=FindSquadByName("squad")
	squad_name := squad.Name
	bu_name := "bu"

	//call service
	bu,err=InsertSquadToBu(squad_name,bu_name)

	//test
	assert.Equal(nil,err)
	assert.Equal(squad.Active,bu.Squads[0].Active)
	assert.Equal(squad.Name,bu.Squads[0].Name)
	assert.Equal(squad.Id,bu.Squads[0].Id)
	assert.Equal(squad.Create_date,bu.Squads[0].Create_date)
	assert.Equal(squad.Devs[0].Active,dev.Active)
}

func Test_buService_InsertSquadToBu_SquadNotFound(t *testing.T) {
	//mock
	squad_name := "squad_notFound"
	bu_name := "bu"

	//call service
	bu,err:=InsertSquadToBu(squad_name,bu_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Bu{},bu)
}

func Test_buService_InsertSquadToBu_BuNotFound(t *testing.T) {
	//mock
	squad_name := "squad"
	bu_name := "bu_notFound"


	//call service
	bu,err:=InsertSquadToBu(squad_name,bu_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Bu{},bu)
}
/*

func Test_buService_InsertSquadToBu_DuplicateSquad(t *testing.T) {
	//mock
	squad_name := "squad"
	bu_name := "bu"

	//call service
	bu,err:=InsertSquadToBu(squad_name,bu_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Bu{},bu)
}

*/

func Test_buService_DeactiveBu_findOneBu_fail(t *testing.T) {
	//mock
	bu_name := "bu_notFound"

	//call service
	bu,err:=DeactiveBu(bu_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Bu{},bu)
}

func Test_buService_DeactiveBu_findOneSquad_fail(t *testing.T) {
	//mock
	squad,err:=FindSquadByName("squad")
	squad_query := bson.M{"name": squad.Name}
	squad.Name = "squad_notFound"
	update("squad",squad_query,&squad)
	bu_name := "bu"

	//call service
	bu,err:=DeactiveBu(bu_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Bu{},bu)

	//reset
	squad_reset,_:=FindSquadByName(squad.Name)
	squad_reset.Name = "squad"
	squad_query = bson.M{"name": squad.Name}
	update("squad",squad_query,&squad_reset)
}

func Test_buService_DeactiveBu_Success(t *testing.T) {
	//mock dev
/*	var dev_insert model.Dev
	dev_insert.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_insert.Dev_name = "dev"
	dev_insert.DevId = "devid"
	_,err := InsertDev(dev_insert)*/
	dev,err:=FindDevByName("dev")

	//mock squad
	squad,err:=FindSquadByName("squad")
	squad.Active = false
	bu_name := "bu"

	//call service
	bu,err:=DeactiveBu(bu_name)

	//mock
	dev,_=FindDevByName("dev")

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(false,bu.Active)
	assert.Equal(false,squad.Active)
	assert.Equal(false,dev.Active)
}


func Test_buService_UpdateBuActiveTrue_Success(t *testing.T) {
	//mock
	var bu model.Bu
	bu.Name = "bu"

	//call dao
	bu_result,err := UpdateBuActiveTrue(bu)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(true,bu_result.Active)

}


func Test_buService_UpdateBuActiveTrue_NotSuccess(t *testing.T) {
	//mock
	var bu model.Bu
	bu.Name = "bu_notFound"

	//call dao
	_,err := UpdateBuActiveTrue(bu)
	assert := assert.New(t)
	assert.Error(err)

}

func Test_buService_UpdateBuActiveFalse_Success(t *testing.T) {
	//mock
	var bu model.Bu
	bu.Name = "bu"

	//call dao
	bu_result,err := UpdateBuActiveFalse(bu)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(false,bu_result.Active)

}

func Test_buService_UpdateBuActiveFalse_NotSuccess(t *testing.T) {
	//mock
	var bu model.Bu
	bu.Name = "bu_notFound"

	//call dao
	_,err := UpdateBuActiveFalse(bu)
	assert := assert.New(t)
	assert.Error(err)

	//reset database
	session,_ := dao.InitialDB()
	session.DB("workshop").C("bu").RemoveAll(nil)
	session.DB("workshop").C("squad").RemoveAll(nil)
	session.DB("workshop").C("dev").RemoveAll(nil)
	defer session.Close()

}
