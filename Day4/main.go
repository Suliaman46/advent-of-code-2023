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
	id             int
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

	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}

	partTwo(file)
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
	id := resolveCardId(line)
	pipeIndex := bytes.IndexByte(line, 124)
	return card{id, resolveWinningNumbers(line, pipeIndex), resolveCardNumbers(line, pipeIndex)}
}

func resolveCardId(line []byte) int {
	// id is immediately before the colon
	idOnesPlaceIdx := bytes.IndexByte(line, 58) - 1
	idLen := idOnesPlaceIdx - 4

	var id = 0
	for i := 0; i < idLen; i++ {
		id *= 10

		elm := line[5+i]
		if elm == 32 {
			continue
		}

		id += int(elm - 48)
	}
	return id
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

func partTwo(file *os.File) {
	var cardTally []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		card := parseCard(scanner.Bytes())
		intersection := intersectNumbers(card)
		addToTally(&cardTally, card.id, len(intersection))
	}

	total := 0
	for _, val := range cardTally {
		total += val
	}

	fmt.Println("Total Part Two: ", total)

}

func addToTally(cardTally *[]int, cardId int, intersectionLength int) {
	exists := cardId <= len(*cardTally)
	if exists {
		handleCardCopies(cardTally, cardId, intersectionLength, (*cardTally)[cardId-1])
	}

	for i := cardId; i <= cardId+intersectionLength; i++ {
		addOrUpdateCardCount(cardTally, i)
	}
}

func handleCardCopies(cardTally *[]int, cardId int, intersectionLength int, cur int) {
	for j := 0; j < cur; j++ {
		for i := cardId + 1; i <= cardId+intersectionLength; i++ {
			addOrUpdateCardCount(cardTally, i)
		}
	}
}

func addOrUpdateCardCount(cardTally *[]int, cardId int) {
	exists := cardId <= len(*cardTally)
	if exists {
		(*cardTally)[cardId-1] = (*cardTally)[cardId-1] + 1
		return
	}
	*cardTally = append(*cardTally, 1)
}
