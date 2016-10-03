package dao

import "gopkg.in/mgo.v2"

func IntialDB() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	return session
}
