package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var XMAS_MAP [][]rune
var XMAS []rune
var XMAS_FOUND int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	XMAS_MAP = createXmasMap(scanner)
	XMAS = []rune{'X', 'M', 'A', 'S'}
	XMAS_FOUND = 0

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1()
	// part2()
	fmt.Println(XMAS_FOUND)
}

func part1() {
	// Starting at 1st index bc we're searching for 'X' at surface level
	xmas_index := 1

	for y, row := range XMAS_MAP {
		for x, char := range row {
			if char == 'X' {
				fmt.Printf("Found %c, searching for %c\n", 'X', XMAS[xmas_index])
				searchRecursively(x, y, xmas_index)
			}
		}
	}
}

func part2() {

}

func searchRecursively(x, y, index int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// fmt.Printf("Looking at [%d][%d]\n", j, i)
			char, err := getAtIndex(i, j)
			if err != nil {
				fmt.Println(err)
			} else {
				if char == XMAS[index] {
					// fmt.Printf("Checking %c against %c\n", char, XMAS[index])
					if char == 'S' {
						fmt.Println("Found S, incrementing XMAS_FOUND")
						XMAS_FOUND++
						return
					} else {
						fmt.Printf("Found %c, searching for %c\n", XMAS[index], XMAS[index+1])
						searchRecursively(i, j, index+1)
					}
				}
			}
		}
	}
}

func getAtIndex(x, y int) (rune, error) {
	if x < 0 || y < 0 || x >= len(XMAS_MAP) || y >= len(XMAS_MAP[x]) {
		return ' ', errors.New("Index out of bounds")
	}
	return XMAS_MAP[y][x], nil
}

func createXmasMap(scanner *bufio.Scanner) [][]rune {
	xmas_map := make([][]rune, 0)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		xmas_map = append(xmas_map, line)
	}
	return xmas_map
}
