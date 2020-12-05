package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	numbers := []int{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	x, y := findPair(numbers, 2020)

	fmt.Println(numbers[x], numbers[y])
	fmt.Println(numbers[x] * numbers[y])

	x, y, z := findTriple(numbers, 2020)

	fmt.Println(numbers[x], numbers[y], numbers[z])
	fmt.Println(numbers[x] * numbers[y] * numbers[z])
}

// [1, 3, 2]
// [4, 2]
func findTriple(list []int, targetSum int) (int, int, int) {
	set := [][]int{}
	x, y, z := 0, 0, 0
	for i := 0; i < len(list); i++ {
		set = append(set, []int{targetSum - list[i], i})
	}

	for j := 0; j < len(set); j++ {
		k1, k2 := findPair(list, set[j][0])

		if list[k1]+list[k2]+list[set[j][1]] == targetSum {
			return k1, k2, set[j][1]
		}
	}

	return x, y, z
}

func findPair(list []int, targetSum int) (int, int) {
	set := [][]int{}
	x, y := 0, 0
	for i := 0; i < len(list); i++ {
		if x != 0 && y != 0 {
			break
		}
		for j := 0; j < len(set); j++ {
			if set[j][0] == list[i] {
				x = set[j][1]
				y = i
				break
			}
		}

		set = append(set, []int{targetSum - list[i], i})
	}

	return x, y
}
