package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

type Obstacle [2]int
type Pos [2]int

func getInputs(
  path string,
) (int, int, []Obstacle, Pos) {
  data, _ := os.ReadFile(path)
  lines := strings.Split(string(data), "\n")
  // extra whitespace in the end of file
  lines = lines[:len(lines)-1]

  ncols := len(lines[0])
  nrows := len(lines)

  var obstacles []Obstacle
  var guard Pos
  for row, line := range lines {
    for col, cell := range line {
      switch string(cell) {
        case "#":
          obstacles = append(
            obstacles, Obstacle { row, col },
          )
        case "^":
          guard = Pos { row, col }
      }
    }
  }

  return ncols, nrows, obstacles, guard
}

func getDir(direction int) Pos {
  switch direction % 4 {
    case 0:
      return Pos { -1, 0 }
    case 1:
      return Pos { 0, 1 }
    case 2:
      return Pos { 1, 0 }
    default:
      return Pos { 0, -1 }
  }
}

func appendUnique(history []Pos, guard Pos) []Pos {
  for _, p := range history {
    if p[0] == guard[0] && p[1] == guard[1] {
      return history
    }
  }
  return append(history, guard)
}

func isObstructed(
  obstacles []Obstacle, head Pos,
) (bool) {
  for _, o := range obstacles {
    if o[0] == head[0] && o[1] == head[1] {
      return true
    }
  }
  return false
}

func getReachable(
  ncols, nrows int, obstacles []Obstacle, guard Pos,
) []Pos {
  finished := false
  direction := 0

  var history []Pos
  history = append(history, guard)
  for !finished {
    dir := getDir(direction)
    head := Pos {
      guard[0] + dir[0], guard[1] + dir[1],
    }

    for isObstructed(obstacles, head) {
      direction++
      dir = getDir(direction)
      head = Pos {
        guard[0] + dir[0], guard[1] + dir[1],
      }
    }
    guard = head

    if guard[0] < 0 ||
      guard[0] >= nrows ||
      guard[1] < 0 ||
      guard[1] >= ncols {
        finished = true
    } else {
      history = appendUnique(history, guard)
    }
  }

  return history
}

func main() {
  start := time.Now()
  ncols, nrows, obstacles, guard := getInputs("day06.input")
  total := len(getReachable(ncols, nrows, obstacles, guard))
  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
