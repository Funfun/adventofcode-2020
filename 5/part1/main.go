package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

const maxRows = 127 // total
const maxCol = 7

func main() {
	input, err := os.Open("input.txt")
	// input, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	maxSeat := 0
	allSeats := []int{}
	for scanner.Scan() {
		pass := scanner.Text()
		theRow := []int{0, maxRows}
		theCol := []int{0, maxCol}
		finalRow := 0
		finalCol := 0
		for pos, char := range pass {
			if pos < len(pass)-3 {
				row := string(char)
				// fmt.Println("ROW", row)
				switch row {
				case "F":
					f := (theRow[1] - theRow[0]) / 2
					if f == 0 {
						finalRow = theRow[0]
						continue
					}
					theRow = []int{theRow[0], f + theRow[0]}
					// fmt.Println("F: theRow", theRow)
				case "B":
					mid := float64(theRow[1]-theRow[0]) / float64(2)
					b := int(math.Ceil(mid))
					if b == 1 {
						finalRow = theRow[1]
						continue
					}
					theRow = []int{theRow[0] + b, theRow[1]}
					// fmt.Println("B: theRow", theRow)
				}
			} else {
				col := string(char)
				// fmt.Println("COL", col)
				switch col {
				case "L":
					f := (theCol[1] - theCol[0]) / 2
					if f == 0 {
						finalCol = theCol[0]
						continue
					}
					theCol = []int{theCol[0], f + theCol[0]}
					// fmt.Println("L: theCol", theCol)
				case "R":
					mid := float64(theCol[1]-theCol[0]) / float64(2)
					b := int(math.Ceil(mid))
					if b == 1 {
						finalCol = theCol[1]
						continue
					}
					theCol = []int{theCol[0] + b, theCol[1]}
					// fmt.Println("R: theCol", theCol)
				}
			}
		}

		seatID := finalRow*8 + finalCol
		allSeats = append(allSeats, seatID)
		if seatID > maxSeat {
			maxSeat = seatID
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	sort.Slice(allSeats, func(i, j int) bool {
		return allSeats[i] < allSeats[j]
	})

	mySeat := 0

	for i := 1; i < len(allSeats)-1; i++ {
		if allSeats[i]-1 != allSeats[i-1] && allSeats[i-1] != allSeats[1] {
			mySeat = allSeats[i] - 1
			break
		}

		if allSeats[i]+1 != allSeats[i+1] && allSeats[i-1] != allSeats[len(allSeats)-2] {
			mySeat = allSeats[i] + 1
			break
		}
	}

	fmt.Printf("Max seatID: %d\n", maxSeat)
	fmt.Printf("mySeats: %v\n", mySeat)
}
