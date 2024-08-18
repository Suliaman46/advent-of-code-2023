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
	var locations []int
	var inTransition []int
	for scanner.Scan() {
		line := scanner.Text()

		if idx == 0 {
			for _, num := range strings.Split(strings.SplitAfter(line, ":")[1][1:], " ") {
				n, _ := strconv.Atoi(num)
				locations = append(locations, n)
				inTransition = append(inTransition, n)
			}
			idx++
			continue
		}

		if len(line) != 0 {
			if line[0] < 48 || line[0] > 57 {
				//map name line

				//copy transition to main array
				copy(locations, inTransition)
			} else {
				//map number line

				// [0] is destination start , [1] is source start, and [2] is range
				mapsNums := strings.Split(line, " ")
				destStart, _ := strconv.Atoi(mapsNums[0])
				sourceStart, _ := strconv.Atoi(mapsNums[1])
				rangeLen, _ := strconv.Atoi(mapsNums[2])
				for i := 0; i < len(locations); i++ {
					//If falls in range to be Mapped
					elm := locations[i]

					if elm >= sourceStart && elm < sourceStart+rangeLen {

						inTransition[i] = destStart + (elm - sourceStart)

					}
				}

			}
		}
		idx++
	}

	copy(locations, inTransition)

	lowest := math.MaxInt
	for _, location := range locations {
		if location < lowest {
			lowest = location
		}
	}

	fmt.Println("Lowest Part One: ", lowest)
}
