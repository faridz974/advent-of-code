package day06

import (
	"advent-of-code/2021/util"
	"fmt"
	"strconv"
	"strings"
)

type Fish struct {
	timer    int
	children []Fish
}

func createFish(timer int) Fish {
	return Fish{timer, []Fish{}}
}

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
	growFishes(80, fishes)
}

func growFishes(days int, fishes []int) {
	for i := 0; i < 80; i++ {
		for idx, f := range fishes {
			if f == 0 {
				fishes[idx] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[idx] = f - 1
			}
		}
	}

	fmt.Println(len(fishes))
}
