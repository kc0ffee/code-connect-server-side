package evaluate

type CodeEvaluator interface {
	CountLines(code string) int
	CountNestedBlocks(code string) int
	CountTokens(code string) int
	ParseToAST(code string) (interface{}, error)
	EvaluateAST(ast interface{}) *EvaluationResult
}

type IndentType int

const (
	IndentTypeSpace IndentType = 0
	IndentTypeTab              = 1
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
	lines             int
	tokens            int
	indent            Indent
	functionCount     int
	averageNameLength float32
}
