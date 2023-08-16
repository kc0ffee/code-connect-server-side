package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
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

type ReturnData struct {
	ID        int64  `json:"id"`
	Theme     string `json:"theme"`
	Code      string `json:"code"`
	Timestamp string `json:"timestamp"`
}

type PostData struct {
	Theme string `json:"theme"`
	Code  string `json:"code"`
}

func GetResultById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}
	var result = ReturnData{}
	row := db.QueryRow("SELECT * FROM results WHERE id = ?", id)
	err = row.Scan(&result.ID, &result.Theme, &result.Code, &result.Timestamp)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting result")
	}
	json, err := json.Marshal(result)
	return c.String(http.StatusOK, string(json))
}

func CreateResult(c echo.Context, db *sql.DB) error {
	body := new(PostData)
	if err := (&body); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}
	result := ReturnData{Theme: body.Theme, Code: body.Code}
	now := time.Now()
	result.Timestamp = now.Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO results (theme, code, timestamp) VALUES (?, ?, ?)", result.Theme, result.Code, result.Timestamp)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error inserting result")
	}
	return c.String(http.StatusOK, "Result posted")
}
