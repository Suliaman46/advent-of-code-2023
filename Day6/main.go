package main

import (
	"aoc23/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
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

	utils.SeekToFileStart(file)

	start := time.Now()

	partTwo(file, findTotal)
	
	elapsed := time.Since(start)
	fmt.Println("Part Two by loop solution took: ", elapsed)

	start = time.Now()

	partTwo(file, findTotalByQuadratic)

	elapsed = time.Since(start)
	fmt.Println("Part Two by quadratic solution took: ", elapsed)
}

func partOne(file *os.File) {
	races := parseFile(file, parseLineIntoSeparateRaces)
	total := findTotal(races)

	fmt.Println("Total Part One: ", total)
}

func partTwo(file *os.File, strategy func([]Race) int) {
	races := parseFile(file, parseLineIntoSingleRace)
	total := strategy(races)

	fmt.Println("Total Part Two: ", total)
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

func findTotalByQuadratic(races []Race) int {
	var solStart int
	var solEnd int

	var total = 1

	for _, race := range races {
		startFloat := 0.5 * (float64(race.time) - math.Sqrt(math.Pow(float64(race.time), 2)-float64(4*race.distance)))
		endFloat := 0.5 * (float64(race.time) + math.Sqrt(math.Pow(float64(race.time), 2)-float64(4*race.distance)))

		solStart = handleQuadraticRounding(startFloat, 1)
		solEnd = handleQuadraticRounding(endFloat, -1)

		total *= solEnd - solStart + 1
	}

	return total
}

func handleQuadraticRounding(float float64, round float64) int {
	if float == math.Trunc(float) {
		return int(float + round)
	}
	if round == 1 {
		return int(math.Ceil(float))
	}
	return int(math.Floor(float))
}

func isRecordBroken(i int, race Race) bool {
	return i*((race.time)-i) > race.distance
}

func parseFile(file *os.File, lineParser func([]byte, *[]int)) []Race {
	scanner := bufio.NewScanner(file)
	var num []int
	for scanner.Scan() {
		line := scanner.Bytes()
		lineParser(line, &num)
	}
	return convertToRaces(num)
}

func parseLineIntoSeparateRaces(line []byte, nums *[]int) {
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

func parseLineIntoSingleRace(line []byte, nums *[]int) {
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
	}

	*nums = append(*nums, num)
}

func convertToRaces(nums []int) []Race {
	var races []Race
	split := len(nums) / 2

	for i := 0; i < split; i++ {
		races = append(races, Race{nums[i], nums[i+split]})
	}

	return races
}
