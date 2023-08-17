package evaluate

import (
	"os/exec"
	"strconv"
	"strings"
)

type PythonEvaluator struct{}

func (e *PythonEvaluator) CountLines(code string) int {
	return strings.Count(code, "\n") + 1
}

func (e *PythonEvaluator) CountNestedBlocks(code string) Indent {
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

func (e *PythonEvaluator) CountTokens(code string) int {
	cmd := exec.Command("python", "./languages/python/CountTokens.py", code)
	out, err := cmd.Output()
	outInt, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		panic(err)
	}
	return outInt
}

func (e *PythonEvaluator) EvaluateAST(code string) EvaluationResult {
	cmd := exec.Command("python", "./languages/python/AST.py", "evaluate", code)
	out, err := cmd.Output()
	if err != nil {
		return EvaluationResult{}
	}
	//parse out to EvaluateResult
	result := EvaluationResult{}
	out = out[:len(out)-2]
	output := strings.Split(string(out), " ")
	result.FunctionCount, err = strconv.Atoi(output[0])
	avarage, err := strconv.ParseFloat(output[1], 32)
	result.AverageNameLength = float32(avarage)
	if err != nil {
		return EvaluationResult{}
	}

	return result
}
