package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Case struct {
	Word     string
	Char     string
	MaxCount int
	MinCount int
}

func parseInputLineToCase(line string) Case {
	splitted := strings.Split(line, " ")
	counters, char, word := splitted[0], splitted[1], splitted[2]

	splitted = strings.Split(counters, "-")
	minCount, _ := strconv.Atoi(splitted[0])
	maxCount, _ := strconv.Atoi(splitted[1])
	char = strings.Split(char, ":")[0]

	return Case{Word: word, Char: char, MaxCount: maxCount, MinCount: minCount}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	cases := []Case{}
	for scanner.Scan() {
		thecase := parseInputLineToCase(scanner.Text())
		cases = append(cases, thecase)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	// go over each case
	// How many passwords are valid according to their policies?
	validCount := 0
	for i := 0; i < len(cases); i++ {
		count := 0
		for _, char := range cases[i].Word {
			if string(char) == cases[i].Char {
				count++
			}
		}
		if count >= cases[i].MinCount && count <= cases[i].MaxCount {
			validCount++
		}
	}

	fmt.Printf("Number of valid password: %d\n", validCount)
}
