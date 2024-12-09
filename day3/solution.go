package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(scanner.Text())

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

	for scanner.Scan() {
		data := scanner.Text()
		pattern := `mul\([0-9]+,[0-9]+\)`

		re := regexp.MustCompile(pattern)

		matches := re.FindAllString(data, -1)

		for _, v := range matches {
			fmt.Println(v)
			sum += computeMul(v)
		}
	}

	return sum
}

func part2(scanner *bufio.Scanner) int {
	sum := 0

	for scanner.Scan() {
		data := scanner.Text()
		do_dont_map := make(map[int]string, 0)
		mul_pattern := `mul\([0-9]+,[0-9]+\)`
		mul_re := regexp.MustCompile(mul_pattern)

		do_pattern := `[d][o][(][)]`
		dont_pattern := `[d][o][n]['][t][(][)]`

		do_matches := regexp.MustCompile(do_pattern).FindAllStringIndex(data, -1)
		dont_matches := regexp.MustCompile(dont_pattern).FindAllStringIndex(data, -1)

		if dont_matches[len(dont_matches)-1][0] < do_matches[len(do_matches)-1][0] {
			last_elem := []int{len(data) - 1, len(data)}
			dont_matches = append(dont_matches, last_elem)
		}

		// Calculate do() mul instructions until first dont()
		matches := mul_re.FindAllString(data[0:dont_matches[0][0]], -1)

		fmt.Println("Do Matches", do_matches)
		fmt.Println("Don't Matches", dont_matches)

		for _, v := range matches {
			fmt.Println(v)
			sum += computeMul(v)
		}

		// Populate a map containing the exact indexes of each do/dont occurrence
		for _, v := range do_matches {
			do_dont_map[v[0]] = "do"
		}
		for _, v := range dont_matches {
			do_dont_map[v[0]] = "dont"
		}

		// Extract keys from map and sort them
		keys := make([]int, 0, len(do_dont_map))
		for key := range do_dont_map {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		last_dont_index := 0

		// Iterate over map using sorted keys
		for i := 0; i < len(keys); i++ {
			curr_instruction := keys[i]
			// Prevents covering the same instructions we've already computed
			if curr_instruction > last_dont_index {
				if do_dont_map[curr_instruction] == "do" {
					// Look for next dont()
					for j := i + 1; j < len(keys); j++ {
						next_instruction := keys[j]
						if do_dont_map[next_instruction] == "dont" {
							fmt.Printf("curr_instruction: %s, next_instruction: %s", do_dont_map[curr_instruction], do_dont_map[next_instruction])
							fmt.Println()
							matches = mul_re.FindAllString(data[curr_instruction:next_instruction], -1)
							for _, v := range matches {
								sum += computeMul(v)
							}
							last_dont_index = next_instruction
							break
						}
					}
				}
			}
		}
	}

	return sum
}

func computeMul(s string) int {
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	numbers := strings.Split(s[start+1:end], ",")
	left, _ := strconv.Atoi(numbers[0])
	right, _ := strconv.Atoi(numbers[1])
	return (left * right)
}
