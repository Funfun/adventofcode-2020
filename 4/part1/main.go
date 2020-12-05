package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Field struct {
	Count    int
	Required bool
	Key      string
	Value    string
}

var elcValues = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var color *regexp.Regexp
var pidre *regexp.Regexp

func init() {
	color = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidre = regexp.MustCompile(`^\d{9}$`)
}
func elcContains(s string) bool {
	for i := 0; i < len(elcValues); i++ {
		if s == elcValues[i] {
			return true
		}
	}

	return false
}
func (f *Field) Valid() bool {
	if f.Key == "byr" {
		v, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}

		return v >= 1920 && v <= 2002
	}

	if f.Key == "iyr" {
		v, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}

		return v >= 2010 && v <= 2020
	}

	if f.Key == "eyr" {
		v, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}

		return v >= 2020 && v <= 2030
	}

	if f.Key == "hgt" {
		if strings.Contains(f.Value, "cm") {
			g := strings.Split(f.Value, "cm")
			if len(g) == 0 {
				return false
			}

			v, err := strconv.Atoi(g[0])
			if err != nil {
				return false
			}

			r := v >= 150 && v <= 193

			return r
		}

		if strings.Contains(f.Value, "in") {
			// fmt.Println(f.Key, f.Value)

			g := strings.Split(f.Value, "in")
			if len(g) == 0 {
				return false
			}

			v, err := strconv.Atoi(g[0])
			if err != nil {
				return false
			}

			r := v >= 59 && v <= 76
			if r {
				fmt.Println(f.Key, f.Value)
			}

			return r
		}
	}

	if f.Key == "hcl" {
		return color.MatchString(f.Value)
	}

	if f.Key == "ecl" {
		return elcContains(f.Value)
	}

	if f.Key == "pid" {
		return pidre.MatchString(f.Value)
	}

	return false
}

var keys = map[string]*Field{
	"byr": {0, true, "", ""},
	"iyr": {0, true, "", ""},
	"eyr": {0, true, "", ""},
	"hgt": {0, true, "", ""},
	"hcl": {0, true, "", ""},
	"ecl": {0, true, "", ""},
	"pid": {0, true, "", ""},
	"cid": {0, false, "", ""},
}

func main() {
	input, err := os.Open("input.txt")
	// input, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	validCount := 0
	for scanner.Scan() {
		raw := scanner.Text()
		if raw != "" {
			data := strings.Split(raw, " ")
			for i := 0; i < len(data); i++ {
				kv := strings.Split(data[i], ":")
				if kv[1] == "" {
					continue
				}
				if field, ok := keys[kv[0]]; ok {
					if field.Count == 0 {
						field.Count = 1
						field.Key = kv[0]
						field.Value = kv[1]
					}

					// fmt.Println(*field, field.Valid())
				}
			}
			continue
		}

		fields := 0
		for _, field := range keys {
			if field.Key == "cid" {
				continue
			}

			if field.Count == 1 && field.Required && field.Valid() {
				fields++
				continue
			}
		}

		if fields >= 7 {
			validCount++
			// } else {
			// 	fmt.Println(keys)
		}

		keys = map[string]*Field{
			"byr": {0, true, "", ""},
			"iyr": {0, true, "", ""},
			"eyr": {0, true, "", ""},
			"hgt": {0, true, "", ""},
			"hcl": {0, true, "", ""},
			"ecl": {0, true, "", ""},
			"pid": {0, true, "", ""},
			"cid": {0, false, "", ""},
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	fmt.Printf("Total count: %d\n", validCount)
}
