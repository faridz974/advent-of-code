package day10

import (
	"advent-of-code/2021/util"
	"fmt"
	"strings"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LBRACK  = "["
	RBRACK  = "]"
	LESS    = "<"
	GREATER = ">"
)

type ParsedLine struct {
	isError      bool
	isIncomplete bool
	score        int
}

func Solve() {
	lines, error := util.ReadLines("./day10/sample.txt")
	if error != nil {
		panic(error)
	}

	score := make([]int, len(lines))
	for idx, line := range lines {
		fmt.Println(line)
		parsedLine := parse(line)
		score[idx] = parsedLine.score
	}
	fmt.Println(util.Sum(score))
}

func Solve2() {
	lines, error := util.ReadLines("./day10/sample.txt")
	if error != nil {
		panic(error)
	}

	score := make([]int, len(lines))
	parsedLines := make([]ParsedLine, len(lines))
	for idx, line := range lines {
		fmt.Println(line)
		parsedLine := parse(line)
		parsedLines[idx] = parsedLine
		score[idx] = parsedLines[idx].score
		fmt.Println(util.Sum(score))
	}
}

func parse(line string) ParsedLine {
	for !isIncomplete(line) {
		before := line
		line = strings.ReplaceAll(line, "()", "")
		line = strings.ReplaceAll(line, "{}", "")
		line = strings.ReplaceAll(line, "[]", "")
		line = strings.ReplaceAll(line, "<>", "")
		if before == line {
			break
		}
	}
	fmt.Println(line)
	if isIncomplete(line) {
		return ParsedLine{isError: false, isIncomplete: true, score: 0}
	}

	return findIllegalChar(line)
}

func isIncomplete(line string) bool {
	return !strings.Contains(line, RBRACE) && !strings.Contains(line, RBRACK) && !strings.Contains(line, RPAREN) && !strings.Contains(line, GREATER)
}

func findIllegalChar(line string) ParsedLine {
	idxRparent := strings.Index(line, ")")
	idxRbrack := strings.Index(line, "]")
	idxRbrace := strings.Index(line, "}")
	idxGreat := strings.Index(line, ">")

	allIndexes := make([]int, 0)
	if idxRparent != -1 {
		allIndexes = append(allIndexes, idxRparent)
	}

	if idxRbrack != -1 {
		allIndexes = append(allIndexes, idxRbrack)
	}

	if idxRbrace != -1 {
		allIndexes = append(allIndexes, idxRbrace)
	}

	if idxGreat != -1 {
		allIndexes = append(allIndexes, idxGreat)
	}

	min, _ := util.MinMax(allIndexes)

	fmt.Println(min)

	errorChar := string(line[min])
	fmt.Println(errorChar)

	if errorChar == ")" {
		return ParsedLine{isError: true, isIncomplete: false, score: 3}
	}
	if errorChar == "]" {
		return ParsedLine{isError: true, isIncomplete: false, score: 57}
	}
	if errorChar == "}" {
		return ParsedLine{isError: true, isIncomplete: false, score: 1197}
	}
	if errorChar == ">" {
		return ParsedLine{isError: true, isIncomplete: false, score: 25137}
	}

	return ParsedLine{isError: false, isIncomplete: false, score: 0}
}
