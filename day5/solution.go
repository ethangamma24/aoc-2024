package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test_input.txt")
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

	// sum += part1(scanner)
	sum += part2(scanner)
	fmt.Println(sum)
}

func part1(scanner *bufio.Scanner) int {
	sum := 0
	ordering_rules := make(map[string][]string)
	instruction_break := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			instruction_break = true
		}
		if instruction_break == false {
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
			// For each page, see if it's a key in ordering_rules. If so,
			// iterate over its values to see if all of its pages have already
			// been seen. If not, return 0. Else, add page to seen_pages. After
			// reaching end of input, return 1.
			updates_out_of_order := false
			pages := strings.Split(scanner.Text(), ",")
			for index, page := range pages {
				if updates_out_of_order {
					break
				}
				preceding_pages, exists := ordering_rules[page]
				if exists && !updates_out_of_order {
					for i := index + 1; i < len(pages); i++ {
						if updates_out_of_order {
							break
						}
						for _, p := range preceding_pages {
							if pages[i] == p {
								updates_out_of_order = true
								break
							}
						}
					}
				}
			}
			if !updates_out_of_order && len(pages) > 1 {
				mid, err := strconv.Atoi(pages[len(pages)/2])
				if err != nil {
					fmt.Println(err)
				} else {
					sum += mid
				}
			}
		}
	}
	return sum
}

func part2(scanner *bufio.Scanner) int {
	sum := 0
	ordering_rules := make(map[string][]string)
	instruction_break := false
	unordered := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			instruction_break = true
			for k, v := range ordering_rules {
				fmt.Println(k, v)
			}
		}
		if instruction_break == false {
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
			unordered = false
			pages := strings.Split(scanner.Text(), ",")
			slices.SortFunc(pages, func(a, b string) int {
				if contains(ordering_rules[a], b) {
					return 1
				}
				if contains(ordering_rules[b], a) {
					unordered = true
					return -1
				}
				return 0
			})
			if len(pages) > 1 {
				mid, err := strconv.Atoi(pages[len(pages)/2])
				if err != nil {
					fmt.Println(err)
				} else if unordered == true {
					sum += mid
				}
			}
		}
	}
	return sum
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
