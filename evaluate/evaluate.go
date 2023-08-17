package evaluate

func Evaluate(theme int, code string, eval CodeEvaluator) EvaluationResult {
	result := eval.EvaluateAST(code)
	result.Lines = eval.CountLines(code)
	result.Tokens = eval.CountTokens(code)
	result.Indent = eval.CountNestedBlocks(code)

	return result
}
