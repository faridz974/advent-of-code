package day09

import (
	"advent-of-code/2021/util"
	"fmt"
	"sort"
)

type Point struct {
	y     int
	x     int
	value int
}

type Bassin struct {
	points []Point
	length int
}

type BassinList []Bassin

func (b BassinList) Len() int           { return len(b) }
func (b BassinList) Less(i, j int) bool { return b[i].length < b[j].length }
func (b BassinList) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func Solve() {
	lines, error := util.ReadLines("./day09/input.txt")
	if error != nil {
		panic(error)
	}

	input := make([][]int, len(lines))
	for idx, line := range lines {
		input[idx] = util.ConvertToIntArray(line)
	}

	points := findLowPoints(input)
	riskLevel := getRiskLevel(points)
	fmt.Println(riskLevel)
}

func Solve2() {
	lines, error := util.ReadLines("./day09/input.txt")
	if error != nil {
		panic(error)
	}

	input := make([][]int, len(lines))
	for idx, line := range lines {
		input[idx] = util.ConvertToIntArray(line)
	}
	points := findLowPoints(input)
	bassins := make([]Bassin, 0)
	for _, point := range points {
		bassin := getBassin(point, input, &points)
		if len(bassin) > 0 {
			bassin = append(bassin, point)
			bassins = append(bassins, Bassin{points: bassin, length: len(bassin)})
		}
	}

	largestBassins := largestBassins(bassins, 3)
	fmt.Println(largestBassins[0].length * largestBassins[1].length * largestBassins[2].length)
}

func findLowPoints(input [][]int) []Point {
	points := make([]Point, 0)
	count := 0
	rows := len(input)
	for i := 0; i < rows; i++ {
		x := input[i]
		columns := len(x)
		for j := 0; j < columns; j++ {
			current := x[j]
			if i == 0 { // first line
				bottom := input[i+1][j]
				if j == 0 { // first column
					right := x[j+1]
					if current < right && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else if j <= columns-2 { // middle columns
					right := x[j+1]
					left := x[j-1]
					if current < right && current < left && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else { // last columns
					left := x[j-1]
					if current < left && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				}
			} else if i == rows-1 { // last line
				top := input[i-1][j]
				if j == 0 { // first column
					right := x[j+1]
					if current < right && current < top {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else if j <= columns-2 { // middle columns
					left := x[j-1]
					right := x[j+1]
					if current < left && current < right && current < top {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else { // last column
					left := x[j-1]
					if current < left && current < top {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				}
			} else { // middle lines
				top := input[i-1][j]
				bottom := input[i+1][j]
				if j == 0 { // first column
					right := x[j+1]
					if current < right && current < top && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else if j <= columns-2 { // middle columns
					left := x[j-1]
					right := x[j+1]
					if current < left && current < right && current < top && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				} else { // last column
					left := x[j-1]
					if current < left && current < top && current < bottom {
						count++
						points = append(points, Point{y: i, x: j, value: current})
					}
				}
			}
		}
	}
	return points
}

func getRiskLevel(points []Point) int {
	riskLevel := 0
	for _, point := range points {
		riskLevel += point.value + 1
	}
	return riskLevel
}

func getBassin(point Point, puzzle [][]int, previousPoints *[]Point) []Point {
	bassins := make([]Point, 0)
	xLength := len(puzzle[point.y])
	yLength := len(puzzle)

	pointsAvailable := make([]Point, 0)
	if point.y != 0 { // top
		x := point.x
		y := point.y - 1
		top := Point{x: x, y: y, value: puzzle[y][x]}
		if top.value != 9 && top.value != -1 && !isPointInList(top, *previousPoints) {
			pointsAvailable = append(pointsAvailable, top)
		}

	}

	if point.y != yLength-1 { // bottom
		x := point.x
		y := point.y + 1
		bottom := Point{x: x, y: y, value: puzzle[y][x]}
		if bottom.value != 9 && bottom.value != -1 && !isPointInList(bottom, *previousPoints) {
			pointsAvailable = append(pointsAvailable, bottom)
		}
	}

	if point.x != 0 { // left
		x := point.x - 1
		y := point.y
		left := Point{x: x, y: y, value: puzzle[y][x]}
		if left.value != 9 && left.value != -1 && !isPointInList(left, *previousPoints) {
			pointsAvailable = append(pointsAvailable, left)
		}
	}

	if point.x != xLength-1 { // right
		x := point.x + 1
		y := point.y
		right := Point{x: x, y: y, value: puzzle[y][x]}
		if right.value != 9 && right.value != -1 && !isPointInList(right, *previousPoints) {
			pointsAvailable = append(pointsAvailable, right)
		}
	}

	if len(pointsAvailable) > 0 {
		bassins = append(bassins, pointsAvailable...)
		*previousPoints = append(*previousPoints, pointsAvailable...)
		for _, point := range pointsAvailable {
			newPoints := getBassin(point, puzzle, previousPoints)
			if len(newPoints) > 0 {
				bassins = append(bassins, newPoints...)
			}
		}
	}

	return bassins
}

func isPointInList(point Point, points []Point) bool {
	for _, p := range points {
		if p.x == point.x && p.y == point.y && p.value == point.value {
			return true
		}
	}
	return false
}

func largestBassins(bassins []Bassin, count int) BassinList {
	sort.Sort(BassinList(bassins))
	return bassins[len(bassins)-count:]
}
