package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

const FilePath = "./data/input.txt"

func check(err error) {
	if err != nil {
		panic(fmt.Errorf("fatal error: %s", err))
	}
}

func main() {
	file, err := os.Open(FilePath)
	check(err)

	defer file.Close()

	var sequenceList [][]byte
	reader := bufio.NewReader(file)
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		sequence := make([]byte, len(line))
		copy(sequence, line)
		sequenceList = append(sequenceList, sequence)
	}

	oxygenGenerationRating := checkOxygenBitCriteria(sequenceList)
	co2ScrubberRating := checkCO2ScrubberRatingCriteria(sequenceList)
	lifeSupportRating := oxygenGenerationRating * co2ScrubberRating
	fmt.Printf("\n*** Life Support Rating: %d ***\n", lifeSupportRating)
}

func getMostCommonValue(counter []int, position int) int {
	if counter[position] < 0 {
		return 0
	} else {
		return 1
	}
}

func getLeastCommonValue(counter []int, position int) int {
	if counter[position] < 0 {
		return 1
	} else {
		return 0
	}
}

func processSequenceList(sequenceList [][]byte, position int) ([2][]int, []int) {
	numBits := len(sequenceList[0])
	var references[2][]int
	var counter = make([]int, numBits)
	for seqNum := 0; seqNum<len(sequenceList); seqNum++ {
		sequence := sequenceList[seqNum]
		if sequence[position] == '1' {
			counter[position]++
			references[1] = append(references[1], seqNum)
		} else if sequence[position] == '0' {
			counter[position]--
			references[0] = append(references[0], seqNum)
		} else {
			check(fmt.Errorf("unexpected binary value! : %s", string(sequence[position])))
		}
	}
	return references, counter
}

func checkOxygenBitCriteria(sequenceList [][]byte) int {
	numBits := len(sequenceList[0])
	for i:=0;i<numBits;i++ {
		references, counter := processSequenceList(sequenceList, i)
		mostCommonValue := getMostCommonValue(counter, i)
		if len(references[mostCommonValue]) > 1 {
			var newSequences [][]byte
			for _, seqNum := range references[mostCommonValue] {
				newSequences = append(newSequences, sequenceList[seqNum])
			}
			sequenceList = newSequences
		} else if len(references[mostCommonValue]) == 1 {
			reference := references[mostCommonValue][0]
			return binarySequenceToDecimal(sequenceList[reference])
		} else if len(references[mostCommonValue]) == 0 {
			break
		}
	}
	check(fmt.Errorf("unexpected error: no valid numbers left"))
	return -1
}

func checkCO2ScrubberRatingCriteria(sequenceList [][]byte) int {
	numBits := len(sequenceList[0])
	for i:=0;i<numBits;i++ {
		references, counter := processSequenceList(sequenceList, i)
		leastCommonValue := getLeastCommonValue(counter, i)
		if len(references[leastCommonValue]) > 1 {
			var newSequences [][]byte
			for _, seqNum := range references[leastCommonValue] {
				newSequences = append(newSequences, sequenceList[seqNum])
			}
			sequenceList = newSequences
		} else if len(references[leastCommonValue]) == 1 {
			reference := references[leastCommonValue][0]
			return binarySequenceToDecimal(sequenceList[reference])
		} else if len(references[leastCommonValue]) == 0 {
			break
		}
	}
	check(fmt.Errorf("unexpected error: no valid numbers left"))
	return -1
}

func binarySequenceToDecimal(sequence []byte) int {
	var decimalValue = 0
	for i:=0;i<len(sequence); i++ {
		if sequence[i] == '1' {
			decimalValue += int(math.Exp2(float64(len(sequence)-1-i)))
		}
	}
	return decimalValue
}