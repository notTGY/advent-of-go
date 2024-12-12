package main

import (
  "fmt"
  "os"
  "strings"
  "slices"
)

func perimeter(plot [][2]int) int {
  diff := 0
  for _, p := range plot {
    candidates := [][2]int {
      [2]int { p[0]+1, p[1] },
      [2]int { p[0]-1, p[1] },
      [2]int { p[0], p[1]+1 },
      [2]int { p[0], p[1]-1 },
    }

    for _, candidate := range candidates {
      if slices.Index(plot, candidate) != -1 {
        diff += 1
      }
    }
  }

  per := 4 * len(plot) - diff

  return per
}

func safeGetV(
  m[]string,
  row, col, nrows, ncols int,
) byte {
  if row >= nrows ||
    row < 0 ||
    col < 0 ||
    col >= ncols {
    return 0
  }
  return m[row][col]
}
func search(
  lines []string,
  row, col int,
  res[][2]int,
) [][2]int {
  nrows := len(lines)
  ncols := len(lines[0])

  letter := safeGetV(lines, row, col, nrows, ncols)
  //fmt.Printf("%d %d: %s\n", row, col, string(letter))

  if slices.Index(res, [2]int{ row+1, col }) == -1 &&
    safeGetV(
      lines, row+1, col, nrows, ncols,
    ) == letter {
    res = append(res, [2]int { row+1, col })
    new_vs := search(lines, row+1, col, res)
    for _, v := range new_vs {
      if slices.Index(res, v) == -1 {
        res = append(res, v)
      }
    }
  }
  if slices.Index(res, [2]int{ row-1, col }) == -1 &&
    safeGetV(
      lines, row-1, col, nrows, ncols,
    ) == letter {
    res = append(res, [2]int { row-1, col })
    new_vs := search(lines, row-1, col, res)
    for _, v := range new_vs {
      if slices.Index(res, v) == -1 {
        res = append(res, v)
      }
    }
  }
  if slices.Index(res, [2]int{ row, col-1 }) == -1 &&
    safeGetV(
      lines, row, col-1, nrows, ncols,
    ) == letter {
    res = append(res, [2]int { row, col-1 })
    new_vs := search(lines, row, col-1, res)
    for _, v := range new_vs {
      if slices.Index(res, v) == -1 {
        res = append(res, v)
      }
    }
  }
  if slices.Index(res, [2]int{ row, col+1 }) == -1 &&
    safeGetV(
      lines, row, col+1, nrows, ncols,
    ) == letter {
    res = append(res, [2]int { row, col+1 })
    new_vs := search(lines, row, col+1, res)
    for _, v := range new_vs {
      if slices.Index(res, v) == -1 {
        res = append(res, v)
      }
    }
  }

  return res
}

func main() {
  data, _ := os.ReadFile("day12.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
  ncols := len(lines[0])

  var used [][]bool
  for row, line := range lines {
    used = append(used, make([]bool, ncols))
    for col, _ := range line {
      used[row][col] = false
    }
  }

  total := 0

  fmt.Printf("Search started\n")
  for slices.IndexFunc(used, func (m []bool) bool {
    idx := slices.Index(m, false)
    return idx != -1
  }) != -1 {
    var col_start int
    row_start := slices.IndexFunc(
      used,
      func (m []bool) bool {
        col_start = slices.Index(m, false)
        return col_start != -1
      },
    )

    plot := search(
      lines,
      row_start,
      col_start,
      [][2]int{ [2]int { row_start, col_start } },
    )
    for _, p := range plot {
      used[p[0]][p[1]] = true
    }
    per := perimeter(plot)
    /*
    fmt.Printf(
      "size: %d; per: %d\n",
      len(plot),
      per,
    )
    */

    total += len(plot)*per
  }

  fmt.Printf("total: %d\n", total)
}

