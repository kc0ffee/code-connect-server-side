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

func (e *GoEvaluator) CountNestedBlocks(code string) Indent {
	lines := strings.Split(code, "\n")
	nestCount := 0
	maxNest := 0
	result := Indent{count: 0}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == ' ' {
			result.indentType = IndentTypeSpace
			spaceCount := 0
			for _, c := range line {
				if c == ' ' {
					spaceCount++
				} else {
					break
				}
			}
			if spaceCount > maxNest {
				nestCount++
				maxNest = spaceCount
			}
		}
		if line[0] == '\t' {
			result.indentType = IndentTypeTab
			tabCount := 0
			for _, c := range line {
				if c == '\t' {
					tabCount++
				} else {
					break
				}
			}
			if tabCount > maxNest {
				nestCount++
				maxNest = tabCount
			}
		}
	}
	result.count = nestCount
	return result
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
