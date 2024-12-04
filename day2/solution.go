package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	// sum += part1(scanner)
	sum += part2(scanner)
	fmt.Println(sum)
}

func part1(scanner *bufio.Scanner) int {
	num_safe_reports := 0

	for scanner.Scan() {
		report := scanner.Text()
		if checkReportSafety(report) == true {
			num_safe_reports++
		}
	}

	return num_safe_reports
}

func part2(scanner *bufio.Scanner) int {
	num_safe_reports := 0

	for scanner.Scan() {
		report := scanner.Text()
		if checkReportSafety(report) == true {
			num_safe_reports++
		} else if checkReportSafetyWithRemoval(report) == true {
			num_safe_reports++
		}
	}

	return num_safe_reports
}

func checkReportSafety(report string) bool {
	increasing := false
	decreasing := false
	last := 0

	levels := strings.Split(report, " ")

	for i, v := range levels {
		num, _ := strconv.Atoi(v)
		// Handle base case
		if i == 0 {
			last = num
			// Handle i == 1
		} else if i == 1 {
			if num < last && (last-num) <= 3 {
				decreasing = true
			} else if num > last && (num-last) <= 3 {
				increasing = true
			} else {
				// Return false because the numbers are equal
				return false
			}
			last = num
		} else {
			if decreasing {
				if num >= last || (last-num) > 3 {
					return false
				}
			} else if increasing {
				if num <= last || (num-last) > 3 {
					return false
				}
			}
			last = num
		}
	}
	return true
}

func checkReportSafetyWithRemoval(report string) bool {
	last := 0

	levels := strings.Split(report, " ")

	for skipped_level := range levels {
		increasing := false
		decreasing := false
		// Make a copy of the levels slice so we don't modify the original
		newLevels := make([]string, len(levels))
		copy(newLevels, levels)
		newLevels = append(newLevels[:skipped_level], newLevels[skipped_level+1:]...)

		failure := false
		for i, v := range newLevels {
			num, _ := strconv.Atoi(v)
			if i == 0 {
				last = num
			} else if i == 1 {
				if num < last && (last-num) <= 3 {
					decreasing = true
				} else if num > last && (num-last) <= 3 {
					increasing = true
				} else {
					failure = true
				}
				last = num
			} else {
				if decreasing {
					if num >= last || (last-num) > 3 {
						failure = true

					}
				} else if increasing {
					if num <= last || (num-last) > 3 {
						failure = true

					}
				}
				last = num
			}
		}
		if failure == false {
			return true
		}
	}
	return false
}
