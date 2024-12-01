package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	sum += part1(scanner)
	// fmt.Println(part2())
}

func part1(scanner *bufio.Scanner) int {
	left, right := []int{}, []int{}
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		left_num, _ := strconv.Atoi(numbers[0])
		right_num, _ := strconv.Atoi(numbers[1])
		left = append(left, left_num)
		right = append(right, right_num)
	}

	return 0
}

// TODO: Finish mergeSort
func mergeSort(array []int, left, right int) {
	if left >= right {
		return
	}

	mid := int(math.Floor(float64(left+right) / 2))
}

func merge(array []int, left, mid, right int) {
	len_left, len_right := mid-left+1, right-mid
	left_array, right_array := make([]int, len_left), make([]int, len_right)

	for i := 0; i < len_left; i++ {
		left_array[i] = array[left+i]
	}
	for i := 0; i < len_right; i++ {
		right_array[i] = array[mid+1+i]
	}

	i, j, k := 0, 0, left

	for i < len_left && j < len_right {
		if left_array[i] <= right_array[j] {
			array[k] = left_array[i]
			i++
		} else {
			array[k] = right_array[j]
			j++
		}
		k++
	}

	// Copy remaining values of arrays if there are any
	for i < len_left {
		array[k] = left_array[i]
		i++
		k++
	}

	for j < len_right {
		array[k] = right_array[j]
		j++
		k++
	}
}
