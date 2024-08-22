package utils

import (
	"bufio"
	"os"
)

func SeekToFileStart(file *os.File) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return
	}
}

func IsDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func SetupScanner(file *os.File) *bufio.Scanner {
	info, _ := file.Stat()
	var maxSize int
	scanner := bufio.NewScanner(file)
	maxSize = int(info.Size())
	buffer := make([]byte, 0, maxSize*4)
	scanner.Buffer(buffer, maxSize*4)
	return scanner
}

