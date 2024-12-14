package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func safeGetV(
	m [][]int,
	row, col, nrows, ncols int,
) int {
	if row >= nrows ||
		row < 0 ||
		col < 0 ||
		col >= ncols {
		return -1
	}
	return m[row][col]
}

func reach(
	m [][]int,
	row, col, nrows, ncols int,
) (res [][2]int) {
	v := safeGetV(m, row, col, nrows, ncols)
	if v == -1 {
		return res
	}

	rp := safeGetV(m, row+1, col, nrows, ncols)
	rm := safeGetV(m, row-1, col, nrows, ncols)
	cp := safeGetV(m, row, col+1, nrows, ncols)
	cm := safeGetV(m, row, col-1, nrows, ncols)

	if rp == v+1 {
		res = append(res, [2]int{row + 1, col})
	}
	if rm == v+1 {
		res = append(res, [2]int{row - 1, col})
	}
	if cp == v+1 {
		res = append(res, [2]int{row, col + 1})
	}
	if cm == v+1 {
		res = append(res, [2]int{row, col - 1})
	}
	return res
}

func countPos(
	m [][]int,
	row, col, nrows, ncols int,
) int {
	var positions [][2]int
	positions = append(positions, [2]int{row, col})
	for i := 0; i < 9; i++ {
		var newPos [][2]int
		for _, pos := range positions {
			reachable := reach(
				m,
				pos[0], pos[1], nrows, ncols,
			)

			for _, p := range reachable {
				if slices.Index(newPos, p) == -1 {
					newPos = append(newPos, p)
				}
			}
		}
		positions = newPos
	}
	return len(positions)
}

func main() {
	data, _ := os.ReadFile("day10.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	nrows := len(lines)
	ncols := len(lines[0])

	var topoMap [][]int

	for row, line := range lines {
		topoMap = append(topoMap, []int{})
		for _, c := range line {
			s := string(c)
			x, err := strconv.Atoi(s)
			if err != nil {
				x = -1
			}
			topoMap[row] = append(topoMap[row], x)
		}
	}

	total := 0
	for row, mapRow := range topoMap {
		for col, cell := range mapRow {
			if cell == 0 {
				total += countPos(
					topoMap, row, col, nrows, ncols,
				)
			}
		}
	}

	fmt.Printf("total: %d\n", total)
}
