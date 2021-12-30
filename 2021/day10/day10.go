package day10

import (
	"advent-of-code/2021/util"
	"fmt"
	"sort"
	"strings"
)

type ParsedLine struct {
	isError      bool
	isIncomplete bool
	score        int
	remaining    string
}

func Solve() {
	lines, error := util.ReadLines("./day10/sample.txt")
	if error != nil {
		panic(error)
	}

	score := make([]int, len(lines))
	for idx, line := range lines {
		parsedLine := parse(line)
		score[idx] = parsedLine.score
	}
	fmt.Println(util.Sum(score))
}

func Solve2() {
	lines, error := util.ReadLines("./day10/input.txt")
	if error != nil {
		panic(error)
	}

	score := make([]int, 0)
	for _, line := range lines {
		parsedLine := parse(line)
		if parsedLine.isIncomplete {
			missingPart := autcomplete(parsedLine)
			lineScore := calculateAutocompleteScore(missingPart)
			score = append(score, lineScore)
		}
	}

	sort.Ints(score)
	middleScore := score[len(score)/2]
	fmt.Println(middleScore)
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
	if isIncomplete(line) {
		return ParsedLine{isError: false, isIncomplete: true, remaining: line, score: 0}
	}

	return findIllegalChar(line)
}

func isIncomplete(line string) bool {
	return !strings.Contains(line, "}") && !strings.Contains(line, "]") && !strings.Contains(line, ")") && !strings.Contains(line, ">")
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
	errorChar := string(line[min])

	switch errorChar {
	case ")":
		return ParsedLine{isError: true, isIncomplete: false, remaining: line, score: 3}
	case "]":
		return ParsedLine{isError: true, isIncomplete: false, remaining: line, score: 57}
	case "}":
		return ParsedLine{isError: true, isIncomplete: false, remaining: line, score: 1197}
	case ">":
		return ParsedLine{isError: true, isIncomplete: false, remaining: line, score: 25137}
	}

	return ParsedLine{isError: false, isIncomplete: false, remaining: line, score: 0}
}

func autcomplete(parsedLine ParsedLine) []rune {
	missing := make([]rune, 0)
	for i := len(parsedLine.remaining) - 1; i >= 0; i-- {
		currentChar := rune(parsedLine.remaining[i])
		switch currentChar {
		case '(':
			missing = append(missing, ')')
		case '[':
			missing = append(missing, ']')
		case '{':
			missing = append(missing, '}')
		case '<':
			missing = append(missing, '>')
		}
	}

	return missing
}

func calculateAutocompleteScore(str []rune) int {
	total := 0

	for _, char := range str {
		switch char {
		case ')':
			total = total*5 + 1
		case ']':
			total = total*5 + 2
		case '}':
			total = total*5 + 3
		case '>':
			total = total*5 + 4
		}
	}

	return total
}
