package models

type (
	Theme struct {
		ID    int    `json:"id"`    // ID はテーマのIDが格納される
		Theme string `json:"theme"` // Theme はテーマ本文が格納される
	}

	ThemeList struct {
		Themes []Theme `json:"themes"`
	}
)
