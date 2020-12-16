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
	Index    int
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

var swap = map[string]string{
	"jmp": "nop",
	"nop": "jmp",
}

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	var commands []*command
	idx := 0
	for scanner.Scan() {
		commandName, val := parseLine(scanner.Text())
		commands = append(commands, &command{commandName, val, false, idx})
		idx++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	next := 0
	acc := 0
	multiExec := false
	execCommands := []command{}
	for next < len(commands) {
		c := command{commands[next].Name, commands[next].Val, false, next}
		execCommands = append(execCommands, c)
		acc, next, multiExec = commands[next].Exec(next, acc)

		if multiExec {
			commands[next].Executed = false
			break
		}
	}

	fmt.Println("acc before loop", acc)

	for i := len(execCommands) - 1; i >= 0; i-- {
		if execCommands[i].Name == "jmp" || execCommands[i].Name == "nop" {
			k := execCommands[i].Index
			commands[k].Name = swap[commands[k].Name]

			for _, v := range commands {
				v.Executed = false
			}

			multiExec = false
			acc = 0
			next = 0
			for next < len(commands) {
				acc, next, multiExec = commands[next].Exec(next, acc)

				if multiExec {
					commands[k].Name = swap[commands[k].Name]
					commands[k].Executed = false
					break
				}
			}

			if !multiExec {
				fmt.Println("exit")
				break
			}
		}
	}

	fmt.Println("verify run")
	for _, v := range commands {
		v.Executed = false
	}
	next = 0
	acc = 0
	for next < len(commands) {
		acc, next, multiExec = commands[next].Exec(next, acc)
	}

	fmt.Println("final acc", acc)
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
