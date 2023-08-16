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
	ID        int    `json:"id"`
	Theme     int    `json:"themeId"`
	Lang      string `json:"language"`
	Code      string `json:"code"`
	Timestamp string `json:"timeStamp"`
}

type PostData struct {
	Theme int    `json:"themeId"`
	Lang  string `json:"language"`
	Code  string `json:"code"`
}

func GetResultById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}
	var result = ReturnData{}
	row := db.QueryRow("SELECT * FROM results WHERE id = ?", id)
	err = row.Scan(&result.ID, &result.Theme, &result.Lang, &result.Code, &result.Timestamp)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting result")
	}
	json, err := json.Marshal(result)
	return c.String(http.StatusOK, string(json))
}

func CreateResult(c echo.Context, db *sql.DB) error {
	body := PostData{}
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}
	now := time.Now()
	Timestamp := now.Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO results (theme, lang, code, timestamp) VALUES (?, ?, ?, ?)", body.Theme, body.Lang, body.Code, Timestamp)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error inserting into database")
	}
	fmt.Println("Inserted into database")
	return c.String(http.StatusOK, "Success")
}
