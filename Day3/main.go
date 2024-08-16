package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
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
