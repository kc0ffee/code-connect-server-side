/// タイプ診断: 透明性､効率性の2軸評価
/// 透明性の評価
///   コード行数が短い
///   命名平均長が長い
///   関数の個数が多い
///
/// 効率性の評価
///   トークン数が少ない
///   ネストが浅い
///   関数の個数が少ない

/*
type EvaluationResult struct {
	lines             int
	tokens            int
	indent            Indent
	functionCount     int
	averageNameLength float32
}
*/

package evaluate

import (
	"encoding/json"
	"io"
	"math"
	"os"
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
	//theme is difficulty, lower is easier, higher is harder

	//透明性の評価
	transparency := calcTransparency(theme, result.averageNameLength, result.lines, result.functionCount)
	efficiency := calcEfficiency(theme, result.tokens, result.lines, result.functionCount)

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

type Themes struct {
	Themes []struct {
		ID    int    `json:"id"`
		Theme string `json:"theme"`
	} `json:"themes"`
}

func difficulty(id int) float32 {
	// load ../themes.json
	f, err := os.Open("../themes.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var theme Themes
	if err := json.Unmarshal(bytes, &theme); err != nil {
		panic(err)
	}
	length := len(theme.Themes)

	return float32(1-id) / float32(length-1)
}
