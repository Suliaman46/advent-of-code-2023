package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
	}
	scanner := bufio.NewScanner(file)

	var total int
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

	fmt.Printf("Total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
