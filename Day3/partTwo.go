package main

import (
	"fmt"
	"os"
)

func partTwo(file *os.File) {
	scanner := setupScanner(file)

	var chars [][]byte

	for scanner.Scan() {
		chars = append(chars, scanner.Bytes())
	}

	total := 0

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars[0]); j++ {
			if isStar(chars[i][j]) {
				gearCoordinates := gearNumberCoordinates(chars, coordinate{i, j})
				if len(gearCoordinates) > 0 {
					var ratio = 1
					for num, _ := range gearCoordinates {
						ratio *= resolveNumber(chars, num)
					}
					total += ratio
				}
			}
		}
	}

	fmt.Println("Total Part Two: ", total)

}

func gearNumberCoordinates(chars [][]byte, coord coordinate) map[coordinate]struct{} {
	neighbours := findNeighbours(coord, dimension{len(chars), len(chars[0])})
	var digitsNeighbouringSymbols []coordinate
	for _, neighbour := range neighbours {
		isDigit := isDigit(chars[neighbour.x][neighbour.y])
		if isDigit {
			digitsNeighbouringSymbols = append(digitsNeighbouringSymbols, neighbour)
		}
	}

	numbers := resolveStartingCoordinates(chars, digitsNeighbouringSymbols)

	if len(numbers) == 2 {
		return numbers
	}
	return map[coordinate]struct{}{}
}

func isStar(char byte) bool {
	return char == 42
}
