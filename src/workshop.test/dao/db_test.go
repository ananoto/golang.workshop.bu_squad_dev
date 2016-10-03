package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"workshop.test/model"
	"gopkg.in/mgo.v2/bson"/*
	"workshop/controller/dao"
	"fmt"*/
	"encoding/json"
	"errors"
)

func Test_InitialDB_Success(t *testing.T) {
	_, err := InitialDB()
	assert := assert.New(t)
	assert.Equal(nil,err)
}


func Test_Dao_InsertDev_Success(t *testing.T){
	//mock
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Create_date = 1475088791
	dev.Update_date = 1475088851
	dev.Active = false

	//call dao
	err:=Insert("dev", &dev)

	//test
	assert := assert.New(t)
	assert.Equal(nil, err)
/*	query := bson.M{"dev_name": dev.Dev_name}
	var dev_dump model.Dev
	FindOne("dev", query, &dev_dump)

	assert.Equal(dev_dump, dev)*/
}

func Test_Dao_InsertDev_CollectionInsertFail(t *testing.T){
	//mock
	var dev model.Dev
	dev.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev.Dev_name = "dev"
	dev.DevId = "devid"
	dev.Create_date = 1475088791
	dev.Update_date = 1475088851
	dev.Active = false

	//call dao
	err:=Insert("dev", &dev)

	//test
	assert := assert.New(t)
	assert.Error(err,"Error")
/*
	query := bson.M{"dev_name": dev.Dev_name}
	var dev_dump model.Dev
	FindOne("dev", query, &dev_dump)

	assert.Equal(dev_dump, dev)*/
}

func Test_Dao_FindOne_Dev_Success(t *testing.T){
	//mock
	var dev_dump model.Dev
	dev_dump.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_dump.Dev_name = "dev"
	dev_dump.DevId = "devid"
	dev_dump.Create_date = 1475088791
	dev_dump.Update_date = 1475088851
	dev_dump.Active = false
	var dev model.Dev
	query := bson.M{"dev_name": "dev"}

	//call dao
	FindOne("dev", query, &dev)

	//test
	assert := assert.New(t)
	assert.Equal(dev_dump, dev)
}

func Test_Dao_FindOne_Dev_CollectionFindFail(t *testing.T){
	//mock
	var dev model.Dev
	query := bson.M{"dev_name": "dev_notFound"}

	//call dao
	err:=FindOne("dev", query, &dev)

	//test
	assert := assert.New(t)
	assert.Error(err,"Error")
}

func Test_Dao_FindOne_Dev_MarshalFail(t *testing.T){
	//mock
	var dev model.Dev
	query := bson.M{"dev_name": "dev"}
	marshal = func(v interface{}) ([]byte,error){
		return nil,errors.New("booommm")}

	//call dao
	err:=FindOne("dev", query, &dev)

	//test
	assert := assert.New(t)
	assert.Error(err,"Error")

	//setting back
	marshal = func(v interface{}) ([]byte, error){
		return json.Marshal(v)	}

}

func Test_Dao_FindAll_Dev_Success(t *testing.T){
	var dev []model.Dev
	err := FindAll("dev",&dev)

	assert := assert.New(t)
	assert.Equal("dev",dev[0].Dev_name)
	assert.Equal(nil,err)
}

func Test_Dao_FindAll_Dev_NotSuccess(t *testing.T){
	var dev []model.Dev
	err:=FindAll("devNotFound",&dev)

	assert := assert.New(t)
	//assert.Equal(nil,err)
	assert.Error(err)
}

func Test_Dao_FindAll_Dev_MarshalFail(t *testing.T){
	//mock
	var dev []model.Dev
	marshal = func(v interface{}) ([]byte,error){
		return nil,errors.New("booommm")}

	//call dao
	err:=FindAll("dev", &dev)

	//test
	assert := assert.New(t)
	assert.Error(err,"Error")

	//setting back
	marshal = func(v interface{}) ([]byte, error){
		return json.Marshal(v)	}

}

func Test_Dao_Update_DevToSquad_Success(t *testing.T){
	//mock
	var dev_dump model.Dev
	dev_dump.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_dump.Dev_name = "dev"
	dev_dump.DevId = "devid"
	dev_dump.Create_date = 1475088791
	dev_dump.Update_date = 1475088851
	dev_dump.Active = false

	var dev_find model.Dev
	query_dev := bson.M{"dev_name": "dev"}
	var squad_find model.Squad
	query_squad := bson.M{"name":"squad"}

	//call dao
	FindOne("dev", query_dev, &dev_find)
	FindOne("squad",query_squad,&squad_find)

	squad_find.Devs = append(squad_find.Devs,dev_find)

	query_update := bson.M{"name": "squad"}
	Update("squad", query_update, &squad_find)

	//test
	assert := assert.New(t)
	assert.Equal(dev_dump, squad_find.Devs[0])
}

func Test_Dao_Update_Dev_Success(t *testing.T){
	//mock
	var dev_dump model.Dev
	dev_dump.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_dump.Dev_name = "dev"
	dev_dump.DevId = "devid"
	dev_dump.Create_date = 1475088791
	dev_dump.Update_date = 1475088851
	dev_dump.Active = true

	var dev_find model.Dev
	query_find := bson.M{"dev_name": "dev"}

	//call dao
	FindOne("dev", query_find, &dev_find)

	dev_find.Active = true
	query_update := bson.M{"dev_name": "dev"}
	Update("dev", query_update, &dev_find)

	//test
	assert := assert.New(t)
	assert.Equal(dev_dump, dev_find)
}

func Test_Dao_Update_Dev_CollectionUpdateFail(t *testing.T){
	//mock
	var dev_dump model.Dev
	dev_dump.Id = bson.ObjectIdHex("57ec1197123f02deb2e91e0d")
	dev_dump.Dev_name = "dev"
	dev_dump.DevId = "devid"
	dev_dump.Create_date = 1475088791
	dev_dump.Update_date = 1475088851
	dev_dump.Active = true

	var dev_find model.Dev
	query_find := bson.M{"dev_name": "dev"}
	dev_find.Active = true
	query_update := bson.M{"dev_name": "devNotFound"}

	//call dao
	FindOne("dev", query_find, &dev_find)
	err:=Update("dev", query_update, "devNotFound")

	assert := assert.New(t)
	assert.Error(err,"Error")

	//reset database
	session,_ := InitialDB()
	session.DB("workshop").C("dev").RemoveAll(nil)
	defer session.Close()
}

/*
func Test_InitialDB_NotSuccess(t *testing.T) {
	url = "bla bla"
	_, err := InitialDB()

	assert := assert.New(t)
	assert.NotNil(err)
}
*/


