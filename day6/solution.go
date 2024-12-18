package main

import (
	"bufio"
	"fmt"
	"os"
)

type COORDS struct {
	x int
	y int
}

var MAP [][]rune

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	MAP = createMap(scanner)
	DIRECTION := map[string]COORDS{
		"UP": {x: 0, y: -1},
		"RIGHT": {x: 1, y: 0},
		"DOWN": {x: 0, y: 1},
		"LEFT": {x: -1, y: 0},
	}

	sum := 0

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var starting_position COORDS
	starting_position_found := false
	// Find guard's starting location
	for i := 0; i < len(MAP); i++ {
		for j := 0; j < len(MAP[i]); j++ {
			if starting_position_found {
				break
			}
			if MAP[i][j] == '^' {
				starting_position := COORDS{j,i}
				starting_position_found = true
			}
		}
		if starting_position_found {
			break
		}
	}

	sum += part1(starting_position)
	// sum += part2(starting_position)
	fmt.Println(sum)
}

func part1(position COORDS) int {
	
}

func part2(position COORDS) int {

}

func createMap(scanner *bufio.Scanner) [][]rune {
	map := make([][]rune, 0)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		map = append(map, line)
	}
	return map
}
