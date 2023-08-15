package server

import (
	"fmt"

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

	return e
}

func StartServer(e *echo.Echo, port int) {
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
