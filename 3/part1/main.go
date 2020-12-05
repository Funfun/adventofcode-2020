package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Case struct {
	Word     string
	Char     string
	MaxCount int
	MinCount int
}

func parseInputLineToCase(line string) []int {
	raw := []int{}
	for _, char := range line {
		e := string(char)
		switch e {
		case ".":
			raw = append(raw, 0)
		case "#":
			raw = append(raw, 1)
		}
	}

	return raw
}

func main() {
	input, err := os.Open("input.txt")
	// input, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	raws := [][]int{}
	for scanner.Scan() {
		raw := parseInputLineToCase(scanner.Text())
		raws = append(raws, raw)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	// Right 3, down 1.
	treeCount := 0
	moveRight := 0
	for i := 1; i < len(raws); i++ {
		moveRight += 3
		var a int
		if moveRight < len(raws[i]) {
			a = raws[i][moveRight]
		} else {
			moveRight = moveRight - len(raws[i])
			a = raws[i][moveRight]
		}
		if a == 1 {
			treeCount++
		}
	}

	fmt.Printf("Tree count: %d\n", treeCount)
}
