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

	// sum += part1(scanner)
	sum += part2(scanner)
	fmt.Println(sum)
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

	mergeSort(left, 0, len(left)-1)
	mergeSort(right, 0, len(right)-1)

	sum := 0

	for i := range len(left) {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}

func part2(scanner *bufio.Scanner) int {
	left, right := []int{}, []int{}
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		left_num, _ := strconv.Atoi(numbers[0])
		right_num, _ := strconv.Atoi(numbers[1])
		left = append(left, left_num)
		right = append(right, right_num)
	}

	mergeSort(left, 0, len(left)-1)
	mergeSort(right, 0, len(right)-1)

	sum := 0

	for i := range len(left) {
		right_num_greater := false
		similarity_score := 0
		iter := 0
		for right_num_greater == false && iter < len(left) {
			fmt.Println("Left:", left[i])
			fmt.Println("Right:", right[iter])
			if left[i] > right[iter] {
				fmt.Println("Right is smaller")
				fmt.Println()
				iter++
				continue
			} else if left[i] == right[iter] {
				similarity_score++
				iter++
				fmt.Println("Incrementing similarity score to", similarity_score)
			} else {
				right_num_greater = true
				fmt.Println("Right is greater, breaking")
			}
			fmt.Println()
		}
		sum += left[i] * similarity_score
		fmt.Printf("Similarity score for %d is %d\n", left[i], left[i]*similarity_score)
		fmt.Printf("Running sum is %d", sum)
		fmt.Println()
		fmt.Println("------------------------------------------")
		fmt.Println()

	}

	return sum
}

// TODO: Finish mergeSort
func mergeSort(array []int, left, right int) {
	if left >= right {
		return
	}

	mid := int(math.Floor(float64(left+right) / 2))
	mergeSort(array, left, mid)
	mergeSort(array, mid+1, right)
	merge(array, left, mid, right)
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
