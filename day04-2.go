package main

import (
  "fmt"
  "os"
  "strings"
)

func testDir(
  s []string,
  i int,
  j int,
  dir [2]int,
) (int) {
  Si := i - dir[0]
  Sj := j - dir[1]

  if Si < 0 || Si > len(s)-1 {
    return 0
  }
  if Sj < 0 || Sj > len(s[i])-1 {
    return 0
  }

  if s[Si][Sj] != "S"[0] {
    return 0
  }

  return 1
}

func findXmas(s []string, i int, j int) (int) {
  total := 0

  dirs := findDir(s, i, j, "M"[0])
  if len(dirs) != 2 {
    return 0
  }
  for _, dir := range dirs {
    r := testDir(s, i, j, dir)
    total += r
    /*
    if r > 0 {
      fmt.Printf("offset: %d %d\n", dir[0], dir[1])
    }
    */
  }
  //fmt.Printf("\n")
  if total == 2 {
    return 1
  }
  return 0
}

func findDir(
  s []string,
  i int,
  j int,
  search byte,
) ([][2]int) {
  var res [][2]int

  if i > 0 {
    if j > 0 {
      if s[i-1][j-1] == search {
        var pos  = [2]int{-1, -1}
        res = append(res, pos)
      }
    }
    if j+1 < len(s[i]) {
      if s[i-1][j+1] == search {
        var pos  = [2]int{-1, 1}
        res = append(res, pos)
      }
    }
    /*
    if s[i-1][j] == search {
      var pos  = [2]int{-1, 0}
      res = append(res, pos)
    }
    */
  }

  if i+1 < len(s) {
    if j > 0 {
      if s[i+1][j-1] == search {
        var pos  = [2]int{1, -1}
        res = append(res, pos)
      }
    }
    if j+1 < len(s[i]) {
      if s[i+1][j+1] == search {
        var pos  = [2]int{1, 1}
        res = append(res, pos)
      }
    }
    /*
    if s[i+1][j] == search {
      var pos  = [2]int{1, 0}
      res = append(res, pos)
    }
    */
  }

  /*
  if j > 0 {
    if s[i][j-1] == search {
      var pos  = [2]int{0, -1}
      res = append(res, pos)
    }
  }
  if j+1 < len(s[i]) {
    if s[i][j+1] == search {
      var pos  = [2]int{0, 1}
      res = append(res, pos)
    }
  }
  */
  return res
}

func main() {
  data, _ := os.ReadFile("day04.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  //fmt.Printf("Len s: %d, Len s[i]: %d\n", len(lines), len(lines[0]))

  total := 0
  for i, s := range lines {
    //fmt.Printf("Line: %s\n", s)
    j := 0
    found_first := false
    for strings.Index(s, "A") != -1 {
      local_j := strings.Index(s, "A")
      s = s[local_j+1:]

      if found_first {
        j++
      }
      found_first = true
      j += local_j
      //fmt.Printf("i: %d, j: %d; l: %d\n", i, j, local_j)

      total += findXmas(lines, i, j)
    }
  }

  fmt.Printf("total: %d\n", total)
}
