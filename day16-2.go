package main

import (
	"fmt"
	"os"
	"strings"
  "slices"

	"image"
	"image/color"
	"image/png"
)

func print_tiles(
  lines []string,
  costs [][]int,
  tiles [][2]int,
) {
  block_size := 10

	width := len(lines[0])*block_size
	height := len(lines)*block_size
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}
	gray := color.RGBA{100, 100, 100, 0xff}
	yellow := color.RGBA{200, 200, 100, 0xff}

  for row, line := range lines {
    for col, cell := range line {
      var cl color.Color = gray
      if cell == 'S' || cell == 'E' {
        cl = yellow
        fmt.Printf("%s", string(cell))
      } else if slices.IndexFunc(tiles, func (x [2]int) bool {
        return x[0] == col && x[1] == row
      }) != -1 {
        fmt.Printf("O")
        cl = cyan
      } else if cell == '#' {
        cl = color.Black
        fmt.Printf("%s", string(cell))
      } else if costs[row][col] != -1 {
        fmt.Printf(",")
      } else {
        fmt.Printf("%s", string(cell))
      }

      for i := 0; i < block_size; i++ {
        for j := 0; j < block_size; j++ {
          img.Set(col*block_size+i, row*block_size+j, cl)
        }
      }
    }
    fmt.Printf("\n")
  }

	f, _ := os.Create("day16.png")
	png.Encode(f, img)
}

func get_input() ([]string) {
	data, _ := os.ReadFile("day16.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func get_test() ([]string) {
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

  new_pos := [2]int { pos[0], pos[1] }
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

func appendUnique(a [][2]int, b [2]int) [][2]int {
  if slices.IndexFunc(a, func (x [2]int) bool {
    return x[0] == b[0] && x[1] == b[1]
  }) == -1 {
    return append(a, b)
  }
  return a
}


func backprop(
  lines []string,
  costs [][]int,
  costs_backwards [][]int,
  facing [][]int,
  facing_backwards [][]int,
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
      temp_costs, _ := djikstra(
        lines,
        [2]int{col, row},
        target,
        facing[row][col],
      )
      c_bw := temp_costs[target[1]][target[0]]


      path := c + c_bw
      //fmt.Printf("%d\n", path)

      if path == total {
        tiles = append(tiles, [2]int{col, row})
      }
    }
  }

  return tiles
}

func djikstra (
  lines []string,
  start_pos [2]int,
  end_pos [2]int,
  facing_start int,
) ([][]int, [][]int) {
  costs := [][]int{}
  facing := [][]int{}

  for row, line := range lines {
    costs = append(costs, []int{})
    facing = append(facing, []int{})
    for _, _ = range line {
      costs[row] = append(costs[row], -1)
      facing[row] = append(facing[row], -1)
    }
  }

  positions := [][2]int{}

  positions = append(positions, start_pos)
  facing[start_pos[1]][start_pos[0]] = facing_start
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

      cur_facing := facing[row][col]
      cur_cost := costs[row][col]

      var new_diff_cost int
      var new_facing int
      var new_cost int

      new_pos := [2]int { col+1, row }

      new_diff_cost, new_facing = get_cost(
        lines, pos, cur_facing, 0, 1,
      )
      new_cost = cur_cost + new_diff_cost
      if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
        costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
        facing[new_pos[1]][new_pos[0]] = new_facing
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col-1, row }

      new_diff_cost, new_facing = get_cost(
        lines, pos, cur_facing, 0, -1,
      )
      new_cost = cur_cost + new_diff_cost
      if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
        costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
        facing[new_pos[1]][new_pos[0]] = new_facing
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col, row-1 }

      new_diff_cost, new_facing = get_cost(
        lines, pos, cur_facing, 1, -1,
      )
      new_cost = cur_cost + new_diff_cost
      if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
        costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
        facing[new_pos[1]][new_pos[0]] = new_facing
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col, row+1 }

      new_diff_cost, new_facing = get_cost(
        lines, pos, cur_facing, 1, 1,
      )
      new_cost = cur_cost + new_diff_cost
      if new_diff_cost != -1 && (new_cost <= costs[new_pos[1]][new_pos[0]] || costs[new_pos[1]][new_pos[0]] == -1) {
        costs[new_pos[1]][new_pos[0]] = cur_cost + new_diff_cost
        facing[new_pos[1]][new_pos[0]] = new_facing
        new_positions = append(
          new_positions, new_pos,
        )
      }

    }

    positions = new_positions
  }

  return costs, facing
}

func main() {
	lines := get_input()

  end_pos := [2]int{ -1, -1 }
  start_pos := [2]int{ -1, -1 }

  for row, line := range lines {
    for col, cell := range line {
      pos := [2]int { col, row }
      if cell == 'S' {
        start_pos = pos
      }
      if cell == 'E' {
        end_pos = pos
      }
    }
  }

  costs, facing := djikstra(
    lines,
    start_pos,
    end_pos,
    1,
  )
  facing_end := facing[end_pos[1]][end_pos[0]]
  costs_backwards, facing_backwards := djikstra(
    lines,
    end_pos,
    start_pos,
    facing_end,
  )


  tiles := backprop(
    lines,
    costs,
    costs_backwards,
    facing,
    facing_backwards,
    start_pos,
    end_pos,
  )
  print_tiles(lines, costs, tiles)
  total := len(tiles)

	fmt.Printf("total: %d\n", total)
}
