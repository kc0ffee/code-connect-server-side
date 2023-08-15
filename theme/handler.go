package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

type (
	Theme struct {
		ID    int    `json:"id"`    // ID はテーマのIDが格納される
		Theme string `json:"theme"` // Theme はテーマ本文が格納される
	}

	GetThemeAPIResponse struct {
		Themes []Theme `json:"themes"`
	}
)

var (
	Response *GetThemeAPIResponse
)

func loadTheme() (*GetThemeAPIResponse, error) {
	f, err := os.Open("themes.json")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	var theme GetThemeAPIResponse
	if err = json.Unmarshal(bytes, &theme); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	return &theme, nil
}

func init() {
	data, err := loadTheme()
	if err != nil {
		return
	}
	Response = data
	return
}

func GetThemeHandler(ctx echo.Context) error {
	return nil
}
