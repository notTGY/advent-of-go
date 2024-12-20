package main

import (
	"fmt"
	"os"
	"strings"
  "time"
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
        fmt.Printf(" %s", string(cell))
      } else if slices.IndexFunc(tiles, func (x [2]int) bool {
        return x[0] == col && x[1] == row
      }) != -1 {
        fmt.Printf("O")
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

func get_input() ([]string) {
	data, _ := os.ReadFile("day20.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
	return lines
}
func get_test() ([]string) {
	data, _ := os.ReadFile("day20.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
	return lines
}

func get_cost(
  lines []string,
  pos [2]int,
  axis int,
  inc int,
) int {
  new_pos := [2]int{ pos[0], pos[1] }
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

func djikstra (
  lines []string,
  start_pos [2]int,
  end_pos [2]int,
) ([][]int) {
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

      new_pos := [2]int { col+1, row }

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

      new_pos = [2]int { col-1, row }

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

      new_pos = [2]int { col, row-1 }

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

      new_pos = [2]int { col, row+1 }

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

func getNeighbourhood(
  lines []string,
  pos [2]int,
) [][2]int {
  neighbourhood := [][2]int{}
  nrows := len(lines)
  ncols := len(lines[0])
  var new_pos [2]int

  new_pos = [2]int{ pos[0]-2, pos[1] }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0]+2, pos[1] }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0], pos[1]+2 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0], pos[1]-2 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0]-1, pos[1]-1 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0]+1, pos[1]+1 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }

  new_pos = [2]int{ pos[0]-1, pos[1]+1 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }


  new_pos = [2]int{ pos[0]+1, pos[1]-1 }

  if new_pos[0] >= 0 && new_pos[0] < ncols &&
    new_pos[1] >= 0 && new_pos[1] < nrows &&
    lines[new_pos[1]][new_pos[0]] != '#' {
    neighbourhood = append(neighbourhood, new_pos)
  }


  return neighbourhood
}

func findCheats(
  lines []string,
  start_pos [2]int,
  end_pos [2]int,
) int {
  total := 0


  costs_forward := djikstra(
    lines,
    start_pos,
    end_pos,
  )
  total_cost := costs_forward[end_pos[1]][end_pos[0]]

  costs_backward := djikstra(
    lines,
    end_pos,
    start_pos,
  )

  saves := make(map[int]int)

  for row, line := range lines {
    for col, cell := range line {
      pos := [2]int{ col, row }
      if cell != '#' {
        cost_fw := costs_forward[row][col]

        neighbours := getNeighbourhood(lines, pos)
        for _, cheat_end := range neighbours {
          cost_bw := costs_backward[cheat_end[1]][cheat_end[0]]
          save := total_cost - (cost_fw + cost_bw + 2)

          if save > 0 {
            _, ok := saves[save]
            if !ok {
              saves[save] = 0
            }
            saves[save]++
          }
        }
      }
    }
  }

  for key, value := range saves {
    //fmt.Printf("%d: %d\n", key, value)

    if key >= 100 {
      total += value
    }
  }

  return total
}

func main() {
  start := time.Now()
	lines := get_input()

  start_pos := [2]int{ -1, -1 }
  end_pos := [2]int{ -1, -1 }
  for row, line := range lines {
    for col, cell := range line {
      pos := [2]int{ col, row }
      if cell == 'S' {
        start_pos = pos
      }
      if cell == 'E' {
        end_pos = pos
      }
    }
  }

  total := findCheats(lines, start_pos, end_pos)


  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
