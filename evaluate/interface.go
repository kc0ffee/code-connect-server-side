package evaluate

type CodeEvaluator interface {
	CountLines(code string) int
	ParseToAST(code string) (interface{}, error)
	EvaluateAST(ast interface{}) *EvaluationResult
  }
  
  type EvaluationResult struct {
	TokenCount int
	FunctionCount int
   // snip...
  }
