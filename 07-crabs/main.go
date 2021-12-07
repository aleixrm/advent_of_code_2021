package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	FilePath = "./data/input.txt"
	PartB = true
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main () {
	fileContent, err := ioutil.ReadFile(FilePath)
	check(err)
	crabs := str2intArray(strings.Split(string(fileContent), ","))
	var totalConsumption, economicPosition int
	if !PartB {
		totalConsumption, economicPosition = computeCrabFuelMethod1(crabs)
	} else {
		totalConsumption, economicPosition = computeCrabFuelMethod2(crabs)
	}
	fmt.Printf("*** total consumption for position %d: %d", economicPosition, totalConsumption)
}

func str2intArray(strArray []string) (intArray []int) {
	for _, str := range strArray {
		intVal, err := strconv.Atoi(str)
		check(err)
		intArray = append(intArray, intVal)
	}
	return
}

func computeCrabFuelMethod1(crabPositions []int) (int, int) {
	sort.Ints(crabPositions)
	medianValue := int(math.Floor(float64(len(crabPositions)/2)))
	economicPosition := crabPositions[medianValue]
	return computeFuelConstantConsumptionToPosition(crabPositions, economicPosition), economicPosition
}

func computeFuelConstantConsumptionToPosition(crabPositions []int, position int) int {
	totalConsumption := 0
	for _, pos := range crabPositions {
		totalConsumption += int(math.Abs(float64(pos - position)))
	}
	return totalConsumption
}

/*** PART B ***/

func computeCrabFuelMethod2(crabPositions []int) (int, int) {
	sort.Ints(crabPositions)
	economicPosition := computeAvgPosition(crabPositions)
	return computeFuelConsumptionToPosition(crabPositions, economicPosition), economicPosition
}

func computeAvgPosition(crabPositions []int) int {
	avg := 0
	for _, pos := range crabPositions {
		avg += pos
	}
	return int(math.Ceil(float64(avg / len(crabPositions))))
}

func computeFuelConsumptionToPosition(crabPositions []int, position int) int {
	totalConsumption := 0
	for _, pos := range crabPositions {
		distance := int(math.Abs(float64(pos - position)))
		totalConsumption += distanceConsumption(distance)
	}
	return totalConsumption
}

func distanceConsumption(distance int) int {
	f := 0
	for i:=1;i<=distance;i++ {
		f += i
	}
	return f
}