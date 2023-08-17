// api GET /api/result
// request body:
//
//	{
//	   "id": number
//	}
//
// 1. fetch code data by request id
// 2. using code content, theme, and lang(changing Evaluator) evaluate
// 3. analyze with evaluation result
// 4. result analyze result
package result

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/kc0ffee/server/database"
	"github.com/kc0ffee/server/evaluate"
	"github.com/kc0ffee/server/models"
	"github.com/labstack/echo/v4"
)

var (
	analyzeResult *models.AnalyzeResponse
)

func GetResult(ctx echo.Context) (*models.AnalyzeResponse, error) {
	id, err := strconv.Atoi(ctx.QueryParam("id"))
	if err != nil {
		return nil, ctx.String(http.StatusBadRequest, "Invalid id")
	}
	data, err := database.FetchCodeDataByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ctx.String(http.StatusBadRequest, "Not found Id")
		}
		return nil, ctx.String(http.StatusInternalServerError, "Error")
	}

	evaluator, err := evaluate.NewEvaluator(data.Lang)
	if err != nil {
		return nil, ctx.String(http.StatusInternalServerError, "Error")
	}

	evaluateResult := evaluate.Evaluate(data.Theme, data.Code, evaluator)
	analyzeResult := evaluate.Analyze(&evaluateResult, data.Theme)

	result := &models.AnalyzeResponse{
		Type:        string(analyzeResult),
		Description: "This is result of analyze",
		Line:        evaluateResult.Lines,
		TokenLen:    evaluateResult.AverageNameLength,
		TokenCount:  evaluateResult.Tokens,
		FuncLen:     evaluateResult.FunctionCount,
	}

	return result, nil
}

func ResultHandler(ctx echo.Context) error {
	result, err := GetResult(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
