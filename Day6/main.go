package main

import (
	"aoc23/utils"
	"bufio"
	"fmt"
	"os"
)

type Race struct {
	time     int
	distance int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
		return
	}
	defer file.Close()

	partOne(file)
}

func partOne(file *os.File) {
	races := parseFile(file)
	total := findTotal(races)

	fmt.Println("Total Part One: ", total)
}

func findTotal(races []Race) int {
	var solStart int
	var solEnd int

	var total = 1
	for _, race := range races {

		for i := 1; i < race.time; i++ {
			if isRecordBroken(i, race) {
				solStart = i
				break
			}
		}

		for i := race.time - 1; i > 0; i-- {
			if isRecordBroken(i, race) {
				solEnd = i
				break
			}
		}

		total *= solEnd - solStart + 1
	}
	return total
}

func isRecordBroken(i int, race Race) bool {
	return i*((race.time)-i) > race.distance
}

func parseFile(file *os.File) []Race {
	scanner := bufio.NewScanner(file)
	var num []int
	for scanner.Scan() {
		line := scanner.Bytes()
		parseLine(line, &num)
	}
	return parseRaces(num)
}

func parseLine(line []byte, nums *[]int) {
	var firstNumIdx int
	for i, char := range line {
		if utils.IsDigit(char) {
			firstNumIdx = i
			break
		}
	}

	num := 0
	for i := firstNumIdx; i < len(line); i++ {
		if utils.IsDigit(line[i]) {
			num *= 10
			num += int(line[i] - 48)
		}

		if num != 0 && !utils.IsDigit(line[i]) {
			*nums = append(*nums, num)
			num = 0
		}
	}

	*nums = append(*nums, num)
}

func parseRaces(nums []int) []Race {
	var races []Race
	split := len(nums) / 2

	for i := 0; i < split; i++ {
		races = append(races, Race{nums[i], nums[i+split]})
	}

	return races
}
