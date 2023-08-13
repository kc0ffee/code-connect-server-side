package evaluate

type GoEvaluator struct{}

func (e *GoEvaluator) CountLines(code string) int {
	// count lines
	return 0
}

func ParseToAST(code string) (interface{}, error) {
	// use go/ast
	return nil, nil
}

func (e *GoEvaluator) EvaluateAST(ast interface{}) *EvaluationResult {
	//evaluate
	return nil
}
