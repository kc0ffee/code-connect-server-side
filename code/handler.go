package code

import (
	"encoding/base64"
	"net/http"

	"github.com/kc0ffee/server/database"
	"github.com/kc0ffee/server/models"
	"github.com/labstack/echo/v4"
)

func PostCodeHandler(ctx echo.Context) error {
	var body models.CodeData
	if err := ctx.Bind(&body); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request")
	}

	var registerData models.CodeData
	decoded, err := base64.StdEncoding.DecodeString(body.Code)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failure decode")
	}

	registerData.Code = string(decoded)
	registerData.Lang = body.Lang
	registerData.Theme = body.Theme

	timestamp, err := database.AddCodeData(&registerData)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failure to update database")
	}

	res, err := database.FetchCodeDataByTimestamp(timestamp)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failure to update database")
	}
	return ctx.JSON(http.StatusOK, res)
}
