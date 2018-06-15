package main

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"
)

func insertValue(db *sql.DB, num int) {
	preString := "insert into go_test(name) values"
	insertValueList := []interface{}{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < num; i++ {
		preString += "(?),"
		value := rand.Int()
		insertValueList = append(insertValueList, i*value)
		// insertValueList = append(insertValueList, i)
	}
	preString = preString[:len(preString)-1]
	// fmt.Println(db)
	// fmt.Println(preString)
	// fmt.Println(insertValueList)
	smtIns, err := db.Prepare(preString)
	if err != nil {
		panic(err.Error())
	}

	rand.Int()

	// fmt.Println(db)
	// return//
	_, err2 := smtIns.Exec(insertValueList...)
	if err2 != nil {
		panic(err.Error())
	}
	// for i := 0; i <= num; i++ {
	// 	_, err := smtIns.Exec(i)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }
}

func TestInsertValue(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(db)

	for i := 0; i < 1000; i++ {
		insertValue(db, 10000)
	}

	defer db.Close()
}
