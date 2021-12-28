package day08

import (
	"advent-of-code/2021/util"
	"fmt"
	"strings"
)

func Solve() {
	lines, error := util.ReadLines("./day08/input.txt")
	if error != nil {
		panic(error)
	}

	outputs := make([]string, len(lines))
	for idx, line := range lines {
		outputs[idx] = strings.TrimSpace(strings.Split(line, "|")[1])
	}

	fmt.Println(countDigits(outputs))
}

func countDigits(entries []string) int {
	count := 0

	for _, entry := range entries {
		words := strings.Split(entry, " ")
		for _, word := range words {
			switch len(word) {
			case 2:
				count++
			case 3:
				count++
			case 4:
				count++
			case 7:
				count++
			}
		}
	}

	return count
}
