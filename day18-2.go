package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
  "slices"
  "time"
)

func get_input() ([][2]int, int, int, int) {
	data, _ := os.ReadFile("day18.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
  positions := [][2]int{}
  for _, line := range lines {
    splitted := strings.Split(line, ",")
    col, _ := strconv.Atoi(splitted[0])
    row, _ := strconv.Atoi(splitted[1])
    positions = append(positions, [2]int{col, row})
  }
  nrows := 71
  ncols := 71
  nsteps := 1024
	return positions, nrows, ncols, nsteps
}
func get_test() ([][2]int, int, int, int) {
	data, _ := os.ReadFile("day18.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
  positions := [][2]int{}
  for _, line := range lines {
    splitted := strings.Split(line, ",")
    col, _ := strconv.Atoi(splitted[0])
    row, _ := strconv.Atoi(splitted[1])
    positions = append(positions, [2]int{col, row})
  }
  nrows := 7
  ncols := 7
  nsteps := 12
	return positions, nrows, ncols, nsteps
}

func print_costs(
  obstacles [][2]int,
  costs[][]int,
) {
  for row, l := range costs {
    for col, c := range l {
      if c == -1 && slices.IndexFunc(obstacles, func(x [2]int) bool {
    return x[0] == col && x[1] == row
  }) != -1 {
        fmt.Printf(" # ")
        continue
      }
      fmt.Printf("%2d ", c)
    }
    fmt.Printf("\n")
  }
}

func djikstra (
  obstacles [][2]int,
  start_pos [2]int,
  end_pos [2]int,
  nrows int,
  ncols int,
) (int) {
  has_obstacle := make(map[int]bool)
  for _, o := range obstacles {
    has_obstacle[o[1]*ncols+o[0]] = true
  }

  positions := [][2]int{}
  visited := make(map[int]bool)
  visited[start_pos[1]*ncols+start_pos[0]] = true

  positions = append(positions, start_pos)

  cur_cost := 0
  for {
    /*
    fmt.Printf("%v\n", positions)
    for _, pos := range positions {
      fmt.Printf(
        "%d %d\n",
        pos[0],
        pos[1],
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

      if col == end_pos[0] && row == end_pos[1] {
        return cur_cost
      }

      var new_diff_cost int
      var v bool
      var ok bool
      var index int

      new_pos := [2]int { col+1, row }

      index = new_pos[1]*ncols+new_pos[0]
      new_diff_cost = 1
      if new_pos[0] < 0 || new_pos[0] >= ncols {
        new_diff_cost = -1
      }
      if new_pos[1] < 0 || new_pos[1] >= nrows {
        new_diff_cost = -1
      }
      v, ok = has_obstacle[index]
      if new_diff_cost != -1 && ok && v {
        new_diff_cost = -1
      }
      v, ok = visited[index]
      if new_diff_cost != -1 && !(ok && v) {
        visited[index] = true
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col-1, row }

      index = new_pos[1]*ncols+new_pos[0]
      new_diff_cost = 1
      if new_pos[0] < 0 || new_pos[0] >= ncols {
        new_diff_cost = -1
      }
      if new_pos[1] < 0 || new_pos[1] >= nrows {
        new_diff_cost = -1
      }
      v, ok = has_obstacle[index]
      if new_diff_cost != -1 && ok && v {
        new_diff_cost = -1
      }
      v, ok = visited[index]
      if new_diff_cost != -1 && !(ok && v) {
        visited[index] = true
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col, row-1 }

      index = new_pos[1]*ncols+new_pos[0]
      new_diff_cost = 1
      if new_pos[0] < 0 || new_pos[0] >= ncols {
        new_diff_cost = -1
      }
      if new_pos[1] < 0 || new_pos[1] >= nrows {
        new_diff_cost = -1
      }
      v, ok = has_obstacle[index]
      if new_diff_cost != -1 && ok && v {
        new_diff_cost = -1
      }
      v, ok = visited[index]
      if new_diff_cost != -1 && !(ok && v) {
        visited[index] = true
        new_positions = append(
          new_positions, new_pos,
        )
      }

      new_pos = [2]int { col, row+1 }

      index = new_pos[1]*ncols+new_pos[0]
      new_diff_cost = 1
      if new_pos[0] < 0 || new_pos[0] >= ncols {
        new_diff_cost = -1
      }
      if new_pos[1] < 0 || new_pos[1] >= nrows {
        new_diff_cost = -1
      }
      v, ok = has_obstacle[index]
      if new_diff_cost != -1 && ok && v {
        new_diff_cost = -1
      }
      v, ok = visited[index]
      if new_diff_cost != -1 && !(ok && v) {
        visited[index] = true
        new_positions = append(
          new_positions, new_pos,
        )
      }

    }
    cur_cost++
    positions = new_positions
  }

  return -1
}

func main() {
  start := time.Now()
	obstacles, nrows, ncols, nsteps := get_input()

  for step := nsteps; step <= len(obstacles); step++ {
    cost := djikstra(
      obstacles[:step],
      [2]int{0, 0},
      [2]int{ncols-1, nrows-1},
      nrows,
      ncols,
    )
    if cost == -1 {
      fmt.Printf(
        "took: %s; result: %d,%d\n",
        time.Since(start),
        obstacles[step-1][0],
        obstacles[step-1][1],
      )
      break
    }
  }
}
