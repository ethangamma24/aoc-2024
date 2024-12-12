package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	ordering_rules := make(map[string][]string)
	instruction_break := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			instruction_break = true
		}
		if instruction_break == false {
			fmt.Println(scanner.Text())
			ordering := strings.Split(scanner.Text(), "|")
			key := ordering[1]
			value := ordering[0]
			_, exists := ordering_rules[key]
			if exists {
				ordering_rules[key] = append(ordering_rules[key], value)
			} else {
				ordering_rules[key] = []string{value}
			}
		} else {
			seen_pages := make([]string{})
			// For each page, see if it's a key in ordering_rules. If so,
			// iterate over its values to see if all of its pages have already
			// been seen. If not, return 0. Else, add page to seen_pages. After
			// reaching end of input, return 1.
		}
	}
	return 0
}

func part2(scanner *bufio.Scanner) int {
	return 0
}
