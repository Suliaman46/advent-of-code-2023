package main

import (
	"aoc23/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExtrapolationDir int

const (
	Forward ExtrapolationDir = iota
	Backward
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

	partTwo(file)

}

func partTwo(file *os.File) {
	sequences := parseSequences(file)
	var total = 0
	for i := 0; i < len(sequences); i++ {
		total += extrapolatePrevVal(sequences[i])
	}

	fmt.Println("Total Part Two", total)

}

func partOne(file *os.File) {
	sequences := parseSequences(file)
	var total = 0
	for i := 0; i < len(sequences); i++ {
		total += extrapolateNextVal(sequences[i])
	}

	fmt.Println("Total Part One", total)
}

func extrapolateNextVal(sequence []int) int {
	var res []int

	extrapolateNextSeq(sequence, &res, Forward)
	var total = 0
	for _, num := range res {
		total += num
	}

	return total + sequence[len(sequence)-1]
}

func extrapolatePrevVal(sequence []int) int {
	var res []int

	extrapolateNextSeq(sequence, &res, Backward)
	var curr = 0

	for i := len(res) - 1; i >= 0; i-- {

		curr = res[i] - curr
	}

	return sequence[0] - curr
}

func extrapolateNextSeq(sequence []int, res *[]int, extrapolationDir ExtrapolationDir) *[]int {
	var temp []int
	var diff int
	var anyNonZeroNum = false
	for i := 1; i < len(sequence); i++ {
		diff = sequence[i] - sequence[i-1]
		if diff != 0 {
			anyNonZeroNum = true
		}
		temp = append(temp, diff)
	}

	if extrapolationDir == Forward {
		*res = append(*res, temp[len(temp)-1])
	} else {
		*res = append(*res, temp[0])
	}

	if !anyNonZeroNum {
		return res
	}

	return extrapolateNextSeq(temp, res, extrapolationDir)
}

func parseSequences(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var sequences [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var temp []int
		for _, n := range line {
			num, _ := strconv.Atoi(n)
			temp = append(temp, num)
		}

		sequences = append(sequences, temp)
	}

	return sequences
}
