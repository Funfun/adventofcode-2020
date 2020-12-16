package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Min, Max int
}

func (r Rule) Include(n int) bool {
	return n >= r.Min && n <= r.Max
}

var rules = []Rule{
	{43, 237}, {251, 961},
	{27, 579}, {586, 953},
	{31, 587}, {608, 967},
	{26, 773}, {784, 973},
	{41, 532}, {552, 956},
	{33, 322}, {333, 972},
	{30, 165}, {178, 965},
	{31, 565}, {571, 968},
	{36, 453}, {473, 963},
	{35, 912}, {924, 951},
	{39, 376}, {396, 968},
	{31, 686}, {697, 974},
	{28, 78}, {96, 971},
	{32, 929}, {943, 955},
	{40, 885}, {896, 968},
	{26, 744}, {765, 967},
	{46, 721}, {741, 969},
	{30, 626}, {641, 965},
	{48, 488}, {513, 971},
	{34, 354}, {361, 973},
}

func parseLine(line string) []int {
	res := []int{}
	splitted := strings.Split(line, ",")
	for _, s := range splitted {
		v, _ := strconv.Atoi(s)
		res = append(res, v)
	}

	return res
}

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	invalidNumbers := []int{}
	for scanner.Scan() {
		ticket := parseLine(scanner.Text())
		result := invalidTicketNumbers(ticket)
		invalidNumbers = append(invalidNumbers, result...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	total := 0
	for _, e := range invalidNumbers {
		total += e
	}
	fmt.Println(total)
}

func invalidTicketNumbers(numbers []int) []int {
	result := []int{}

	for i := 0; i < len(numbers); i++ {
		number := numbers[i]

		validCounter := 0
		for j := 0; j < len(rules); j++ {
			if rules[j].Include(number) {
				validCounter++
			}
		}

		if validCounter == 0 {
			result = append(result, number)
		}
	}

	return result
}
