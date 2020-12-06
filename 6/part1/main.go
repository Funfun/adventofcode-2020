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

	allGroups := [][][]string{}
	group := [][]string{}

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

		allGroups = append(allGroups, group)
		group = [][]string{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	total := 0
	// allGroups.all?{|group| group.each_person_answers_are_equal }
	for i := 0; i < len(allGroups); i++ {
		thegroup := allGroups[i]
		cross := map[string]int{}
		for j := 0; j < len(thegroup); j++ {
			answers := thegroup[j]

			for k := 0; k < len(answers); k++ {
				if c, ok := cross[answers[k]]; ok {
					cross[answers[k]] = c + 1
				} else {
					cross[answers[k]] = 1
				}
			}
		}

		groupTotal := 0

		for _, v := range cross {
			if v == len(thegroup) {
				groupTotal++
				total++
			}
		}
		fmt.Println(groupTotal, "----------")
	}

	fmt.Printf("total: %v\n", total)
}
