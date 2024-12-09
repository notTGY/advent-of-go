package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

type Pos [2]int
type Obstacle struct {
  pos Pos
}

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
            obstacles, Obstacle { pos: Pos {row, col} },
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
    if o.pos[0] == head[0] && o.pos[1] == head[1] {
      return true
    }
  }
  return false
}

func makeStep(
  obstacles []Obstacle, head Pos, direction int,
) Pos {
  new_head := Pos { -1, -1 }

  for _, ob := range obstacles {
    switch direction % 4 {
      case 0:
        if ob.pos[0] < head[0] &&
          ob.pos[1] == head[1] &&
          (ob.pos[0] >= new_head[0]) {
          new_head[0] = ob.pos[0] + 1
          new_head[1] = head[1]
        }
      case 2:
        if ob.pos[0] > head[0] &&
          ob.pos[1] == head[1] &&
          (ob.pos[0] <= new_head[0] || new_head[0] == -1) {
          new_head[0] = ob.pos[0] - 1
          new_head[1] = head[1]
        }
      case 1:
        if ob.pos[1] > head[1] &&
          ob.pos[0] == head[0] &&
          (ob.pos[1] <= new_head[1] || new_head[1] == -1) {
          new_head[0] = head[0]
          new_head[1] = ob.pos[1] - 1
        }
      case 3:
        if ob.pos[1] < head[1] &&
          ob.pos[0] == head[0] &&
          (ob.pos[1] >= new_head[1] || new_head[1] == -1) {
          new_head[0] = head[0]
          new_head[1] = ob.pos[1] + 1
        }
    }
  }

  return new_head
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

func isLoop(
  ncols, nrows int,
  obstacles []Obstacle,
  guard Pos,
  ob Obstacle,
) int {
  finished := false
  direction := 0

  full_obstacles := append(obstacles, ob)

  step := 0
  for !finished {
    new_head := makeStep(
      full_obstacles, guard, direction,
    )
    if new_head[0] == -1 && new_head[1] == -1 {
      return 0
    }
    guard = Pos { new_head[0], new_head[1] }
    direction = (direction+1)%4

    step++
    if step > nrows*ncols*2 {
      return 1
    }
  }

  return 0
}

func main() {
  start := time.Now()
  ncols, nrows, obstacles, guard := getInputs("day06.input")
  reachable := getReachable(ncols, nrows, obstacles, guard)

  total := 0
  for i, p := range reachable {
    if i == 0 {
      continue
    }

    inc := isLoop(
      ncols,
      nrows,
      obstacles,
      guard,
      Obstacle { pos: Pos {p[0], p[1]} },
    )
    //if inc == 1 {
      //fmt.Printf("Got %v\n", p)
    //}
    total += inc
  }

  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
