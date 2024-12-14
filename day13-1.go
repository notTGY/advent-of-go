package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binSearch(x, y, x_req, y_req int) int {
	if x_req%x != 0 || y_req%y != 0 {
		return -1
	}
	Bx := x_req / x
	By := y_req / y
	if Bx == By {
		return Bx
	}
	return -1
}

func getMinTokens(x1, y1, x2, y2, x3, y3 int) int {
	/*
	  fmt.Printf("X: %d; Y: %d\n", x1, y1)
	  fmt.Printf("X: %d; Y: %d\n", x2, y2)
	  fmt.Printf("X: %d; Y: %d\n", x3, y3)
	*/

	A_max := min(x3/x1, y3/y1)
	tokens := 0
	for A := 0; A < A_max+1; A++ {
		x_req := x3 - A*x1
		y_req := y3 - A*y1
		if x_req < 0 || y_req < 0 {
			continue
		}

		B := binSearch(x2, y2, x_req, y_req)
		if B >= 0 {
			sum := 3*A + B
			if tokens == 0 || sum < tokens {
				tokens = sum
			}
		}
	}
	//fmt.Printf("tokens: %d\n", tokens)
	return tokens
}

func main() {
	data, _ := os.ReadFile("day13.input")
	blocks := strings.Split(string(data), "\n\n")
	//blocks = blocks[:len(blocks)-1]
	blocks[len(blocks)-1] = blocks[len(blocks)-1][:len(blocks[len(blocks)-1])-1]

	total := 0
	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		line1 := lines[0]
		xStr1 := line1[strings.Index(line1, "X+")+2 : strings.Index(line1, ",")]
		x1, _ := strconv.Atoi(string(xStr1))
		yStr1 := line1[strings.Index(line1, "Y+")+2:]
		y1, _ := strconv.Atoi(string(yStr1))

		line2 := lines[1]
		xStr2 := line2[strings.Index(line2, "X+")+2 : strings.Index(line2, ",")]
		x2, _ := strconv.Atoi(string(xStr2))
		yStr2 := line2[strings.Index(line2, "Y+")+2:]
		y2, _ := strconv.Atoi(string(yStr2))

		line3 := lines[2]
		xStr3 := line3[strings.Index(line3, "X=")+2 : strings.Index(line3, ",")]
		x3, _ := strconv.Atoi(string(xStr3))
		yStr3 := line3[strings.Index(line3, "Y=")+2:]
		y3, _ := strconv.Atoi(string(yStr3))

		total += getMinTokens(x1, y1, x2, y2, x3, y3)
	}

	fmt.Printf("total: %d\n", total)
}
