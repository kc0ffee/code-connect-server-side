package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kc0ffee/server/models"
)

var (
	connection *sql.DB
)

// NewDBConnection create new connection to database
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

// ConnectionInitialize はデータベースとの接続を初期化します。
func ConnectionInitialize(db *sql.DB) {
	connection = db
}

// FetchCodeDataByID はデータベースに登録されているコードをIDを元に取得します。
func FetchCodeDataByID(id int) (*models.CodeDataResponse, error) {
	var result = models.CodeDataResponse{}
	row := connection.QueryRow("SELECT * FROM results WHERE id = ?", id)
	err := row.Scan(&result.ID, &result.Theme, &result.Code, &result.Timestamp, &result.Lang)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FetchCodeDataByTimestamp はtimestampを元にデータをデータベースから検索し返します
func FetchCodeDataByTimestamp(timestamp string) (*models.CodeDataResponse, error) {
	var result = models.CodeDataResponse{}
	row := connection.QueryRow("SELECT * FROM results WHERE timestamp = ?", timestamp)
	err := row.Scan(&result.ID, &result.Theme, &result.Lang, &result.Code, &result.Timestamp)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddCodeData はサーバーにPOSTされたデータをデータベースに保存し、保存時のtimestampを返します。
func AddCodeData(body *models.CodeData) (string, error) {
	now := time.Now()
	Timestamp := now.Format("2006-01-02 15:04:05")
	_, err := connection.Exec("INSERT INTO results (theme, lang, code, timestamp) VALUES (?, ?, ?, ?)", body.Theme, body.Lang, body.Code, Timestamp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Inserted into database")
	return Timestamp, nil
}
