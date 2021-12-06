package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	FilePath = "./data/input.txt"
	MaxDays = 256
	MaxDaysToBirth = 8
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fileContent, err := ioutil.ReadFile(FilePath)
	check(err)
	fish := str2intArray(strings.Split(string(fileContent), ","))
	totalFish := reproduceFish(fish)
	fmt.Printf("*** Amount of fish after %d days: %d", MaxDays, totalFish)
}

func str2intArray(strArray []string) (intArray []int) {
	for _, str := range strArray {
		intVal, err := strconv.Atoi(str)
		check(err)
		intArray = append(intArray, intVal)
	}
	return
}

func reproduceFish(fish []int) int {
	var pregnancyWindow = make([]int, MaxDaysToBirth+1)
	for _, fishDays := range fish {
		pregnancyWindow[fishDays]++
	}
	for day:=0; day<MaxDays; day++ {
		newPregnancyWindow := make([]int, MaxDaysToBirth+1)
		for daysToBirth:= MaxDaysToBirth; daysToBirth >=0; daysToBirth-- {
			if daysToBirth == 0 {
				newPregnancyWindow[6] += pregnancyWindow[daysToBirth]
				newPregnancyWindow[8] += pregnancyWindow[daysToBirth]
			} else {
				newPregnancyWindow[daysToBirth-1] = pregnancyWindow[daysToBirth]
			}
		}
		pregnancyWindow = newPregnancyWindow
	}
	totalFish := 0
	for _, numOfFish := range pregnancyWindow {
		totalFish += numOfFish
	}
	return totalFish
}
