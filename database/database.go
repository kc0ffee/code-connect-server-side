package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection(ADDRESS string, USER string, DB_NAME string) *sql.DB {
	db, err := sql.Open("mysql", "$USER@tcp($ADDRESS)/$DB_NAME?charset=utf8mb4&parseTime=True&loc=Local")

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func Insert(db *sql.DB, them int, code string) {
	_, err := db.Exec("INSERT INTO Items (theme, code) VALUES (?, ?)", them, code)
	if err != nil {
		fmt.Println(err)
		return
	}
}
