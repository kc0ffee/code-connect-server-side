package evaluate

import (
	"testing"
)

func TestPythonEvaluator_CountLines(t *testing.T) {
	code := "print(\"Hello, World!\")"
	evaluator := PythonEvaluator{}
	lines := evaluator.CountLines(code)
	if lines != 1 {
		t.Errorf("Expected 1, got %d", lines)
	}
}

func TestPythonEvaluator_CountNestedBlocks_Tab(t *testing.T) {
	code := `def hello():
	print("Hello, World!")
	print("Hello, World!")
	print("Hello, World!")`
	evaluator := PythonEvaluator{}
	indent := evaluator.CountNestedBlocks(code)
	if indent.count != 1 {
		t.Errorf("Expected 1, got %d", indent.count)
	}
	if indent.indentType != IndentTypeTab {
		t.Errorf("Expected IndentTypeTab, got %d", indent.indentType)
	}
}

func TestPythonEvaluator_CountNestedBlocks_Space2(t *testing.T) {
	code := `def hello():
    print("Hello, World!")`
	evaluator := PythonEvaluator{}
	indent := evaluator.CountNestedBlocks(code)
	if indent.count != 1 {
		t.Errorf("Expected 1, got %d", indent.count)
	}
	if indent.indentType != IndentTypeSpace {
		t.Errorf("Expected IndentTypeSpace, got %d", indent.indentType)
	}
}

func TestPythonEvaluater_CountTokens(t *testing.T) {
	code := "print(\"Hello, World!\")" // 6 tokens including newline and EOF
	evaluator := PythonEvaluator{}
	tokens := evaluator.CountTokens(code)
	if tokens != 6 {
		t.Errorf("Expected 6, got %d", tokens)
	}
}

func TestPythonEvaluator_CountFunctionCalls(t *testing.T) {
	code := "def hello():\n\tprint(\"Hello, World!\")\n\nhello()"
	evaluator := PythonEvaluator{}
	result := evaluator.EvaluateAST(code)
	if result.FunctionCount != 2 {
		t.Errorf("Expected 2, got %d", result.FunctionCount)
	}
	if result.AverageNameLength != 5.0 {
		t.Errorf("Expected 4, got %f", result.AverageNameLength)
	}
}
