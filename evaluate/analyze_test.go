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
				Lines:             5,
				Tokens:            10,
				FunctionCount:     3,
				AverageNameLength: 10.0,
			},
			theme:        1,
			expectedType: analyst,
		},
		{
			result: &EvaluationResult{
				Lines:             5,
				Tokens:            7,
				FunctionCount:     3,
				AverageNameLength: 15.0,
			},
			theme:        1,
			expectedType: gatekeeper,
		},
		{
			result: &EvaluationResult{
				Lines:             10,
				Tokens:            10,
				FunctionCount:     1,
				AverageNameLength: 5.0,
			},
			theme:        3,
			expectedType: artist,
		},
		{
			result: &EvaluationResult{
				Lines:             10,
				Tokens:            20,
				FunctionCount:     1,
				AverageNameLength: 5.0,
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
