package handler

import "github.com/labstack/echo/v4"

type (
	Theme struct {
		ID    int    `json:"id"`    // ID はテーマのIDが格納される
		Theme string `json:"theme"` // Theme はテーマ本文が格納される
	}

	GetThemeAPIResponse struct {
		Themes []Theme `json:"themes"`
	}
)

func GetThemeHandler(ctx echo.Context) error {
	return nil
}
