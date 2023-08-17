package code

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/kc0ffee/server/database"
	"github.com/kc0ffee/server/models"
	"github.com/labstack/echo/v4"
)

func PostCodeHandler(ctx echo.Context) error {
	body := models.CodeData{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request")
	}
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
	res, err := database.FetchCodeDataByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusBadRequest, "Not found Id")
		}
		return ctx.String(http.StatusInternalServerError, "Error")
	}
	return ctx.JSON(http.StatusOK, res)
}
