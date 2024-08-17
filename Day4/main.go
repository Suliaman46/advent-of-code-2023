package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

type card struct {
	winningNumbers []int
	cardNumbers    []int
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
	var total = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCard(scanner.Bytes())
		intersection := intersectNumbers(card)
		total += resolvePoints(intersection)
	}
	fmt.Println("Total Part One: ", total)
}

func resolvePoints(intersection []int) int {
	return int(math.Pow(2, float64(len(intersection)-1)))
}

func intersectNumbers(card card) []int {
	interMap := map[int]struct{}{}

	var intersection []int

	for _, number := range card.winningNumbers {
		interMap[number] = struct{}{}
	}

	for _, number := range card.cardNumbers {
		_, present := interMap[number]
		if present {
			intersection = append(intersection, number)
		}
	}

	return intersection
}

func parseCard(line []byte) card {
	pipeIndex := bytes.IndexByte(line, 124)
	return card{resolveWinningNumbers(line, pipeIndex), resolveCardNumbers(line, pipeIndex)}
}

func resolveCardNumbers(line []byte, pipeIndex int) []int {
	var cardNumbers []int
	lineLength := len(line)

	for i := pipeIndex + 1; i < lineLength; i++ {
		if (i-pipeIndex)%3 == 2 {

			//whitespace character
			if line[i] == 32 {
				cardNumbers = append(cardNumbers, int(line[i+1]-48))
				i++
				continue
			}

			cardNumbers = append(cardNumbers, resolveTwoDigitNumber(line, i))
		}
	}
	return cardNumbers
}

func resolveWinningNumbers(line []byte, pipeIndex int) []int {
	var winningNumbers []int
	colonIndex := bytes.IndexByte(line, 58)
	for i := colonIndex + 1; i < pipeIndex-1; i++ {
		if (i-colonIndex)%3 == 2 {
			//whitespace character
			if line[i] == 32 {
				winningNumbers = append(winningNumbers, int(line[i+1]-48))
				i++
				continue
			}
			winningNumbers = append(winningNumbers, resolveTwoDigitNumber(line, i))
		}
	}
	return winningNumbers
}

func resolveTwoDigitNumber(line []byte, idx int) int {
	num := strconv.Itoa(int(line[idx] - 48))
	num += strconv.Itoa(int(line[idx+1] - 48))
	number, _ := strconv.Atoi(num)

	return number
}
