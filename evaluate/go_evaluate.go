package evaluate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type GoEvaluator struct{}

func (e *GoEvaluator) CountLines(code string) int {
	return strings.Count(code, "\n") + 1
}

func (e *GoEvaluator) ParseToAST(code string) (interface{}, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (e *GoEvaluator) EvaluateAST(targetAst interface{}) *EvaluationResult {
	result := &EvaluationResult{}
	node, ok := targetAst.(*ast.File)
	if !ok {
		return result
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			result.FunctionCount++
		case *ast.Ident:
			result.TokenCount++
		}
		return true
	})

	return result
}
