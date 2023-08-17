package theme

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kc0ffee/server/models"
	"github.com/labstack/echo/v4"
)

var (
	themeList *models.ThemeList
)

func loadTheme() (*models.ThemeList, error) {
	f, err := os.Open("themes.json")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	var theme models.ThemeList
	if err = json.Unmarshal(bytes, &theme); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	return &theme, nil
}

func init() {
	data, err := loadTheme()
	if err != nil {
		panic("テーマファイルの読み込みに失敗しました。")
	}
	themeList = data
	return
}

func GetThemeList() *models.ThemeList {
	return themeList
}

func GetThemeHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, themeList)
}
