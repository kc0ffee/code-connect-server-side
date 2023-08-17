package evaluate

func Evaluate(theme int, code string, eval CodeEvaluator) *EvaluationResult {
	result := eval.EvaluateAST(code)
	result.lines = eval.CountLines(code)
	result.tokens = eval.CountTokens(code)
	result.indent = eval.CountNestedBlocks(code)

	return result
}
