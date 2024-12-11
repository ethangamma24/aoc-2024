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

	XMAS_FOUND += part1()
	// XMAS_FOUND += part2()
	fmt.Println(XMAS_FOUND)
}

func part1() int {
	sum := 0
	// Starting at 2nd index bc we're searching for 'X' at surface level, then
	// when we find an 'M' around it, we'll be searching in a straight line
	xmas_index := 1

	for y, row := range XMAS_MAP {
		for x, char := range row {
			fmt.Printf("Looking for %c at [%d][%d]\n", 'X', y+1, x+1)
			if char == 'X' {
				fmt.Printf("Found %c at [%d][%d], searching for %c\n", 'X', y+1, x+1, XMAS[xmas_index])
				// searchRecursively(x, y, xmas_index)
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						char, err := getAtIndex(i, j)
						fmt.Printf("Looking at [%d][%d]: %c\n", j+1, i+1, char)
						if err != nil {
							fmt.Println(err)
						} else {
							// If we find M, continue searching in that "line"
							// for the last 2 characters
							// fmt.Printf("Found %c at [%d][%d], searching for %c\n", XMAS[xmas_index], j+1, i+1, XMAS[xmas_index+1])
							if char == 'M' {
								fmt.Printf("Found %c at [%d][%d], searching for %c\n", XMAS[xmas_index], j+1, i+1, XMAS[xmas_index+1])
								sum += searchStraightLine(i, j, i-x, j-y, xmas_index+1)
							}
						}
					}
				}
			}
		}
	}
	return sum
}

func part2() int {
	return 0
}

func searchStraightLine(x, y, dir_x, dir_y, index int) int {
	new_x, new_y := x+dir_x, y+dir_y
	fmt.Printf("Direction: %d, %d\n", dir_y, dir_x)
	fmt.Printf("Looking at [%d][%d]\n", new_y+1, new_x+1)
	char, err := getAtIndex(new_x, new_y)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if char == XMAS[index] {
		if char == 'S' {
			fmt.Println("Found S, incrementing XMAS_FOUND")
			fmt.Println("--------------------------------")
			return 1
		} else {
			fmt.Printf("Found %c at [%d][%d], searching for %c\n", XMAS[index], new_y+1, new_x+1, XMAS[index+1])
			return searchStraightLine(new_x, new_y, dir_x, dir_y, index+1)
		}
	} else {
		fmt.Printf("Found %c, exiting\n", char)
		return 0
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
