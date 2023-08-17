package server

import (
	"fmt"

	"github.com/kc0ffee/server/code"
	"github.com/kc0ffee/server/theme"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewAPIServer is a function to create new instance of server
func NewAPIServer() *echo.Echo {
	e := echo.New()

	// Set middleware
	// TODO : ロガーをEchoデフォルトのロガーではなくSLogとかに変えてもいいかも
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: APIのハンドラーの追加
	e.GET("/api/theme", theme.GetThemeHandler)
	e.POST("/api/code", code.PostCodeHandler)

	// e.GET("/api/result", func(c echo.Context) error {
	// 	return database.GetResultById(c, db)
	// })

	return e
}

func StartServer(e *echo.Echo, port int) {
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
