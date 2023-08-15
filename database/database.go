package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection(ADDRESS string, USER string, DB_NAME string) *sql.DB {
	str := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, ADDRESS, DB_NAME)
	db, err := sql.Open("mysql", str)
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Connected to database")
	return db
}

func Insert(db *sql.DB, them int, code string) {
	fmt.Println("Inserting into database")
	_, err := db.Exec("INSERT INTO results (theme, code) VALUES (?, ?)", them, code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted into database")
}
