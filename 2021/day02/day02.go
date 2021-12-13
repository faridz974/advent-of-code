package day02

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Solve() {
	content, err := ioutil.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	strContent := string(content)
	array := strings.Split(strContent, "\n")

	var horizontal int64 = 0
	var depth int64 = 0
	for i := 0; i < len(array); i++ {
		instruction := strings.Split(array[i], " ")
		command := instruction[0]
		unitStr := instruction[1]

		unit, errUnit := strconv.ParseInt(unitStr, 10, 64)
		if errUnit != nil {
			log.Fatalf("unable to parse unit: %v", errUnit)
			break
		}

		switch command {
		case "forward":
			horizontal = unit + horizontal
		case "down":
			depth = depth + unit
		case "up":
			depth = depth - unit
		}
	}

	fmt.Println(horizontal * depth)
}

func Solve2() {
	content, err := ioutil.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	strContent := string(content)
	array := strings.Split(strContent, "\n")

	var horizontal int64 = 0
	var aim int64 = 0
	var depth int64 = 0
	for i := 0; i < len(array); i++ {
		instruction := strings.Split(array[i], " ")
		command := instruction[0]
		unitStr := instruction[1]

		unit, errUnit := strconv.ParseInt(unitStr, 10, 64)
		if errUnit != nil {
			log.Fatalf("unable to parse unit: %v", errUnit)
			break
		}

		fmt.Println(instruction)

		switch command {
		case "forward":
			horizontal = unit + horizontal
			depth = aim*unit + depth
		case "down":
			aim = aim + unit
		case "up":
			aim = aim - unit
		}

		fmt.Println(horizontal)
		fmt.Println(aim)
		fmt.Println(depth)
		fmt.Println("========")
	}

	fmt.Println(horizontal * depth)
}
