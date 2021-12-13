package day01

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Solve() {
	content, err := ioutil.ReadFile("./day01/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	strContent := string(content)
	array := strings.Split(strContent, "\n")
	count := 0

	for i := 1; i < len(array); i++ {
		current, errCurrent := strconv.ParseInt(array[i], 10, 64)
		if errCurrent != nil {
			log.Fatalf("unable to parse current line: %v", errCurrent)
		}
		previous, errPrevious := strconv.ParseInt(array[i-1], 10, 64)
		if errPrevious != nil {
			log.Fatalf("unable to parse current line: %v", errPrevious)
		}

		if current > previous {
			count++
		}
	}
	fmt.Println(count)
}

func Solve2() {
	content, err := ioutil.ReadFile("./day01/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	strContent := string(content)
	strArray := strings.Split(strContent, "\n")
	arrLen := len(strArray)
	var slice []int64
	for i := 0; i < arrLen; i++ {
		current, errCurrent := strconv.ParseInt(strArray[i], 10, 64)
		if errCurrent != nil {
			log.Fatalf("unable to parse current line: %v", errCurrent)
		}

		slice = append(slice, current)
	}

	var previousSum int64
	count := 0
	for i := 0; i < len(slice); i++ {
		if i+2 > len(slice)-1 {
			break
		}

		if i == 0 {
			previousSum = slice[i] + slice[i+1] + slice[i+2]
		} else {
			sum := slice[i] + slice[i+1] + slice[i+2]
			if sum > previousSum {
				count++
			}
			previousSum = sum
		}
	}

	fmt.Println(count)
}
