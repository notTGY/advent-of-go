package main

import (
  "fmt"
  "os"
  "strings"
  "math"
  "time"
)

func powInt(x, y int) int {
  return int(math.Pow(float64(x), float64(y)))
}

func appendUnique(
  a [][6]int, x [2]int, d int,
) ([][6]int, int) {
  for i, b := range a {
    if b[0] == x[0] &&
      b[1] == x[1] {

      idx := (d%4) + 1
      a[i][idx] = b[idx] + 1

      if a[i][idx] > 3 {
        return a, 1000
      }

      return a, a[i][2]
    }
  }
  var new_val [6]int
  new_val[0] = x[0]
  new_val[1] = x[1]

  new_val[2] = 0
  new_val[3] = 0
  new_val[4] = 0
  new_val[5] = 0

  idx := (d%4) + 1
  new_val[idx] = new_val[idx] + 1

  return append(a, new_val), 1
}

func isObstructed(
  obstacles [][2]int, guard [2]int,
) (bool) {
  for _, obstacle := range obstacles {
    if obstacle[0] == guard[0] &&
      obstacle[1] == guard[1] {
        return true
    }
  }
  return false
}

func getDir(d int) ([2]int) {
  d = d % 4
  if d == 0 {
    return [2]int { -1, 0 }
  }
  if d == 1 {
    return [2]int { 0, 1 }
  }
  if d == 2 {
    return [2]int { 1, 0 }
  }
  return [2]int { 0, -1 }
}


func appendUnique2(
  obs [][2]int, ob [2]int,
) ([][2]int) {
  for _, i := range obs {
    if i[0] == ob[0] &&
      i[1] == ob[1] {
      return obs
    }
  }
  return append(obs, ob)
}

func explore(
  gc [2]int,
  ob [2]int,
  obstacles [][2]int,
  num_rows int,
  num_cols int,
) (int, [][6]int) {
  finished := false
  direction := 0

  var history [][6]int
  history, _ = appendUnique(history, gc, direction)
  steps := 0
  for !finished {
    steps++
    dir := getDir(direction)
    next_gc := [2]int {
      gc[0] + dir[0],
      gc[1] + dir[1],
    }

    n := 1
    for isObstructed(obstacles, next_gc) ||
      isObstructed([][2]int {ob}, next_gc) {
      direction++
      dir = getDir(direction)
      next_gc = [2]int {
        gc[0] + dir[0],
        gc[1] + dir[1],
      }

      n++
      if n > 4 {
        return -1, history
      }
    }

    gc = next_gc

    if gc[0] < 0 ||
      gc[0] >= num_rows ||
      gc[1] < 0 ||
      gc[1] >= num_cols {
        finished = true
    }
    h, delve := appendUnique(history, gc, direction)
    history = h
    if delve == 1000 {
      return -1, history
    }

    /*
    if steps > num_cols*num_rows*6 {
      return -1, history
    }
    */
  }


  return len(history) - 1, history
}

type V [2]int

func (g V) Copy() ([2]int) {
  return [2]int{ g[0], g[1] }
}

func main() {
  start := time.Now()
  data, _ := os.ReadFile("day6.test")
  lines := strings.Split(string(data), "\n")
  // extra whitespace in the end of file
  lines = lines[:len(lines)-1]


  num_cols := len(lines[0])
  num_rows := len(lines)

  var obstacles [][2]int
  var guard V

  for row, line := range lines {
    for col, cell := range line {
      if string(cell) == "#" {
        obstacles = append(
          obstacles, [2]int { row, col },
        )
      }
      if string(cell) == "^" {
        guard = V { row, col }
      }
    }
  }


  total := 0
  gc := guard.Copy()
  _, h := explore(
    gc,
    obstacles[0],
    obstacles,
    num_rows,
    num_cols,
  )
  for i, ob2 := range h {
    if i == 0 {
      continue
    }
    ob := [2]int { ob2[0], ob2[1] }
    gc = guard.Copy()
    delve, _ := explore(
      gc, ob, obstacles, num_rows, num_cols,
    )
    if delve == -1 {
      total++
    }
  }

  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}




/// Binary search

// 1733 -- too low

// 1912 -- too high
