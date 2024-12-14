package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() ([]string, int, int, int) {
	data, _ := os.ReadFile("day14.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	nrows := 103
	ncols := 101
	nsteps := 100
	return lines, nrows, ncols, nsteps
}

func get_test() ([]string, int, int, int) {
	data, _ := os.ReadFile("day14.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	nrows := 7
	ncols := 11
	nsteps := 100
	return lines, nrows, ncols, nsteps
}

func main() {
	lines, nrows, ncols, nsteps := get_input()

	midrow := (nrows - 1) / 2
	midcol := (ncols - 1) / 2

	//fmt.Printf("%d %d\n\n", midcol, midrow)

	quads := [4]int{0, 0, 0, 0}
	for _, line := range lines {
		var p [2]int
		var v [2]int
		for i, pov_i := range strings.Split(line, " ") {
			coordStrings := strings.Split(
				strings.Split(pov_i, "=")[1],
				",",
			)
			var coords [2]int
			for j, coordString := range coordStrings {
				coord, _ := strconv.Atoi(coordString)
				coords[j] = coord
			}
			if i == 0 {
				p = coords
			} else {
				v = coords
			}
		}
		// p, v

		if v[0] < 0 {
			v[0] += ncols
		}
		if v[1] < 0 {
			v[1] += nrows
		}

		x := (p[0] + nsteps*v[0]) % ncols
		y := (p[1] + nsteps*v[1]) % nrows
		//x = (x+2*ncols)%ncols
		//y = (y+2*nrows)%nrows

		//fmt.Printf("%d %d\n%d %d\n\n", x, y, p[0], p[1])

		x -= midcol
		y -= midrow

		//fmt.Printf("%v\n", v)

		selected_quad := -1
		switch {
		case x > 0 && y > 0:
			selected_quad = 0
		case x > 0 && y < 0:
			selected_quad = 3
		case x < 0 && y > 0:
			selected_quad = 1
		case x < 0 && y < 0:
			selected_quad = 2
		default:
			if x != 0 && y != 0 {
				fmt.Printf("fail\n")
			}
		}
		//fmt.Printf("%d %d\nquad %d\n", x, y, selected_quad)
		if selected_quad != -1 {
			quads[selected_quad]++
		}
	}

	total := 1
	for _, quad := range quads {
		total *= quad
	}
	fmt.Printf("%v\n", quads)

	fmt.Printf("total: %d\n", total)
}
