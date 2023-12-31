package evaluate

import (
	"math"

	"github.com/kc0ffee/server/models"
	"github.com/kc0ffee/server/theme"
)

type AnalysisResult string

const (
	analyst    AnalysisResult = "analyst"
	gatekeeper AnalysisResult = "gatekeeper"
	diplomat   AnalysisResult = "diplomat"
	artist     AnalysisResult = "artist"
)

const CORRECTION_FACTOR float32 = 0.5

func Analyze(result *EvaluationResult, theme int) AnalysisResult {
	transparency := calcTransparency(theme, result.AverageNameLength, result.Lines, result.FunctionCount)
	efficiency := calcEfficiency(theme, result.Tokens, result.Lines, result.FunctionCount)

	if transparency > 0.5 && efficiency > 0.5 {
		return analyst
	} else if transparency > 0.5 && efficiency <= 0.5 {
		return gatekeeper
	} else if transparency <= 0.5 && efficiency > 0.5 {
		return diplomat
	} else {
		return artist
	}
}

func calcTransparency(theme int, averageNameLength float32, lines int, functionCount int) float32 {
	transparency := averageNameLength - float32(lines) + float32(functionCount)
	transparency = transparency + CORRECTION_FACTOR*difficulty(theme)
	return float32(math.Tanh(float64(transparency)))
}

func calcEfficiency(theme, tokens int, lines int, functionCount int) float32 {
	efficiency := float32(tokens) - float32(lines) - float32(functionCount)
	efficiency = efficiency + CORRECTION_FACTOR*difficulty(theme)
	return float32(math.Tanh(float64(efficiency)))
}

func difficulty(id int) float32 {
	// load ../themes.json

	var themeData models.ThemeList
	themeData = *theme.GetThemeList()
	length := len(themeData.Themes)

	return float32(1-id) / float32(length-1)
}
