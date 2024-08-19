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

type MappingRule struct {
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

	partTwo(file)
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

func shouldMapNumber(num int, mappingDetails MappingRule) bool {
	return num >= mappingDetails.sourceStart && num < mappingDetails.sourceStart+mappingDetails.rangeLen
}

func resolveMap(line string) MappingRule {
	// [0] is destination start , [1] is source start, and [2] is range
	mapsNums := strings.Split(line, " ")
	destStart, _ := strconv.Atoi(mapsNums[0])
	sourceStart, _ := strconv.Atoi(mapsNums[1])
	rangeLen, _ := strconv.Atoi(mapsNums[2])
	return MappingRule{destStart, sourceStart, rangeLen}
}

func resolveInitialDiscreteSeedValues(line string, baseNumbers *[]int, numsInTransition *[]int) {
	for _, num := range strings.Split(strings.SplitAfter(line, ":")[1][1:], " ") {
		n, _ := strconv.Atoi(num)
		*baseNumbers = append(*baseNumbers, n)
		*numsInTransition = append(*numsInTransition, n)
	}

}

func partTwo(file *os.File) {
	var rangeStarts []int
	var rangeLens []int
	var rangeStartsInTransition []int
	var rangeLensInTransition []int

	mappingRules := parseMappingRules(file, &rangeStarts, &rangeLens)
	copyTransformation(&rangeStartsInTransition, &rangeStarts)
	copyTransformation(&rangeLensInTransition, &rangeLens)

	for _, rule := range mappingRules {
		for _, line := range rule {
			splitRangesIfNeeded(line, &rangeStarts, &rangeLens, &rangeStartsInTransition, &rangeLensInTransition)
		}
		for _, line := range rule {
			applyMappingToRange(line, &rangeStarts, &rangeStartsInTransition)
		}
		// copy transition to main
		copyTransformation(&rangeStarts, &rangeStartsInTransition)
		copyTransformation(&rangeLens, &rangeLensInTransition)
	}

	lowest := math.MaxInt
	for _, location := range rangeStarts {
		if location < lowest {
			lowest = location
		}
	}

	fmt.Println("Lowest Part Two: ", lowest)

}

func parseMappingRules(file *os.File, rangeStarts *[]int, rangeLens *[]int) [][]string {
	scanner := bufio.NewScanner(file)
	idx := 0
	var mappingRules [][]string
	temp := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if idx == 0 {
			resolveInitialContSeedValues(line, rangeStarts, rangeLens)
			idx++
			continue
		}

		if len(line) != 0 {
			if isMapNameLine(line) {
				if len(temp) > 0 {
					mappingRules = append(mappingRules, temp)
					temp = make([]string, 0)
				}
			} else {
				temp = append(temp, line)
			}
		}

		idx++
	}
	return mappingRules
}

func copyTransformation(destination *[]int, source *[]int) {
	*destination = append((*source)[:0:0], *source...)
}

func resolveInitialContSeedValues(line string, rangeStarts *[]int, rangeLens *[]int) {

	for i, num := range strings.Split(strings.SplitAfter(line, ":")[1][1:], " ") {
		n, _ := strconv.Atoi(num)
		if i%2 == 0 {
			*rangeStarts = append(*rangeStarts, n)
			continue
		}
		*rangeLens = append(*rangeLens, n)
	}
}

func applyMappingToRange(line string, rangeStarts *[]int, rangeStartsInTransition *[]int) {
	rule := resolveMap(line)

	for i, start := range *rangeStarts {
		if shouldMapNumber(start, rule) {
			(*rangeStartsInTransition)[i] = rule.destStart + (start - rule.sourceStart)
		}
	}

}

func splitRangesIfNeeded(line string, rangeStarts *[]int, rangeLens *[]int, rangeStartsInTransition *[]int, rangeLensInTransition *[]int) {
	rule := resolveMap(line)

	var lengthOfFirstRange int
	var startOfSecondRange int
	var lengthOfSecondRange int

	for i, start := range *rangeStarts {
		currentRange := (*rangeLens)[i]

		if doesRuleOverlapToTheRightWithRange(start, rule, currentRange) {
			// Split Case 1:
			// |________________| -> Source Rule range
			//            |___________________|  -> Seed Range

			// After Split
			// |__________|___________________| -> Seed Range

			lengthOfFirstRange = rule.sourceStart + rule.rangeLen - start
			startOfSecondRange = rule.sourceStart + rule.rangeLen
			lengthOfSecondRange = start + currentRange - (rule.sourceStart + rule.rangeLen)
		} else if doesRuleOverlapToTheLeftWithRange(rule, start, currentRange) {
			{
				// Split Case 2:
				// 				|________________| -> Source Rule range
				//|___________________|  -> Seed Range

				// After Split
				//|_____________|________________| -> Seed Range

				lengthOfFirstRange = rule.sourceStart - start
				startOfSecondRange = rule.sourceStart
				lengthOfSecondRange = currentRange - lengthOfFirstRange
			}

			(*rangeLens)[i] = lengthOfFirstRange
			*rangeStarts = append(*(rangeStarts), startOfSecondRange)
			*rangeLens = append(*rangeLens, lengthOfSecondRange)
		}

	}

	copyTransformation(rangeStartsInTransition, rangeStarts)
	copyTransformation(rangeLensInTransition, rangeLens)
}

func doesRuleOverlapToTheRightWithRange(start int, rule MappingRule, currentRange int) bool {
	// x|________________|y -> Source Rule range
	//            a|___________________|b  -> Seed Range
	// checks if x < a < y && b > y

	return start > rule.sourceStart && start < rule.sourceStart+rule.rangeLen && currentRange > rule.sourceStart+rule.rangeLen
}

func doesRuleOverlapToTheLeftWithRange(rule MappingRule, start int, currentRange int) bool {
	// Split Case 2:
	// 				x|________________|y -> Source Rule range
	//a|___________________|b  -> Seed Range
	// checks if a < x < b && y > b
	return rule.sourceStart > start && rule.sourceStart < start+currentRange && rule.sourceStart+rule.rangeLen > start+currentRange
}
