package service

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/stretchr/testify/assert"
	"workshop.test/model"
	"workshop.test/dao"
)

func Test_squadService_InsertSquad_Success(t *testing.T){
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	_,err := InsertSquad(squad)
	assert := assert.New(t)
	assert.Equal(nil, err)

	/*	query := bson.M{"Squad_Name": squad_result.name}
		dao.FindOne("squad", query, &squad_result)
		assert.Equal(squad_result, squad)*/
}

func Test_squadService_InsertSquad_Duplicate(t *testing.T){
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	_,err := InsertSquad(squad)

	assert := assert.New(t)
	assert.Error(err)
}


/*func Test_squadService_DaoInsertSquad_fail(t *testing.T){
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	squad_result,err := InsertSquad(squad)

	assert := assert.New(t)
	assert.Equal(nil, err)

	query := bson.M{"squad_name": squad_result.name}
	dao.FindOne("squad", query, &squad_result)

	assert.Equal(squad_result, squad)
}*/

func Test_squadService_FindByName_success(t *testing.T)  {
	//mock
	name := "squad"
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	squad.Active = false

	//call service
	_,err := FindSquadByName(name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	/*assert.Equal(squad,squad_result)*/
}

func Test_squadService_FindByName_NotSuccess(t *testing.T)  {
	//mock
	name := "squad_notFound"
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	squad.Active = false

	//call service
	squad_result,err := FindSquadByName(name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.NotEqual(squad,squad_result)
}


func Test_squadService_FindAll_success(t *testing.T)  {
	//mock
	name := "squad"
	var squad model.Squad
	squad.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	squad.Name = "squad"
	squad.Active = false

	var squad2 model.Squad
	squad2.Id = bson.ObjectIdHex("52ec1197123f02deb2e91e0d")
	squad2.Name = "squad2"
	InsertSquad(squad2)

	//call service
	var squad_result []model.Squad
	err := findAll(name,&squad_result)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(squad.Name,squad_result[0].Name)
	assert.Equal(squad.Id,squad_result[0].Id)
	assert.Equal(squad.Active,squad_result[0].Active)
	assert.Equal(squad2.Name,squad_result[1].Name)
	assert.Equal(squad2.Id,squad_result[1].Id)
	assert.Equal(squad2.Active,squad_result[1].Active)
}

func Test_squadService_FindAll_Notsuccess(t *testing.T)  {
	//mock
	name := "squadNotFound"

	//call service
	var squad_result []model.Squad
	err := findAll(name,&squad_result)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal([]model.Squad(nil),squad_result)
}


func Test_squadService_InsertDevToSquad_Success(t *testing.T) {
	//mock dev
	var dev_insert model.Dev
	dev_insert.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_insert.Dev_name = "dev"
	dev_insert.DevId = "devid"
	_,err := InsertDev(dev_insert)

	//mock
	dev,err:=FindDevByName("dev")
	dev.Active = true
	dev_name := "dev"
	squad_name := "squad"

	//call service
	squad,err:=InsertDevToSquad(dev_name,squad_name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(dev.Active,squad.Devs[0].Active)
	assert.Equal(dev.Dev_name,squad.Devs[0].Dev_name)
	assert.Equal(dev.Id,squad.Devs[0].Id)
	assert.Equal(dev.Create_date,squad.Devs[0].Create_date)
}

func Test_squadService_InsertDevToSquad_DevNotFound(t *testing.T) {
	//mock
	dev_name := "dev_notFound"
	squad_name := "squad"

	//call service
	squad,err:=InsertDevToSquad(dev_name,squad_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Squad{},squad)
}

func Test_squadService_InsertDevToSquad_SquadNotFound(t *testing.T) {
	//mock
	dev_name := "dev"
	squad_name := "squad_notFound"


	//call service
	squad,err:=InsertDevToSquad(dev_name,squad_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Squad{},squad)
}

func Test_squadService_InsertDevToSquad_DuplicateDev(t *testing.T) {
	//mock
	dev_name := "dev"
	squad_name := "squad"

	//call service
	squad,err:=InsertDevToSquad(dev_name,squad_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Squad{},squad)
}


func Test_squadService_DeactiveSquad_findOneSquad_fail(t *testing.T) {
	//mock
	squad_name := "squad_notFound"

	//call service
	squad,err:=DeactiveSquad(squad_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Squad{},squad)
}

func Test_squadService_DeactiveSquad_findOneDev_fail(t *testing.T) {
	//mock
	dev,err:=FindDevByName("dev")
	dev_query := bson.M{"dev_name": dev.Dev_name}
	dev.Dev_name = "dev_notFound"
	update("dev",dev_query,&dev)
	squad_name := "squad"

	//call service
	squad,err:=DeactiveSquad(squad_name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal(model.Squad{},squad)

	//reset
	dev_reset,_:=FindDevByName(dev.Dev_name)
	dev_reset.Dev_name = "dev"
	dev_query = bson.M{"dev_name": dev.Dev_name}
	update("dev",dev_query,&dev_reset)
}

func Test_squadService_DeactiveSquad_Success(t *testing.T) {
	//mock
	dev,err:=FindDevByName("dev")
	dev.Active = false
	squad_name := "squad"

	//call service
	squad,err:=DeactiveSquad(squad_name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(false,squad.Active)
	assert.Equal(false,dev.Active)
}


func Test_squadService_UpdateSquadActiveTrue_Success(t *testing.T) {
	//mock
	var squad model.Squad
	squad.Name = "squad"

	//call dao
	squad_result,err := UpdateSquadActiveTrue(squad)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(true,squad_result.Active)

}


func Test_squadService_UpdateSquadActiveTrue_NotSuccess(t *testing.T) {
	//mock
	var squad model.Squad
	squad.Name = "squad_notFound"

	//call dao
	_,err := UpdateSquadActiveTrue(squad)
	assert := assert.New(t)
	assert.Error(err)

}

func Test_squadService_UpdateSquadActiveFalse_Success(t *testing.T) {
	//mock
	var squad model.Squad
	squad.Name = "squad"

	//call dao
	squad_result,err := UpdateSquadActiveFalse(squad)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(false,squad_result.Active)

}

func Test_squadService_UpdateSquadActiveFalse_NotSuccess(t *testing.T) {
	//mock
	var squad model.Squad
	squad.Name = "squad_notFound"

	//call dao
	_,err := UpdateSquadActiveFalse(squad)
	assert := assert.New(t)
	assert.Error(err)

	//reset database
	session,_ := dao.InitialDB()
	session.DB("workshop").C("squad").RemoveAll(nil)
	session.DB("workshop").C("dev").RemoveAll(nil)
	defer session.Close()
}


