package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sum += part1(scanner)
	// sum += part2(scanner.Text())
	fmt.Println(sum)
}

func part1(scanner *bufio.Scanner) int {

}

func part2(scanner *bufio.Scanner) int {

}
