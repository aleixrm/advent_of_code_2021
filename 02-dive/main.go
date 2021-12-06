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
	"strings"
)

const FilePath = "./data/input.txt"
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
	var depth = 0
	var horizontal = 0
	var aim = 0
	for line, _, err:= reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		command := strings.Split(string(line), " ")
		direction := command[0]
		units, err := strconv.Atoi(command[1])
		check(err)

		switch direction {
		case "forward":
			horizontal += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	finalPos := horizontal * depth
	fmt.Println("final position: ", finalPos)
}