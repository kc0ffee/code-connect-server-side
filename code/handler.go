package code

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

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

	timestamp, err := database.AddCodeData(&body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failure to update database")
	}

	res, err := database.FetchCodeDataByTimestamp(timestamp)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failure to update database")
	}
	return ctx.JSON(http.StatusOK, res)
}

func ResultHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.QueryParam("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid id")
	}
	data, err := database.FetchCodeDataByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusBadRequest, "Not found Id")
		}
		return ctx.String(http.StatusInternalServerError, "Error")
	}

	res := models.CodeDataResponse{
		ID:        data.ID,
		Theme:     data.Theme,
		Lang:      data.Lang,
		Code:      base64.StdEncoding.EncodeToString([]byte(data.Code)),
		Timestamp: data.Timestamp,
	}

	return ctx.JSON(http.StatusOK, res)
}
