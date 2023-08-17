package evaluate

import "errors"

type CodeEvaluator interface {
	CountLines(code string) int
	CountNestedBlocks(code string) Indent
	CountTokens(code string) int
	EvaluateAST(code string) EvaluationResult
}

func NewEvaluator(lang string) (CodeEvaluator, error) {
	switch lang {
	case "go":
		return &GoEvaluator{}, nil
	case "python":
		return &PythonEvaluator{}, nil
	default:
		return nil, errors.New("Invalid language")
	}
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
	Lines             int
	Tokens            int
	Indent            Indent
	FunctionCount     int
	AverageNameLength float32
}
