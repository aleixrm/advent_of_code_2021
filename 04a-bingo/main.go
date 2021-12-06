package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const FilePath = "./data/input.txt"

type numberLocations map[string][]location

type hitsBoard [][][]int // <0 if there is a hit

type location struct {
	board int
	row int
	col int
}


func check(err error) {
	if err != nil {
		panic(fmt.Errorf("unexpected error: %s", err))
	}
}

func main() {
	file, err := os.Open(FilePath)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	check(err)
	numSequence := strings.Split(string(line), ",")
	reader.ReadLine() //discard blank line
	row := 0
	boardNum := 0
	numberMap := numberLocations{}
	var hits hitsBoard
	var boardArray [][]int
	for line, _, err := reader.ReadLine(); err != io.EOF;  {
		if len(line) == 0 {
			line, _, err = reader.ReadLine()
			continue
		}

		lineArray := strings.Fields(string(line))
		rowArray := make([]int, len(lineArray))
		for col, num := range lineArray {
			rowArray[col], err = strconv.Atoi(num)
			check(err)
			numberMap[num] = append(numberMap[num], location{ boardNum, row, col})
		}
		boardArray = append(boardArray, rowArray)
		row++

		line, _, err = reader.ReadLine()
		if len(line) == 0 {
			row = 0
			hits = append(hits, boardArray)
			boardArray = [][]int{}
			boardNum++
		}
	}
	victory, num, result := playBingo(numSequence, hits, numberMap)
	if victory {
		fmt.Printf("BINGO with num: %s! result: %d", num, result)
	} else {
		fmt.Errorf("game finished without winner")
	}
}

func playBingo(numSequence []string, hits hitsBoard, locations numberLocations) (bool, string, int) {
	for _, num := range numSequence {
		for _, loc := range locations[num] {
			hits[loc.board][loc.row][loc.col] = -1
			if checkRow(hits, loc.row, loc.board) || checkCol(hits, loc.col, loc.board) {
				sum := unmarkedSum(hits, loc.board)
				intNum, err := strconv.Atoi(num)
				check(err)
				return true, num, sum * intNum
			}
		}
	}
	return false, "", -1
}

func checkRow(hits hitsBoard, row int, boardNum int) bool {
	for _, num := range hits[boardNum][row] {
		if num >= 0 {
			return false
		}
	}
	return true
}

func checkCol(hits hitsBoard, col int, boardNum int) bool {
	for i:=0;i<len(hits[boardNum]);i++ {
		if hits[boardNum][i][col] >= 0 {
			return false
		}
	}
	return true
}

func unmarkedSum(hits hitsBoard, boardNum int) int {
	sum := 0
	for _, row := range hits[boardNum] {
		for _, num := range row {
			if num >0 {
				sum += num
			}
		}
	}
	return sum
}
