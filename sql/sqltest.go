package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insert(db *sql.DB, nums int) bool {
	stmtIns, err := db.Prepare("insert into go_test(name) values(?)")
	if err != nil {
		panic(err.Error())
	}

	defer stmtIns.Close()
	for i := 0; i <= nums; i++ {
		_, err := stmtIns.Exec(i * i)
		if err != nil {
			panic(err.Error())
		}
	}
	return true
}

func getNumber(db *sql.DB, nums int) int {
	return 1
}
func main() {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	stmtOut, err := db.Prepare("select name from go_test where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var squareNum int
	err = stmtOut.QueryRow(12).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("the square number of 13 is: %d", squareNum)

	err = stmtOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}

	insert(db, 1)
	fmt.Printf("the square number of 1 is : %d", squareNum)
	for i := 1; i < 25; i++ {
		err = stmtOut.QueryRow(i).Scan(&squareNum)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("the square number of 1 is : %d", squareNum)

	}
	// rows, err := db.Query("select * from go_test")
	// db query select
	// rows, err := db.Query("select * from go_test")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// columns, err := rows.Columns()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(values))

	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rows.Next() {
	// 	err := rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	var value string
	// 	for i, col := range values {
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}

	// 		fmt.Printf(columns[i], ":", value, "\n")
	// 	}

	// 	fmt.Printf("-------------------------")
	// }

}
