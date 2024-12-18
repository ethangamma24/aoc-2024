package main

import (
	"bufio"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, _ := os.Open("test_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expected := 0

	result := part2(scanner)
	if result != expected {
		t.Errorf("Result was %d, expected %d", result, expected)
	}

}

func TestPart2(t *testing.T) {
	file, _ := os.Open("test_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expected := 0

	result := part2(scanner)
	if result != expected {
		t.Errorf("Result was %d, expected %d", result, expected)
	}
}
