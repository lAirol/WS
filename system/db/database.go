package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := "user=root password=root dbname=test sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func dbquery() {
	var err error
	// Пример выполнения запроса:
	_, err = db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)", "Apple", 72000)
	if err != nil {
		panic(err)
	}
}

func CheckUrlExists(url string) (bool, interface{}) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM your_table WHERE url like $1)", url).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
