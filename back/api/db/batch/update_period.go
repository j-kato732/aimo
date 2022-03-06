package main

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

var query = "INSERT INTO  period_models(period) VALUES(?)"

func main() {
	db, err := sql.Open("sqlite3", "../test.db")
	checkErr(err)
	defer db.Close()

	year, month, _ := time.Now().Date()
	var period = ""

	if int(month) >= 5 && int(month) <= 10 {
		period = strconv.Itoa(year) + "05"
	} else {
		period = strconv.Itoa(year) + "11"
	}

	db.Exec(query, period)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
