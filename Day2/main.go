package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type game struct {
	id    int
	blue  int
	green int
	red   int
}

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

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
		game := resolveGame(scanner.Text())
		if isLegalGame(game) {
			total += game.id
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func resolveGame(line string) game {
	gameId := resolveGameId(line)

	blueCount := resolveMaxColorAppearance(line, "blue")
	greenCount := resolveMaxColorAppearance(line, "green")
	redCount := resolveMaxColorAppearance(line, "red")

	return game{
		id:    gameId,
		blue:  blueCount,
		green: greenCount,
		red:   redCount}

}

func resolveGameId(line string) int {
	r := regexp.MustCompile(`Game (?P<GameId>\d+):`)
	matches := r.FindAllStringSubmatch(line, -1)
	val, err := strconv.Atoi(matches[0][len(matches[0])-1])
	if err != nil {
		fmt.Println("Error on game id conversion")
		return -1
	}
	return val
}

func resolveMaxColorAppearance(line string, color string) int {
	r := regexp.MustCompile(`(?P<Num>\d+) ` + color)
	matches := r.FindAllStringSubmatch(line, -1)

	largest := -1
	for _, match := range matches {
		val, err := strconv.Atoi(match[len(match)-1])
		if err != nil {
			fmt.Println("Error on color count conversion")
			return -1
		}
		if val > largest {
			largest = val
		}
	}
	return largest
}

func isLegalGame(game game) bool {
	return game.red <= maxRed && game.blue <= maxBlue && game.green <= maxGreen
}
