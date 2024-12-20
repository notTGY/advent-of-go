package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"image"
	"image/color"
	"image/png"
)

func print_tiles(
	lines []string,
	costs [][]int,
	tiles [][2]int,
	cheat [][2]int,
) {
	block_size := 10

	width := len(lines[0]) * block_size
	height := len(lines) * block_size
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}
	gray := color.RGBA{100, 100, 100, 0xff}
	yellow := color.RGBA{200, 200, 100, 0xff}
	red := color.RGBA{200, 100, 100, 0xff}

	for row, line := range lines {
		for col, cell := range line {
			var cl color.Color = gray
			if cell == 'S' || cell == 'E' {
				cl = yellow
				fmt.Printf(" %s", string(cell))
			} else if slices.IndexFunc(cheat, func(x [2]int) bool {
				return x[0] == col && x[1] == row
			}) != -1 {
				fmt.Printf(" C")
				cl = red
			} else if slices.IndexFunc(tiles, func(x [2]int) bool {
				return x[0] == col && x[1] == row
			}) != -1 {
				fmt.Printf(" O")
				cl = cyan
			} else if cell == '#' {
				cl = color.Black
				fmt.Printf(" %s", string(cell))
			} else if costs[row][col] != -1 {
				fmt.Printf(" ,")
			} else {
				fmt.Printf(" %s", string(cell))
			}

			for i := 0; i < block_size; i++ {
				for j := 0; j < block_size; j++ {
					img.Set(col*block_size+i, row*block_size+j, cl)
				}
			}
		}
		fmt.Printf("\n")
	}

	f, _ := os.Create("day20.png")
	png.Encode(f, img)
}

func get_input() ([]string, [2]int, [2]int) {
	data, _ := os.ReadFile("day20.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	start_pos := [2]int{-1, -1}
	end_pos := [2]int{-1, -1}
	for row, line := range lines {
		for col, cell := range line {
			pos := [2]int{col, row}
			if cell == 'S' {
				start_pos = pos
			}
			if cell == 'E' {
				end_pos = pos
			}
		}
	}
	return lines, start_pos, end_pos
}
func get_test() ([]string, [2]int, [2]int) {
	data, _ := os.ReadFile("day20.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	start_pos := [2]int{-1, -1}
	end_pos := [2]int{-1, -1}
	for row, line := range lines {
		for col, cell := range line {
			pos := [2]int{col, row}
			if cell == 'S' {
				start_pos = pos
			}
			if cell == 'E' {
				end_pos = pos
			}
		}
	}
	return lines, start_pos, end_pos
}

func get_cost(
	lines []string,
	pos [2]int,
	axis int,
	inc int,
) int {
	new_pos := [2]int{pos[0], pos[1]}
	new_pos[axis] += inc

	nrows := len(lines)
	ncols := len(lines[0])

	if new_pos[0] < 0 || new_pos[0] >= ncols {
		return -1
	}

	if new_pos[1] < 0 || new_pos[1] >= nrows {
		return -1
	}

	if lines[new_pos[1]][new_pos[0]] == '#' {
		return -1
	}

	return 1
}

func djikstra(
	lines []string,
	start_pos [2]int,
) [][]int {
	costs := [][]int{}

	for row, line := range lines {
		costs = append(costs, []int{})
		for _, _ = range line {
			costs[row] = append(costs[row], -1)
		}
	}

	positions := [][2]int{}

	positions = append(positions, start_pos)
	costs[start_pos[1]][start_pos[0]] = 0

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

			cur_cost := costs[row][col]

			var new_diff_cost int
			var new_cost int

			new_pos := [2]int{col + 1, row}

			new_diff_cost = get_cost(
				lines, pos, 0, 1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = new_cost
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col - 1, row}

			new_diff_cost = get_cost(
				lines, pos, 0, -1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = new_cost
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col, row - 1}

			new_diff_cost = get_cost(
				lines, pos, 1, -1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = new_cost
				new_positions = append(
					new_positions, new_pos,
				)
			}

			new_pos = [2]int{col, row + 1}

			new_diff_cost = get_cost(
				lines, pos, 1, 1,
			)
			new_cost = cur_cost + new_diff_cost
			if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
				costs[new_pos[1]][new_pos[0]] = new_cost
				new_positions = append(
					new_positions, new_pos,
				)
			}

		}

		positions = new_positions
	}

	return costs
}

func testPos(
	lines []string,
	pos [2]int,
) bool {
	nrows := len(lines)
	ncols := len(lines[0])
	return pos[0] >= 0 && pos[0] < ncols &&
		pos[1] >= 0 && pos[1] < nrows &&
		lines[pos[1]][pos[0]] != '#'
}

func notInArr(arr [][2]int, b [2]int) bool {
	return slices.IndexFunc(arr, func(x [2]int) bool {
		return x[0] == b[0] && x[1] == b[1]
	}) == -1
}

func getNeighbourhoodFor(
	lines []string,
	pos [2]int,
	max_length int,
) ([][2]int, []int) {
	res := [][2]int{}
	l := []int{}

	for dy := 0; dy <= max_length; dy++ {
		for dx := 0; dx+dy <= max_length; dx++ {
			cur_l := dx + dy
			if cur_l < 2 {
				continue
			}
			candidates := [][2]int{
				[2]int{pos[0] + dx, pos[1] + dy},
				[2]int{pos[0] - dx, pos[1] + dy},
				[2]int{pos[0] + dx, pos[1] - dy},
				[2]int{pos[0] - dx, pos[1] - dy},
			}
			for _, cand := range candidates {
				if testPos(lines, cand) && notInArr(res, cand) {
					res = append(res, cand)
					l = append(l, cur_l)
				}
			}
		}
	}

	return res, l
}

func findCheats(
	lines []string,
	start_pos [2]int,
	end_pos [2]int,
) (int, map[int][][2]int) {
	costs_forward := djikstra(
		lines,
		start_pos,
	)
	total_cost := costs_forward[end_pos[1]][end_pos[0]]

	costs_backward := djikstra(
		lines,
		end_pos,
	)

	cheats := make(map[int][][2]int)
	saves := make(map[int]int)
	total := 0
	for row, line := range lines {
		for col, cell := range line {
			if cell == '#' {
				continue
			}
			pos := [2]int{col, row}

			cost_fw := costs_forward[row][col]

			neighbours, ls := getNeighbourhoodFor(lines, pos, 20)

			for i, cheat_end := range neighbours {
				cheat_len := ls[i]
				// REMOVE this to get solution
				if cheat_len != 2 {
					continue
				}
				cost_bw := costs_backward[cheat_end[1]][cheat_end[0]]

				save := total_cost - (cost_fw + cost_bw + cheat_len)

				if save > 0 {
					_, ok := saves[save]
					if !ok {
						saves[save] = 0
						cheats[save] = [][2]int{
							pos,
							cheat_end,
						}
					}
					saves[save]++
				}
			}

		}
	}

	for key, value := range saves {
		fmt.Printf("%d: %d\n", key, value)

		if key >= 100 {
			total += value
		}
	}

	return total, cheats
}

func backprop(
	lines []string,
	costs [][]int,
	pos [2]int,
	target [2]int,
) [][2]int {
	tiles := [][2]int{}

	total := costs[target[1]][target[0]]

	for row, line := range lines {
		for col, cell := range line {
			if cell == '#' {
				continue
			}

			c := costs[row][col]
			temp_costs := djikstra(
				lines,
				[2]int{col, row},
			)
			c_bw := temp_costs[target[1]][target[0]]
			/*
			   c_bw := costs_backwards[row][col]
			*/

			path := c + c_bw
			if path == total {
				tiles = append(tiles, [2]int{col, row})
			}
		}
	}

	return tiles
}

func main() {
	start := time.Now()
	lines, start_pos, end_pos := get_input()

	total, cheats := findCheats(lines, start_pos, end_pos)

	m := -1
	cheat := [][2]int{}
	for key, value := range cheats {
		if key > m {
			m = key
			cheat = value
		}
	}
	if len(cheat) != 2 {
		fmt.Printf("No cheats")
		return
	}
	cheat_start := cheat[0]
	cheat_end := cheat[1]

	costs_start := djikstra(
		lines,
		start_pos,
	)
	costs_cheat := djikstra(
		lines,
		cheat_end,
	)

	tiles_0 := backprop(
		lines,
		costs_start,
		start_pos,
		cheat_start,
	)

	tiles_1 := backprop(
		lines,
		costs_cheat,
		cheat_end,
		end_pos,
	)

	tiles := append(tiles_0, tiles_1...)

	print_tiles(
		lines,
		costs_start,
		tiles,
		cheat,
	)

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
