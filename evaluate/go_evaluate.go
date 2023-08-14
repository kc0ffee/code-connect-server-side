package evaluate

import (
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"strings"
)

type GoEvaluator struct{}

const space = 0x20

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
		if line[0] == space {
			result.indentType = IndentTypeSpace
			spaceCount := 0
			for _, c := range line {
				if c == space {
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

func (e *GoEvaluator) CountTokens(code string) int {
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(code))
	var s scanner.Scanner
	s.Init(file, []byte(code), nil, 0)

	count := 0
	for {
		_, tok, _ := s.Scan()
		if tok == token.EOF {
			break
		}
		count++
	}
	return count
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
	tokenLen := []int{}
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
			tokenLen = append(tokenLen, len(n.(*ast.Ident).Name))
		}
		return true
	})
	// calculate average name length
	sum := 0
	for _, l := range tokenLen {
		sum += l
	}
	result.AverageNameLength = float32(sum) / float32(len(tokenLen))

	return result
}
