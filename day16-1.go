package main

import (
	"fmt"
	"os"
	"strings"
)

func get_input() []string {
	data, _ := os.ReadFile("day16.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func get_test() []string {
	data, _ := os.ReadFile("day16.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

/*
new_diff_cost, new_facing = get_cost(
  lines, pos, facing, 0, 1,
)
*/

func get_cost(
	lines []string,
	pos [2]int,
	facing int,
	axis int,
	inc int,
) (int, int) {
	cost := -1

	new_facing := 1
	switch {
	case axis == 0 && inc < 0:
		new_facing = 3
	case axis == 1 && inc < 0:
		new_facing = 2
	case axis == 1 && inc > 0:
		new_facing = 0
	}

	new_pos := [2]int{pos[0], pos[1]}
	new_pos[axis] += inc

	if lines[new_pos[1]][new_pos[0]] == '#' {
		return cost, new_facing
	}

	cost = 1

	rotations := 2
	if facing == new_facing {
		rotations = 0
	}
	if (facing-new_facing)*(facing-new_facing) == 1 {
		rotations = 1
	}

	//rotations := (3+(facing - new_facing)) % 3
	cost += rotations * 1000

	return cost, new_facing
}

func main() {
	lines := get_input()

	end_pos := [2]int{-1, -1}

	costs := [][]int{}
	facing := [][]int{}

	positions := [][2]int{}
	for row, line := range lines {
		costs = append(costs, []int{})
		facing = append(facing, []int{})

		for col, cell := range line {
			pos := [2]int{col, row}
			costs[row] = append(costs[row], -1)
			facing[row] = append(facing[row], -1)

			if cell == 'S' {
				positions = append(positions, pos)
				costs[row][col] = 0
				facing[row][col] = 1
			}
			if cell == 'E' {
				end_pos = pos
			}
		}
	}

	for {
		/*
		   fmt.Printf("%v\n", positions)
		   for _, pos := range positions {
		     fmt.Printf(
		       "%d %d: %d\n",
		       pos[0],
		       pos[1],
		       costs[pos[1]][pos[0]],
		     )
		   }
		*/

		if len(positions) == 0 {
			break
		}
		new_positions := [][2]int{}

		for _, pos := range positions {
			col := pos[0]
			row := pos[1]

			cur_facing := facing[row][col]
			cur_cost := costs[row][col]

			var new_diff_cost int
			var new_facing int
			var new_cost int

			new_pos := [2]int{col + 1, row}

			new_diff_cost, new_facing = get_cost(
				lines, pos, cur_facing, 0, 1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost < costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
				facing[new_pos[1]][new_pos[0]] = new_facing
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col - 1, row}

			new_diff_cost, new_facing = get_cost(
				lines, pos, cur_facing, 0, -1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost < costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
				facing[new_pos[1]][new_pos[0]] = new_facing
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col, row - 1}

			new_diff_cost, new_facing = get_cost(
				lines, pos, cur_facing, 1, -1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost < costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
				facing[new_pos[1]][new_pos[0]] = new_facing
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col, row + 1}

			new_diff_cost, new_facing = get_cost(
				lines, pos, cur_facing, 1, 1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost < costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
				facing[new_pos[1]][new_pos[0]] = new_facing
				new_positions = append(
					new_positions, new_pos,
				)
			}

		}

		positions = new_positions
	}

	total := costs[end_pos[1]][end_pos[0]]

	fmt.Printf("total: %d\n", total)
}
