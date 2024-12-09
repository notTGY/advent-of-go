package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func appendUnique(a [][2]int, x [2]int) ([][2]int) {
  for _, b := range a {
    if b[0] == x[0] &&
      b[1] == x[1] {
      return a
    }
  }
  return append(a, x)
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

func main() {
  start := time.Now()
  data, _ := os.ReadFile("day06.input")
  lines := strings.Split(string(data), "\n")
  // extra whitespace in the end of file
  lines = lines[:len(lines)-1]


  num_cols := len(lines[0])
  num_rows := len(lines)

  var obstacles [][2]int

  var guard [2]int

  for row, line := range lines {
    for col, cell := range line {
      if string(cell) == "#" {
        obstacles = append(
          obstacles, [2]int { row, col },
        )
      }
      if string(cell) == "^" {
        guard = [2]int { row, col }
      }
    }
  }

  finished := false
  direction := 0

  var history [][2]int
  history = append(history, guard)
  for !finished {
    dir := getDir(direction)
    guard[0] += dir[0]
    guard[1] += dir[1]

    for isObstructed(obstacles, guard) {
      guard[0] -= dir[0]
      guard[1] -= dir[1]
      direction++
      dir = getDir(direction)
      guard[0] += dir[0]
      guard[1] += dir[1]
    }

    if guard[0] < 0 ||
      guard[0] >= num_rows ||
      guard[1] < 0 ||
      guard[1] >= num_cols {
        finished = true
    }
    history = appendUnique(history, guard)
  }

  total := len(history) - 1
  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
