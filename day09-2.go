package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("day09.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	line := lines[0]

	total_length := 0
	var digits []int
	for _, c := range line {
		dig, _ := strconv.Atoi(string(c))
		digits = append(digits, dig)
		total_length += dig
	}

	var memory []int
	for i, dig := range digits {
		for j := 0; j < dig; j++ {
			val := -1
			if i%2 == 0 {
				val = i / 2
			}
			memory = append(memory, val)
		}
	}

	for last_id := (len(digits) - 1) / 2; last_id > 0; last_id-- {

		//fmt.Printf("%d: %v\n", last_id, memory)
		l := digits[2*last_id]

		if l == 0 {
			continue
		}
		pos := 0

		found := false
		for !found {
			avail_start := slices.Index(memory[pos:], -1)
			if avail_start == -1 {
				break
			}
			pos += avail_start
			avail_end := slices.IndexFunc(
				memory[pos:],
				func(n int) bool {
					return n >= 0
				},
			)
			if avail_end == -1 {
				avail_end = len(memory[pos:])
			}
			if avail_end < l {
				pos += avail_end
				continue
			}
			found = true
		}

		cpy_start := slices.Index(memory, last_id)
		if cpy_start < pos || !found {
			continue
		}

		for i := 0; i < l; i++ {
			memory[pos+i] = last_id
			memory[cpy_start+i] = -1
		}
	}
	//fmt.Printf("%v\n", memory)

	total := 0
	for i, v := range memory {
		s := 0
		if v != -1 {
			s = i * v
		}
		total += s
	}
	fmt.Printf("total: %d\n", total)
}

// 10180210073447 - too high
// 6311986989618 - too high

// 0620201 <- use this to understand why
// 32211.
// 012345
// 024340 = 13
