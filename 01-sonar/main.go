/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const FilePath = "./data/sonar.txt"
const WinSize = 3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open(FilePath)
	defer file.Close()
	check(err)

	reader := bufio.NewReader(file)
	var increments = 0
	var lineNum = 0
	window := make([]int, WinSize+1)
	for line, _, err:= reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		current, err := strconv.Atoi(string(line))
		check(err)

		window[lineNum%WinSize+1] = current

		if lineNum >= WinSize {
			prev := sum(window, lineNum-WinSize%WinSize+1, WinSize)
			curr := sum(window, lineNum-WinSize-1%WinSize+1, WinSize)
			if prev < curr {
				increments++
			}
		}
		lineNum++
	}
	fmt.Println("number of increments: ", increments)
}

func sum(window []int, from int, winSize int) int {
	sum := 0
	for i:=from; i<from+winSize; i++ {
		sum += window[i%len(window)]
	}
	return sum
}
