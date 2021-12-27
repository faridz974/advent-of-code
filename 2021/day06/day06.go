package day06

import (
	"advent-of-code/2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	lines, err := util.ReadLines("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	entries := strings.Split(lines[0], ",")
	fishes := make([]int, 0)
	for i := 0; i < len(entries); i++ {
		timer, _ := strconv.Atoi(entries[i])
		fishes = append(fishes, timer)
	}
	growFishes(256, fishes)
}

func Solve2() {
	lines, err := util.ReadLines("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	entries := strings.Split(lines[0], ",")
	fishes := make([]int, 0)
	for i := 0; i < len(entries); i++ {
		timer, _ := strconv.Atoi(entries[i])
		fishes = append(fishes, timer)
	}

	growBigFishes(256, fishes)
}

func growFishes(days int, fishes []int) {
	for i := 0; i < days; i++ {
		for idx, f := range fishes {
			if f == 0 {
				fishes[idx] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[idx] = f - 1
			}
		}

		fmt.Println(len(fishes))
	}

	fmt.Println(len(fishes))
}

func growBigFishes(days int, fishes []int) {
	m := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, f := range fishes {
		m[f] = m[f] + 1
	}

	for i := 0; i < days; i++ {
		fishPregnants := m[0]
		for x := 1; x <= len(m); x++ {
			m[x-1] = m[x]
		}
		m[6] += fishPregnants
		m[8] = fishPregnants
	}

	sum := 0
	for x := 0; x < len(m); x++ {
		sum += m[x]
	}

	fmt.Println(sum)
}
