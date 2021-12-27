package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ParseLineToIntArray(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	entries := strings.Split(lines[0], ",")
	values := make([]int, 0)
	for i := 0; i < len(entries); i++ {
		pos, _ := strconv.Atoi(entries[i])
		values = append(values, pos)
	}

	return values, nil
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MinMaxMap(m map[int]int) (int, int) {
	var max int = m[0]
	var min int = m[0]
	for _, value := range m {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
