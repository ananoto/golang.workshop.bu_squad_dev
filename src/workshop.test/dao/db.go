package dao

import (
//	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

var (
	url = "localhost:27017"
	dial = mgo.Dial
	marshal =  json.Marshal
	unmarshal = json.Unmarshal


)
/*func RessetAllDB(){
	session,_ := InitialDB()
	session.DB("workshop").C("squad").RemoveAll(nil)
	session.DB("workshop").C("dev").RemoveAll(nil)
	session.DB("workshop").C("bu").RemoveAll(nil)
	defer session.Close()
}*/

func InitialDB() (*mgo.Session, error){
	session, err := dial(url)
	if err != nil {
		return nil,err
	}else{
		return session,nil
	}
}
func Insert(collectionName string, obj interface{}) (error) {
	//var result interface{}

	//Database Connecting
	session,_ := InitialDB()
	collection := session.DB("workshop").C(collectionName)
	defer session.Close()

	//Insert
	err := collection.Insert(&obj)
	if err != nil {
		return err }

	//Convert
	/*byte, err := marshall(&result)
	if err != nil {	return err}
	err = json.Unmarshal(byte, &obj)
	if err != nil {	return err}*/
	return nil
}

func FindOne(collectionName string, query bson.M, obj interface{})error {
	var result interface{}
	//Database Connecting
	session,_ := InitialDB()
	collection := session.DB("workshop").C(collectionName)
	defer session.Close()

	//Find One
	err := collection.Find(query).One(&result)
	if err != nil {
		return errors.New(err.Error()) }

	//Convert
	byte, err := marshal(&result)
	if err != nil {
		return errors.New(err.Error())}
	unmarshal(byte, &obj)

	return nil
}

func FindAll(collectionName string,obj interface{}) error {
	var results []interface{}
	//Database Connecting
	session,_ := InitialDB()
	collection := session.DB("workshop").C(collectionName)
	defer session.Close()

	//Find One
	err := collection.Find(nil).All(&results)
	if results == nil {
		return errors.New("service Can't findAll Dev!!") }

	//Convert
	byte, err := marshal(&results)
	if err != nil {
		return errors.New(err.Error())}
	unmarshal(byte, obj)

	return nil
}

func Update(collectionName string, query bson.M, obj interface{}) error{
	//var result interface{}
	//Database Connecting
	session,_ := InitialDB()
	collection := session.DB("workshop").C(collectionName)
	defer session.Close()

	//Update bu
	err := collection.Update(query,&obj)
	if err != nil {
		return errors.New(err.Error()) }

	return nil

}
