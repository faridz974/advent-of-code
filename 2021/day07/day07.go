package day07

import (
	"advent-of-code/2021/util"
	"fmt"
	"math"
)

func Solve() {
	numbers, err := util.ParseLineToIntArray("./day07/input.txt")
	if err != nil {
		panic(err)
	}

	min, max := util.MinMax(numbers)
	dict := make(map[int]int)
	for i := min; i <= max; i++ {
		dict[i] = 0
		for _, n := range numbers {
			diff := int(math.Abs(float64(i - n)))
			dict[i] += diff
		}
	}

	minMap, _ := util.MinMaxMap(dict)
	fmt.Println(minMap)
}

func Solve2() {
	numbers, err := util.ParseLineToIntArray("./day07/input.txt")
	if err != nil {
		panic(err)
	}

	min, max := util.MinMax(numbers)
	dict := make(map[int]int)
	for i := min; i <= max; i++ {
		dict[i] = 0
		for _, n := range numbers {
			diff := int(math.Abs(float64(i - n)))
			fuel := 0
			for i := 1; i <= diff; i++ {
				fuel += i
			}
			dict[i] += fuel
		}
	}

	minMap, _ := util.MinMaxMap(dict)
	fmt.Println(minMap)
}
