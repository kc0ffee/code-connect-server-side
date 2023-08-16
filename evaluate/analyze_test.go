package evaluate

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		result       *EvaluationResult
		theme        int
		expectedType AnalysisResult
	}{
		{
			result: &EvaluationResult{
				lines:             5,
				tokens:            10,
				functionCount:     3,
				averageNameLength: 10.0,
			},
			theme:        1,
			expectedType: analyst,
		},
		{
			result: &EvaluationResult{
				lines:             5,
				tokens:            7,
				functionCount:     3,
				averageNameLength: 15.0,
			},
			theme:        1,
			expectedType: gatekeeper,
		},
		{
			result: &EvaluationResult{
				lines:             10,
				tokens:            10,
				functionCount:     1,
				averageNameLength: 5.0,
			},
			theme:        3,
			expectedType: artist,
		},
		{
			result: &EvaluationResult{
				lines:             10,
				tokens:            20,
				functionCount:     1,
				averageNameLength: 5.0,
			},
			theme:        1,
			expectedType: diplomat,
		},
	}

	for _, test := range tests {
		actualType := Analyze(test.result, test.theme)
		if actualType != test.expectedType {
			t.Errorf("Expected %v, but got %v", test.expectedType, actualType)
		}
	}
}

func TestCalcTransparency(t *testing.T) {
	// 透明性の計算に対するテスト

}

func TestCalcEfficiency(t *testing.T) {
	// 効率性の計算に対するテスト
}

func TestDifficulty(t *testing.T) {
	// 難易度の計算に対するテスト
}
