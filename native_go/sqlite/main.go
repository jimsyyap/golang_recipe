package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"strconv"
)

type Record struct {
	Id   string
	Name string
}

var (
	dbHandle *sql.DB
	curNames = [5]string{"USD", "EUR", "GBP", "JPY", "CNY"}
	dbname   = "local.db"
)

func main() {
	log.Println("init db - ", dbname)
	dbHandle = InitDB(dbname)

	log.Println("create table", dbname)
	InitTable(dbHandle)

	log.Println("inserting into table", dbname)
	records := []Record{}

	for i := 0; i < 10; i++ {
		records = append(records, Record{Id: strconv.Itoa(i), Name: curNames[rand.Intn(5)]})
	}
}
