package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/* input example
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
*/
type Case struct {
	Word      string
	Char      string
	Positions []int
}

func parseInputLineToCase(line string) Case {
	splitted := strings.Split(line, " ")
	counters, char, word := splitted[0], splitted[1], splitted[2]

	splitted = strings.Split(counters, "-")
	pos1, _ := strconv.Atoi(splitted[0])
	pos2, _ := strconv.Atoi(splitted[1])
	char = strings.Split(char, ":")[0]

	return Case{Word: word, Char: char, Positions: []int{pos1, pos2}}
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
		for j := 0; j < len(cases[i].Positions); j++ {
			if cases[i].Char == string(cases[i].Word[cases[i].Positions[j]-1]) {
				count++
			}
		}
		if count == 1 {
			fmt.Println(cases[i].Word, cases[i].Char, cases[i].Positions)
			validCount++
		}
	}

	fmt.Printf("Number of valid password: %d\n", validCount)
}
