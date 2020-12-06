package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	group := [][]string{}
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			dic := []string{}
			person := map[string]int{}
			for _, char := range line {
				if _, ok := person[string(char)]; !ok {
					person[string(char)] = 1
					dic = append(dic, string(char))
				}
			}
			// [a, b, c]
			group = append(group, dic)

			continue
		}

		cross := map[string]int{}
		for j := 0; j < len(group); j++ {
			answers := group[j]

			for k := 0; k < len(answers); k++ {
				if c, ok := cross[answers[k]]; ok {
					cross[answers[k]] = c + 1 // counter number of same answers to 'yes'
				} else {
					cross[answers[k]] = 1
				}
			}
		}

		// a = 5 people
		// b = 4
		// group size: 5
		groupTotal := 0
		for _, v := range cross {
			if v == len(group) {
				groupTotal++
				total++
			}
		}
		group = [][]string{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	fmt.Printf("total: %v\n", total)
}
