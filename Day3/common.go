package main

import (
	"bufio"
	"os"
)

type coordinate struct {
	x int
	y int
}

type dimension struct {
	x int
	y int
}

func setupScanner(file *os.File) *bufio.Scanner {
	info, _ := file.Stat()
	var maxSize int
	scanner := bufio.NewScanner(file)
	maxSize = int(info.Size())
	buffer := make([]byte, 0, maxSize*4)
	scanner.Buffer(buffer, maxSize*4)
	return scanner
}
