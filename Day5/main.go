package main

import (
	"aoc23/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MappingDetails struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
		return
	}
	defer file.Close()

	partOne(file)

	utils.SeekToFileStart(file)
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	idx := 0
	var baseNumbers []int
	var numsInTransition []int
	for scanner.Scan() {
		line := scanner.Text()

		if idx == 0 {
			resolveInitialDiscreteSeedValues(line, &baseNumbers, &numsInTransition)
			idx++
			continue
		}

		if len(line) != 0 {
			if isMapNameLine(line) {
				//copy transition to main array i.e. Mapping stage Completed
				copy(baseNumbers, numsInTransition)
			} else {
				applyMapping(line, baseNumbers, numsInTransition)
			}
		}
		idx++
	}

	copy(baseNumbers, numsInTransition)

	lowest := math.MaxInt
	for _, location := range baseNumbers {
		if location < lowest {
			lowest = location
		}
	}

	fmt.Println("Lowest Part One: ", lowest)
}

func isMapNameLine(line string) bool {
	return line[0] < 48 || line[0] > 57
}

func applyMapping(line string, locations []int, inTransition []int) {
	//map number line
	mappingDetails := resolveMap(line)
	for i := 0; i < len(locations); i++ {
		num := locations[i]

		//If num falls in range to be Mapped
		if shouldMapNumber(num, mappingDetails) {
			inTransition[i] = mappingDetails.destStart + (num - mappingDetails.sourceStart)
		}
	}
}

func shouldMapNumber(num int, mappingDetails MappingDetails) bool {
	return num >= mappingDetails.sourceStart && num < mappingDetails.sourceStart+mappingDetails.rangeLen
}

func resolveMap(line string) MappingDetails {
	// [0] is destination start , [1] is source start, and [2] is range
	mapsNums := strings.Split(line, " ")
	destStart, _ := strconv.Atoi(mapsNums[0])
	sourceStart, _ := strconv.Atoi(mapsNums[1])
	rangeLen, _ := strconv.Atoi(mapsNums[2])
	return MappingDetails{destStart, sourceStart, rangeLen}
}

func resolveInitialDiscreteSeedValues(line string, baseNumbers *[]int, numsInTransition *[]int) {
	for _, num := range strings.Split(strings.SplitAfter(line, ":")[1][1:], " ") {
		n, _ := strconv.Atoi(num)
		*baseNumbers = append(*baseNumbers, n)
		*numsInTransition = append(*numsInTransition, n)
	}

}
