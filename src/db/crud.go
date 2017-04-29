package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"constants"
	"log"
	"jsons"
	"fmt"
	"utils"
	"time"
)

func GetSession() *mgo.Session {
	fmt.Println("new session")
	session, err := mgo.Dial(constants.MongoUrl)
	if err != nil {
		log.Fatal("DB: ", err)
		panic(err)
	}
	return session
}

func InsertOwnershipDocuments(session *mgo.Session, collection string, documents []jsons.OwnershipDocument) {
	c := session.DB(constants.MongoDb).C(collection)
	for _, document := range documents {
		err := c.Insert(document)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func InsertLastUpdated(session *mgo.Session, collection string, lastUpdated jsons.LastUpdated) {
	c := session.DB(constants.MongoDb).C(collection)
	c.RemoveAll(nil)
	err := c.Insert(lastUpdated)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetLastUpdated(session *mgo.Session) time.Time {
	c := session.DB(constants.MongoDb).C(constants.LastUpdated)
	var lastUpdated jsons.LastUpdated
	err := c.Find(bson.M{}).One(&lastUpdated)
	var returnValue time.Time
	if err != nil {
		fmt.Println("No record in database")
		returnValue = time.Now().Add(-25*time.Hour)
	} else {
		fmt.Println("reached else")
		returnValue = lastUpdated.LastUpdated
	}
	return returnValue
}

func GetAllLastUpdated(session *mgo.Session) jsons.LastUpdated {
	c := session.DB(constants.MongoDb).C(constants.LastUpdated)
	var lastUpdated jsons.LastUpdated
	err := c.Find(bson.M{}).All(&lastUpdated)
	utils.HandleError(err)
	return lastUpdated
}

func DeleteAllDocuments(session *mgo.Session, collection string) {
	session.DB(constants.MongoDb).C(collection).RemoveAll(nil)
}