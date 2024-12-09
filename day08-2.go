package main

import (
  "fmt"
  "os"
  "strings"
  "unicode"
)

func inGrid(
  a [2]int, inc[2]int,
  i, nrows, ncols int,
) bool {
  c := [2]int {
    a[0] + inc[0]*i,
    a[1] + inc[1]*i,
  }
  return c[0] >= 0 &&
    c[0] < nrows &&
    c[1] < ncols &&
    c[1] >= 0
}

func main() {
  data, _ := os.ReadFile("day08.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  nrows := len(lines)
  ncols := len(lines[0])

  antennas := make(map[string][][2]int)


  for row, line := range lines {
    for col, c := range line {
      s := string(c)
      if unicode.IsDigit(c) || unicode.IsLetter(c) {
        antennas[s] = append(antennas[s], [2]int { row, col })
      }
    }
  }

  var antidotes [][2]int

  for _, same_type_antennas := range antennas {
    if len(same_type_antennas) < 2 {
      continue
    }

    for i, a := range same_type_antennas {
      for j, b := range same_type_antennas {
        if i == j {
          continue
        }

        inc := [2]int {
          a[0] - b[0],
          a[1] - b[1],
        }
        for i := 0; inGrid(a, inc, i, nrows, ncols); i++ {
          c := [2]int {
            a[0] + inc[0]*i,
            a[1] + inc[1]*i,
          }
          antidotes = append(antidotes, c)
        }
      }
    }
  }

  // dedup
  antidotes_map := make(map[int]map[int]bool)
  for _, a := range antidotes {
    if antidotes_map[a[0]] == nil {
      antidotes_map[a[0]] = make(map[int]bool)
    }
    antidotes_map[a[0]][a[1]] = true
  }

  total := 0
  for _, row_of_unique_antidotes := range antidotes_map {
    for _, unique_antidote := range row_of_unique_antidotes {
      if unique_antidote {
        total++
      }
    }
  }

  fmt.Printf("total: %d\n", total)
}
