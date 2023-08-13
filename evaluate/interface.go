package evaluate

type CodeEvaluator interface {
	CountLines(code string) int
	CountNestedBlocks(code string) int
	ParseToAST(code string) (interface{}, error)
	EvaluateAST(ast interface{}) *EvaluationResult
}

type IndentType int

const (
	IndentTypeSpace IndentType = iota
	IndentTypeTab
)

type Indent struct {
	count      int
	indentType IndentType
}

func (i IndentType) String() string {
	switch i {
	case IndentTypeSpace:
		return "space"
	case IndentTypeTab:
		return "tab"
	default:
		return "unknown"
	}
}

type EvaluationResult struct {
	TokenCount            int
	FunctionCount         int
	Indent                IndentType
	AverageFunctionLength float32
	// snip...
}
