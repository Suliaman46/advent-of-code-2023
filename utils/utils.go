package utils

import (
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
