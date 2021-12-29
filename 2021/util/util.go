package util

import (
	"bufio"
	"os"
	"sort"
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

func Sum(array []int) int {
	var sum int = 0
	for _, value := range array {
		sum += value
	}
	return sum
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func Intersect(s1 string, s2 string) []rune {
	m := make(map[rune]bool)
	result := make([]rune, 0)
	for _, s := range s1 {
		m[s] = true
	}

	for _, s := range s2 {
		if _, ok := m[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

func ConvertToIntArray(s string) []int {
	result := make([]int, 0)
	for _, r := range s {
		number, _ := strconv.Atoi(string(r))
		result = append(result, number)
	}
	return result
}
