package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const FilePath = "./data/input.txt"
const ExerciseB = true //enables exercise Part B

type vectorMap map[string]int

func setVector(diagram vectorMap, ori [2]int, dest [2]int) {
	skip, amount, xIncrement, yIncrement := getAmountAndIncrements(ori, dest)
	if skip {
		return
	}
	for i:=0;i<=amount;i++ {
		x := ori[0]+i*xIncrement
		y := ori[1]+i*yIncrement
		diagram[coordinate2string(x,y)]++
	}
}

func getAmountAndIncrements(ori [2]int, dest [2]int) (skip bool, amount int, xIncrement int, yIncrement int) {
	xAmount := dest[0]-ori[0]
	yAmount := dest[1]-ori[1]
	if xAmount == 0 {
		amount = yAmount
		yIncrement = 1
	} else if yAmount == 0 {
		amount = xAmount
		xIncrement = 1
	} else if math.Abs(float64(xAmount)) == math.Abs(float64(yAmount)) && ExerciseB {
		amount = int(math.Abs(float64(xAmount)))
		xIncrement = 1
		yIncrement = 1
	} else {
		return true, 0, 0, 0
	}
	if xAmount < 0 {
		xIncrement = -1
	}
	if yAmount < 0 {
		yIncrement = -1
	}
	if amount < 0 {
		amount *= -1
	}
	return
}

func getNumOfDangerPoints(diagram vectorMap) int {
	numDangerPoints := 0
	for _, v := range diagram {
		if v > 1 {
			numDangerPoints++
		}
	}
	return numDangerPoints
}

func check(err error) {
	if err != nil {
		panic(fmt.Errorf("fatal error: %s", err))
	}
}

func parseCoordinate(coordinate []string) [2]int {
	xCoordinate, err := strconv.Atoi(coordinate[0])
	check(err)
	yCoordinate, err := strconv.Atoi(coordinate[1])
	check(err)
	return [2]int{xCoordinate, yCoordinate}
}

func coordinate2string(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func main() {
	file, err := os.Open(FilePath)
	check(err)

	defer file.Close()

	var diagram vectorMap
	diagram = make(map[string]int)
	reader := bufio.NewReader(file)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		coords := strings.Split(string(line), " -> ")
		origin := strings.Split(coords[0], ",")
		dest := strings.Split(coords[1], ",")
		setVector(diagram, parseCoordinate(origin), parseCoordinate(dest))
	}

	fmt.Printf("\n*** Number of dangerous points: %d ***\n", getNumOfDangerPoints(diagram))
}

