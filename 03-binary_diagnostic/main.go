package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const FilePath = "./data/input.txt"
//const FilePath = "./data/test_input.txt"
const NumBits = 12

func check(err error) {
	if err != nil {
		panic(fmt.Errorf("fatal error: %s", err))
	}
}

func main() {
	fmt.Printf(os.Getwd())
	file, err := os.Open(FilePath)
	check(err)

	defer file.Close()

	var count = make([]int, NumBits)
	reader := bufio.NewReader(file)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		checkBytes(&count, line)
	}

	binaryGamma, gamma := getGamma(count)
	epsilon := getEpsilon(binaryGamma)
	consumption := gamma * epsilon
	fmt.Printf("submarine consumption: %d", consumption)
}

func checkBytes(count *[]int, line []byte)  {
	for i:=0;i<len(*count); i++ {
		if line[i] == '1' {
			(*count)[i]++
		} else if line[i] == '0'{
			(*count)[i]--
		} else {
			check(fmt.Errorf("unexpected binary value! : %s", string(line[i])))
		}
	}
}

func getGamma(count []int) ([]byte, int) {
	gamma := 0
	binaryGamma := make([]byte, NumBits)
	for i:=0;i<len(count); i++ {
		switch val := count[i]; {
		case val > 0:
			binaryGamma[i] = 1
			gamma += int(math.Exp2(float64(len(count)-1-i)))
		case val < 0:
			binaryGamma[i] = 0
		case val == 0:
			check(fmt.Errorf("unexpected error: binary draw! "))
		}
	}
	return binaryGamma, gamma
}

func getEpsilon(binaryGamma []byte) int {
	epsilon := 0
	for i:=0;i<len(binaryGamma); i++ {
		if binaryGamma[i] == 0 {
			epsilon += int(math.Exp2(float64(len(binaryGamma)-1-i)))
		}
	}
	return epsilon
}
