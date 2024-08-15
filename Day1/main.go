package main

import (
	"bufio"
	"fmt"
	"log"
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

	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}
	partTwo(file)
}

func partOne(file *os.File) {
	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		var first string
		var last string
		for _, val := range bytes {
			if val > 47 && val < 58 {
				first = strconv.Itoa(int(val - 48))
				break
			}
		}

		for i := len(bytes) - 1; i >= 0; i-- {
			val := bytes[i]
			if val > 47 && val < 58 {
				last = strconv.Itoa(int(val - 48))
				break
			}
		}

		joined, _ := strconv.Atoi(first + last)
		total += joined
	}

	fmt.Printf("Total Part 1: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partTwo(file *os.File) {
	possibleDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	var total int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		firstIdxOfDigits := make(map[string]int)
		lastIndexOfDigits := make(map[string]int)

		for key, _ := range possibleDigits {
			firstIdx := strings.Index(line, key)
			lastIdx := strings.LastIndex(line, key)

			if firstIdx != -1 {
				firstIdxOfDigits[key] = firstIdx
			}
			if lastIdx != -1 {
				lastIndexOfDigits[key] = lastIdx
			}
		}

		last := resolveLastDigit(lastIndexOfDigits)

		first := resolveFirstDigit(firstIdxOfDigits)

		joined, _ := strconv.Atoi(possibleDigits[first] + possibleDigits[last])
		total += joined
	}

	fmt.Printf("Total Part 2: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func resolveFirstDigit(firstIndexOfInteresting map[string]int) string {
	var first string
	smallest := math.MaxInt
	for key, value := range firstIndexOfInteresting {
		if value < smallest {
			smallest = value
			first = key
		}
	}
	return first
}

func resolveLastDigit(lastIndexOfInteresting map[string]int) string {
	var last string
	largest := -1
	for key, value := range lastIndexOfInteresting {
		if value > largest {
			largest = value
			last = key
		}
	}
	return last
}
