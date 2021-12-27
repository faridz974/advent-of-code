package day05

import (
	"advent-of-code/2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Line struct {
	X int
	Y int
}

type Entry struct {
	from Line
	to   Line
}

func Solve() {
	lines, err := util.ReadLines("./day05/input.txt")
	if err != nil {
		panic(err)
	}

	entries := make([]Entry, len(lines))

	for i := 0; i < len(lines); i++ {
		entries[i] = parseLine(lines[i])
	}

	points := buildHorizontalVerticalPoints(entries)
	fmt.Println(getOverlap(points))
}

func Solve2() {
	lines, err := util.ReadLines("./day05/input.txt")
	if err != nil {
		panic(err)
	}

	entries := make([]Entry, len(lines))

	for i := 0; i < len(lines); i++ {
		entries[i] = parseLine(lines[i])
	}

	points := buildHorizontalVerticalDiagonalPoints(entries)
	fmt.Println(getOverlap(points))
}

func parseLine(line string) Entry {
	parse := func(s string) Line {
		v := strings.Split(s, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(v[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(v[1]))
		return Line{x, y}
	}

	values := strings.Split(line, "->")
	return Entry{
		from: parse(strings.TrimSpace(values[0])),
		to:   parse(strings.TrimSpace(values[1])),
	}
}

func buildHorizontalVerticalPoints(entries []Entry) []Line {
	lines := make([]Line, 0)

	for _, entry := range entries {
		if entry.from.X == entry.to.X || entry.from.Y == entry.to.Y {
			if entry.from.X > entry.to.X && entry.from.Y == entry.to.Y {
				newLines := make([]Line, 0)
				for i := entry.from.X; i >= entry.to.X; i-- {
					newLine := Line{i, entry.from.Y}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else if entry.from.X < entry.to.X && entry.from.Y == entry.to.Y {
				newLines := make([]Line, 0)
				for i := entry.from.X; i <= entry.to.X; i++ {
					newLine := Line{i, entry.from.Y}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else if entry.from.Y > entry.to.Y && entry.from.X == entry.to.X {
				newLines := make([]Line, 0)
				for i := entry.from.Y; i >= entry.to.Y; i-- {
					newLine := Line{entry.from.X, i}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else {
				fmt.Println("Entry invalid")
			}
		}
	}

	return lines
}

func buildHorizontalVerticalDiagonalPoints(entries []Entry) []Line {
	lines := make([]Line, 0)

	for _, entry := range entries {
		if entry.from.X == entry.to.X || entry.from.Y == entry.to.Y {
			if entry.from.X > entry.to.X && entry.from.Y == entry.to.Y {
				newLines := make([]Line, 0)
				for i := entry.from.X; i >= entry.to.X; i-- {
					newLine := Line{i, entry.from.Y}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else if entry.from.X < entry.to.X && entry.from.Y == entry.to.Y {
				newLines := make([]Line, 0)
				for i := entry.from.X; i <= entry.to.X; i++ {
					newLine := Line{i, entry.from.Y}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else if entry.from.Y > entry.to.Y && entry.from.X == entry.to.X {
				newLines := make([]Line, 0)
				for i := entry.from.Y; i >= entry.to.Y; i-- {
					newLine := Line{entry.from.X, i}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else if entry.from.Y < entry.to.Y && entry.from.X == entry.to.X {
				newLines := make([]Line, 0)
				for i := entry.from.Y; i <= entry.to.Y; i++ {
					newLine := Line{entry.from.X, i}
					newLines = append(newLines, newLine)
				}
				lines = append(lines, newLines...)
			} else {
				fmt.Println("Entry invalid")
			}
		} else if math.Abs(float64(entry.from.X-entry.to.X)) == math.Abs(float64(entry.from.Y-entry.to.Y)) {
			newLines := make([]Line, 0)
			if entry.from.X >= entry.to.X {
				i := entry.from.X
				j := entry.from.Y
				for i >= entry.to.X {
					newLine := Line{i, j}
					newLines = append(newLines, newLine)
					i--
					if entry.from.Y <= entry.to.Y {
						j++
					} else {
						j--
					}
				}
			} else {
				i := entry.from.X
				j := entry.from.Y
				for i <= entry.to.X {
					newLine := Line{i, j}
					newLines = append(newLines, newLine)
					i++
					if entry.from.Y <= entry.to.Y {
						j++
					} else {
						j--
					}
				}

			}
			lines = append(lines, newLines...)
		}
	}

	return lines
}

func getOverlap(lines []Line) int {
	overlap := 0

	var dic = make(map[string]int)
	for _, pt := range lines {
		key := fmt.Sprintf("%d_%d", pt.X, pt.Y)
		if elem, ok := dic[key]; ok {
			dic[key] = elem + 1
		} else {
			dic[key] = 1
		}
	}

	for _, elem := range dic {
		if elem >= 2 {
			overlap++
		}
	}

	return overlap
}
