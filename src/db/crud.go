package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"constants"
	"log"
	"jsons"
	"fmt"
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

func InsertDocuments(session *mgo.Session, collection string, transactions []jsons.Transaction) {
	c := session.DB(constants.MongoDb).C(collection)
	for _, transaction := range transactions {
		err := c.Insert(transaction)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func InsertTransactionCollection(transactionCollection []jsons.Transaction, session *mgo.Session) {
	InsertDocuments(session, constants.TransactionsCollection, transactionCollection)
}

func GetTransactionsByDate(session *mgo.Session, startDate time.Time, endDate time.Time, ticker string) []jsons.Transaction {
	return nil
}

func GetAllDocuments(session *mgo.Session, collection string) []jsons.Transaction {
	c := session.DB(constants.MongoDb).C(collection)
	var transactions []jsons.Transaction
	err := c.Find(bson.M{}).All(&transactions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return transactions
}

func DeleteAllDocuments(session *mgo.Session, collection string) {
	session.DB(constants.MongoDb).C(collection).RemoveAll(nil)
}