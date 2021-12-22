package day03

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Solve() {
	content, err := ioutil.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	array := strings.Split(string(content), "\n")
	dictionary := make(map[int][]string)
	for i := 0; i < len(array); i++ {
		line := strings.Split(array[i], "")
		for j := 0; j < len(line); j++ {
			elem := dictionary[j]
			if len(elem) == 0 {
				var slice []string
				slice = append(slice, line[j])
				dictionary[j] = slice
			} else {
				elem = append(elem, line[j])
				dictionary[j] = elem
			}
		}
	}

	gamma := calculateGamma(dictionary)
	epsilon := calculateEpsilon(gamma)
	fmt.Printf("power: %v\n", multiplyBinaryNumbers(gamma, epsilon))
}

func calculateGamma(d map[int][]string) string {
	gamma := ""
	for i := 0; i < len(d); i++ {
		zeroCount := 0
		oneCount := 0
		for j := 0; j < len(d[i]); j++ {
			if d[i][j] == "0" {
				zeroCount++
			}

			if d[i][j] == "1" {
				oneCount++
			}
		}

		if zeroCount > oneCount {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}

	return gamma
}

func calculateEpsilon(gamma string) string {
	epsilon := ""
	chars := strings.Split(gamma, "")
	for i := 0; i < len(chars); i++ {
		if chars[i] == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}
	return epsilon
}

func multiplyBinaryNumbers(x, y string) int64 {
	xInt, xErr := strconv.ParseInt(x, 2, 64)
	yInt, yErr := strconv.ParseInt(y, 2, 64)

	if xErr != nil {
		log.Fatalf("unable to convert x to number: %v", xErr)
	}

	if yErr != nil {
		log.Fatalf("unable to convert y to number: %v", yErr)
	}

	return xInt * yInt
}

func Solve2() {
	content, err := ioutil.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	numbers := strings.Split(string(content), "\n")
	oxygenRating, oxygenError := findOxygenGeneratorRating(numbers, 1)
	if oxygenError != nil {
		log.Fatalf("error: %v", oxygenError)

	}
	fmt.Println("oxygen", oxygenRating)

	co2Rating, co2RatingError := findCO2ScrubberRating(numbers, 1)
	if co2RatingError != nil {
		log.Fatalf("error: %v", co2RatingError)
	}

	fmt.Println("co2", co2Rating)

	fmt.Println("life support rating:", multiplyBinaryNumbers(oxygenRating, co2Rating))

}

func findOxygenGeneratorRating(numbers []string, position int) (string, error) {
	if len(numbers) == 1 {
		return numbers[0], nil
	}

	countZero := 0
	countOne := 0
	var numbersWithZero []string
	var numbersWithOne []string
	for i := 0; i < len(numbers); i++ {
		var maxIndex int
		if position > len(numbers[i]) {
			maxIndex = len(numbers[i])
		} else {
			maxIndex = position
		}
		firstCharacter := numbers[i][maxIndex-1 : maxIndex]
		if firstCharacter == "0" {
			countZero++
			numbersWithZero = append(numbersWithZero, numbers[i])
		} else {
			countOne++
			numbersWithOne = append(numbersWithOne, numbers[i])
		}
	}

	// fmt.Println("Numbers with 0")
	// fmt.Println(countZero)
	// fmt.Println(numbersWithZero)
	// fmt.Println("Numbers with 1")
	// fmt.Println(countOne)
	// fmt.Println(numbersWithOne)

	if countOne >= countZero {
		return findOxygenGeneratorRating(numbersWithOne, position+1)
	} else {
		return findOxygenGeneratorRating(numbersWithZero, position+1)
	}
}

func findCO2ScrubberRating(numbers []string, position int) (string, error) {
	if len(numbers) == 1 {
		return numbers[0], nil
	}

	countZero := 0
	countOne := 0
	var numbersWithZero []string
	var numbersWithOne []string
	for i := 0; i < len(numbers); i++ {
		var maxIndex int
		if position > len(numbers[i]) {
			maxIndex = len(numbers[i])
		} else {
			maxIndex = position
		}
		firstCharacter := numbers[i][maxIndex-1 : maxIndex]
		if firstCharacter == "0" {
			countZero++
			numbersWithZero = append(numbersWithZero, numbers[i])
		} else {
			countOne++
			numbersWithOne = append(numbersWithOne, numbers[i])
		}
	}

	// fmt.Println("Numbers with 0")
	// fmt.Println(countZero)
	// fmt.Println(numbersWithZero)
	// fmt.Println("Numbers with 1")
	// fmt.Println(countOne)
	// fmt.Println(numbersWithOne)

	if countZero <= countOne {
		return findCO2ScrubberRating(numbersWithZero, position+1)
	} else {
		return findCO2ScrubberRating(numbersWithOne, position+1)
	}
}
