package day04

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Board struct {
	puzzle [][]int
	Sum    int
	Id int
}

func (b *Board) Mark(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if n == b.puzzle[i][j] {
				b.puzzle[i][j] = -1
			} else if b.puzzle[i][j] != -1 {
				b.Sum = b.Sum + b.puzzle[i][j]
			}
		}
	}
}

func (b Board) CheckWin() bool {
	checkRowWin := func() bool {
		result := 0 
		for i := 0; i < 5; i++ {
			for _, v := range b.puzzle[i] {  
				result += v  
			   }
			if result == -5 {
				return true
			} else {
				result = 0
			}
		}
		return false
	}

	checkColumnWin := func() bool {
		result := 0 
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				result += b.puzzle[j][i]
			}
			if result == -5 {
				return true
			} else {
				result = 0
			}
		}

		return false
	}

	return checkRowWin() || checkColumnWin()
}

func removeEmptyString(array []string) []string {
	var newArray []string
	for _, v := range array {
		if v != "" {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func createBoard(array []string, id int) Board {
	var puzzle [][]int

	for i := 0; i < 5; i++ {
		line := removeEmptyString(strings.Split(array[i], " "))
		var lineArray []int
		for j := 0; j < len(line); j++ {
			xInt, _ := strconv.Atoi(line[j])
			lineArray = append(lineArray, xInt)
		}
		puzzle = append(puzzle, lineArray)
	}

	return Board{puzzle,0,id}
}

func createBoards(boardsLines []string) []Board {
	var boards []Board

	line := 0
	for line < len(boardsLines) {
		boardLines := boardsLines[line : line+5]
		boards = append(boards, createBoard(boardLines,line))
		line = line + 5
	}

	return boards
}

func Solve() {
	content, err := ioutil.ReadFile("./day04/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	lines := removeEmptyString(strings.Split(string(content), "\n"))
	numbers := strings.Split(lines[0], ",")
	boards := createBoards(lines[1:])

	found := false
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		for _, board := range boards {
			board.Mark(number)
			if board.CheckWin() {
				fmt.Println(board.Sum * number)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func Solve2() {
	content, err := ioutil.ReadFile("./day04/input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	lines := removeEmptyString(strings.Split(string(content), "\n"))
	numbers := strings.Split(lines[0], ",")
	boards := createBoards(lines[1:])

	m := make(map[int]Board)
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		for _, board := range boards {
			if _ , ok := m[board.Id]; !ok {
				board.Mark(number)
				if board.CheckWin() {
					m[board.Id] = board
					if(len(m) == len(boards)){
						fmt.Println(board.Sum * number)
						break
					}
				}

			}
		}
		if len(m) == len(boards) {
			break
		}
	}
}
