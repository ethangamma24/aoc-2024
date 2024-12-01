package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadEntireFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func FileLinesScanner(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return bufio.NewScanner(file)
}
