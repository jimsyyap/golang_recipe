package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"log"
)

type Transaction struct {
	CCNum      string  `bson:"ccnum"`
	Date       string  `bson:"date"`
	Amount     float64 `bson:"amount"`
	Cvv        string  `bson:"cvv"`
	Expiration string  `bson:"expiration"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Panicln(err)
	}
	defer session.Close()

	results := make([]Transaction, 0)
	if err := session.DB("test").C("transactions").Find(nil).All(&results); err != nil {
		log.Panicln(err)
	}
	for _, txn := range results {
		fmt.Println(txn.CCNum, txn.Date, txn.Amount, txn.Cvv, txn.Expiration)
	}
}
