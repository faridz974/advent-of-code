package day08

import (
	"advent-of-code/2021/util"
	"fmt"
	"strconv"
	"strings"
)

type Digit struct {
	number   int
	original string
	value    string
}

type UnknowDigit struct {
	number   []int
	original string
	value    string
}

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

func Solve2() {
	lines, error := util.ReadLines("./day08/input.txt")
	if error != nil {
		panic(error)
	}

	outputs := make([]string, len(lines))
	inputs := make([]string, len(lines))
	for idx, line := range lines {
		entries := strings.Split(line, "|")
		inputs[idx] = strings.TrimSpace(entries[0])
		outputs[idx] = strings.TrimSpace(entries[1])
	}

	fmt.Println(addDigits(inputs, outputs))
}

func addDigits(inputs []string, entries []string) int {
	numbers := make([]int, 0)

	for inputIdx, input := range inputs {
		mapNumberValue := make(map[int]Digit)
		words := strings.Split(input, " ")
		unknown5Digits := make(map[string]UnknowDigit)
		unknown6Digits := make(map[string]UnknowDigit)

		for _, word := range words {
			switch len(word) {
			case 2:
				mapNumberValue[1] = Digit{number: 1, original: word, value: util.Sort(word)}
			case 3:
				mapNumberValue[7] = Digit{number: 7, original: word, value: util.Sort(word)}
			case 4:
				mapNumberValue[4] = Digit{number: 4, original: word, value: util.Sort(word)}
			case 5:
				unknown5Digits[word] = UnknowDigit{number: []int{2, 3, 5}, original: word, value: util.Sort(word)}
			case 6:
				unknown6Digits[word] = UnknowDigit{number: []int{0, 6, 9}, original: word, value: util.Sort(word)}
			case 7:
				mapNumberValue[8] = Digit{number: 8, original: word, value: util.Sort(word)}
			}
		}

		three := findOccurrences(mapNumberValue[7], unknown5Digits, 3)
		delete(unknown5Digits, three.original)
		mapNumberValue[3] = Digit{number: 3, original: three.original, value: util.Sort(three.original)}

		five := findOccurrences(mapNumberValue[4], unknown5Digits, 3)
		delete(unknown5Digits, five.original)
		mapNumberValue[5] = Digit{number: 5, original: five.original, value: util.Sort(five.original)}

		var two UnknowDigit
		for _, digit := range unknown5Digits {
			two = digit
		}
		delete(unknown5Digits, two.original)
		mapNumberValue[2] = Digit{number: 2, original: two.original, value: util.Sort(two.original)}

		nine := findOccurrences(mapNumberValue[4], unknown6Digits, 4)
		delete(unknown6Digits, nine.original)
		mapNumberValue[9] = Digit{number: 9, original: nine.original, value: util.Sort(nine.original)}

		zero := findOccurrences(mapNumberValue[7], unknown6Digits, 3)
		delete(unknown6Digits, zero.original)
		mapNumberValue[0] = Digit{number: 0, original: zero.original, value: util.Sort(zero.original)}

		var six UnknowDigit
		for _, digit := range unknown6Digits {
			six = digit
		}
		delete(unknown6Digits, six.original)
		mapNumberValue[6] = Digit{number: 6, original: six.original, value: util.Sort(six.original)}

		entry := strings.Split(entries[inputIdx], " ")
		final := ""
		for _, word := range entry {
			sortedWord := util.Sort(word)
			for _, assoc := range mapNumberValue {
				if assoc.value == sortedWord {
					final += strconv.Itoa(assoc.number)
				}
			}
		}

		newNumber, _ := strconv.Atoi(final)
		numbers = append(numbers, newNumber)
	}

	return util.Sum(numbers)
}

func findOccurrences(digitToMatch Digit, digits map[string]UnknowDigit, expectedLength int) UnknowDigit {
	for _, digit := range digits {
		inter := util.Intersect(digit.value, digitToMatch.value)
		if len(inter) == expectedLength {
			return digit
		}
	}
	return UnknowDigit{}
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
