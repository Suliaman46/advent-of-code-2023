package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

type dimension struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
		return
	}
	defer file.Close()

	partOne(file)
}

func partOne(file *os.File) {
	info, _ := file.Stat()
	var maxSize int
	scanner := bufio.NewScanner(file)
	maxSize = int(info.Size())
	buffer := make([]byte, 0, maxSize*4)
	scanner.Buffer(buffer, maxSize*4)

	var chars [][]byte

	for scanner.Scan() {
		chars = append(chars, scanner.Bytes())
	}

	var symbolCoordinates []coordinate
	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars[0]); j++ {
			if isSymbol(chars[i][j]) {
				symbolCoordinates = append(symbolCoordinates, coordinate{i, j})
			}
		}
	}

	dim := dimension{len(chars), len(chars[0])}

	var digitsNeighbouringSymbols []coordinate
	for _, symbolCord := range symbolCoordinates {
		neighbours := findNeighbours(symbolCord, dim)
		for _, neighbour := range neighbours {
			isDigit := isDigit(chars[neighbour.x][neighbour.y])
			if isDigit {
				digitsNeighbouringSymbols = append(digitsNeighbouringSymbols, neighbour)
			}
		}
	}

	numbers := resolveStartingCoordinates(chars, digitsNeighbouringSymbols)

	total := 0
	for num, _ := range numbers {
		total += resolveNumber(chars, num)

	}

	fmt.Println("Total Part One: ", total)

}

func findNeighbours(cord coordinate, dim dimension) []coordinate {
	var neighbours []coordinate

	xAddends := [3]int{0, 1, -1}
	yAddends := [3]int{0, 1, -1}

	for _, xAddend := range xAddends {
		for _, yAddend := range yAddends {
			possibleNeighbour := coordinate{cord.x + xAddend, cord.y + yAddend}
			if isOutOfBound(possibleNeighbour, dim) {
				continue
			}

			if xAddend == 0 && xAddend == yAddend {
				continue
			}

			neighbours = append(neighbours, coordinate{cord.x + xAddend, cord.y + yAddend})
		}
	}

	return neighbours
}

func isSymbol(char byte) bool {
	return !isDigit(char) && !isPeriod(char)
}

func isDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func isPeriod(char byte) bool {
	return char == 46
}

func isOutOfBound(coord coordinate, dim dimension) bool {
	return coord.x < 0 || coord.x >= dim.x || coord.y < 0 || coord.y >= dim.y
}

func resolveStartingCoordinates(chars [][]byte, coordinates []coordinate) map[coordinate]struct{} {

	startingCoordinates := map[coordinate]struct{}{}

	for _, coord := range coordinates {
		for j := coord.y; j >= 0; j-- {

			if hasNumberEnded(chars, coordinate{coord.x, j - 1}) {
				startingCoordinates[coordinate{coord.x, j}] = struct{}{}
				break
			}
		}
	}

	return startingCoordinates
}

func resolveNumber(chars [][]byte, coord coordinate) int {
	var joined string
	for j := coord.y; j < len(chars[0]); j++ {
		if hasNumberEnded(chars, coordinate{coord.x, j}) {
			break
		}
		joined += strconv.Itoa(int(chars[coord.x][j] - 48))

	}
	num, _ := strconv.Atoi(joined)

	return num
}

func hasNumberEnded(chars [][]byte, coord coordinate) bool {
	dim := dimension{len(chars), len(chars[0])}
	return isOutOfBound(coord, dim) || isPeriod(chars[coord.x][coord.y]) || isSymbol(chars[coord.x][coord.y])
}
