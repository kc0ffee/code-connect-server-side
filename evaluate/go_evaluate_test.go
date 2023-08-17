package evaluate

import (
	"testing"
)

func TestGoEvaluator_CountLines(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\n\nfunc main() {\n\tprintln(\"Hello, world!\")\n}"
	lines := evaluator.CountLines(code)
	if lines != 5 {
		t.Errorf("Expected 5 lines, but got %d", lines)
	}
}

func TestGoEvaluator_CountNestedBlocks_CaseTab(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\n\nfunc main() {\n\tprintln(\"Hello, world!\")\n}"
	result := evaluator.CountNestedBlocks(code)
	nestedBlocks := result.count
	indentType := result.indentType
	if nestedBlocks != 1 {
		t.Errorf("Expected 1 nested blocks, but got %d", nestedBlocks)
	}
	if indentType != IndentTypeTab {
		t.Errorf("Expected indent type tab, but got %s", indentType)
	}
}

func TestGoEvaluator_CountNestedBlocks_CaseSpace2(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\n\nfunc main() {\n  println(\"Hello, world!\")\n}"
	result := evaluator.CountNestedBlocks(code)
	nestedBlocks := result.count
	indentType := result.indentType
	if nestedBlocks != 1 {
		t.Errorf("Expected 1 nested blocks, but got %d", nestedBlocks)
	}
	if indentType != IndentTypeSpace {
		t.Errorf("Expected indent type space, but got %s", indentType)
	}
}

func TestGoEvaluator_CountTOkens(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\n\nfunc main() {\n\tprintln(\"Hello, world!\")\n}"
	tokens := evaluator.CountTokens(code)
	if tokens != 15 {
		t.Errorf("Expected 15 tokens, but got %d", tokens)
	}
}

func TestGoEvaluator_CountNestedBlocks_CaseSpace4(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\n\nfunc main() {\n    println(\"Hello, world!\")\n}"
	result := evaluator.CountNestedBlocks(code)
	nestedBlocks := result.count
	indentType := result.indentType
	if nestedBlocks != 1 {
		t.Errorf("Expected 1 nested blocks, but got %d", nestedBlocks)
	}
	if indentType != IndentTypeSpace {
		t.Errorf("Expected indent type space, but got %s", indentType)
	}
}

func TestGoEvaluator_ParseToAST(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\nfunc main() {}"
	_, err := evaluator.ParseToAST(code)
	if err != nil {
		t.Errorf("Failed to parse code: %v", err)
	}
}

func TestGoEvaluator_EvaluateAST_FuncCount(t *testing.T) {
	evaluator := &GoEvaluator{}
	code := "package main\nfunc main() {}\nfunc test() {}\nfunc test2() {}"
	_, err := evaluator.ParseToAST(code)
	if err != nil {
		t.Fatalf("Failed to parse code: %v", err)
	}

	result := evaluator.EvaluateAST(code)
	if result.FunctionCount != 3 {
		t.Errorf("Expected 2 functions, but got %d", result.FunctionCount)
	}
	if result.AverageNameLength != 4.25 {
		t.Errorf("Expected 4.5 average name length, but got %f", result.AverageNameLength)
	}
}
