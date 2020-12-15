package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	Name     string
	Val      int
	Executed bool
}

func (c *command) Exec(next int, acc int) (int, int, bool) {
	if c.Executed {
		return acc, next, true
	}

	c.Executed = true

	switch c.Name {
	case "acc":
		acc = acc + c.Val
		next++
	case "jmp":
		next = next + c.Val
	case "nop":
		// no op
		next++
	}

	return acc, next, false
}

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	var commands []command
	for scanner.Scan() {
		commandName, val := parseLine(scanner.Text())
		commands = append(commands, command{commandName, val, false})
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	next := 0
	acc := 0
	multiExec := false

	for next < len(commands) {
		acc, next, multiExec = commands[next].Exec(next, acc)

		if multiExec {
			break
		}
	}

	fmt.Println(acc)
}

func parseLine(line string) (string, int) {
	s := strings.Split(line, " ")
	c, val := s[0], s[1]
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return c, v
}
