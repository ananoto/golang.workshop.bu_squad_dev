package service

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/stretchr/testify/assert"
	"workshop.test/model"
	"workshop.test/dao"
)

func Test_devService_InsertDev_Success(t *testing.T){
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	_,err := InsertDev(dev)
	assert := assert.New(t)
	assert.Equal(nil, err)

/*	query := bson.M{"dev_name": dev_result.Dev_name}
	dao.FindOne("dev", query, &dev_result)
	assert.Equal(dev_result, dev)*/
}

func Test_devService_InsertDev_Duplicate(t *testing.T){
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	_,err := InsertDev(dev)

	assert := assert.New(t)
	assert.Error(err)
}

/*func Test_devService_DaoInsertDev_fail(t *testing.T){
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev_result,err := InsertDev(dev)

	assert := assert.New(t)
	assert.Equal(nil, err)

	query := bson.M{"dev_name": dev_result.Dev_name}
	dao.FindOne("dev", query, &dev_result)

	assert.Equal(dev_result, dev)
}*/

func Test_devService_FindByName_success(t *testing.T)  {
	//mock
	name := "dev"
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Active = false

	//call service
	dev_result,err := FindDevByName(name)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(dev.Dev_name,dev_result.Dev_name)
	assert.Equal(dev.Id,dev_result.Id)
	assert.Equal(dev.Active,dev_result.Active)
}

func Test_devService_FindAll_success(t *testing.T)  {
	//mock
	name := "dev"
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Active = false

	var dev2 model.Dev
	dev2.Id = bson.ObjectIdHex("52ec1197123f02deb2e91e0d")
	dev2.Dev_name = "dev2"
	dev2.DevId = "devid2"
	InsertDev(dev2)

	//call service
	var dev_result []model.Dev
	err := findAll(name,&dev_result)

	//test
	assert := assert.New(t)
	assert.Equal(nil,err)
	assert.Equal(dev.Dev_name,dev_result[0].Dev_name)
	assert.Equal(dev.Id,dev_result[0].Id)
	assert.Equal(dev.Active,dev_result[0].Active)
	assert.Equal(dev2.Dev_name,dev_result[1].Dev_name)
	assert.Equal(dev2.Id,dev_result[1].Id)
	assert.Equal(dev2.Active,dev_result[1].Active)
}

func Test_devService_FindAll_Notsuccess(t *testing.T)  {
	//mock
	name := "devNotFound"

	//call service
	var dev_result []model.Dev
	err := findAll(name,&dev_result)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.Equal([]model.Dev(nil),dev_result)
}

func Test_devService_FindByName_NotSuccess(t *testing.T)  {
	//mock
	name := "dev_notFound"
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Active = false

	//call service
	dev_result,err := FindDevByName(name)

	//test
	assert := assert.New(t)
	assert.Error(err)
	assert.NotEqual(dev,dev_result)
}

func Test_devService_UpdateDevActiveTrue_Success(t *testing.T) {
	//mock
	var dev model.Dev
	dev.Dev_name = "dev"

	//call dao
	dev_result,err := UpdateDevActiveTrue(dev)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(true,dev_result.Active)

}

func Test_devService_UpdateDevActiveTrue_NotSuccess(t *testing.T) {
	//mock
	var dev model.Dev
	dev.Dev_name = "dev_notFound"

	//call dao
	_,err := UpdateDevActiveTrue(dev)
	assert := assert.New(t)
	assert.Error(err)

}

func Test_devService_UpdateDevActiveFalse_Success(t *testing.T) {
	//mock
	var dev model.Dev
	dev.Dev_name = "dev"

	//call dao
	dev_result,err := UpdateDevActiveFalse(dev)
	assert := assert.New(t)
	assert.Equal(nil, err)
	assert.Equal(false,dev_result.Active)

}

func Test_devService_UpdateDevActiveFalse_NotSuccess(t *testing.T) {
	//mock
	var dev model.Dev
	dev.Dev_name = "dev_notFound"

	//call dao
	_,err := UpdateDevActiveFalse(dev)
	assert := assert.New(t)
	assert.Error(err)

	//reset database
	session,_ := dao.InitialDB()
	session.DB("workshop").C("dev").RemoveAll(nil)
	defer session.Close()
}
