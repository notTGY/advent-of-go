package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func abs(x, y int) (int) {
  if x > y {
    return x - y
  }
  return y - x
}


func isSafe(l []int) (int) {
  if isSingleSafe(l, -1) == 1 {
    return 1
  }

  for j, _ := range l {
    if isSingleSafe(l, j) == 1 {
      return 1
    }
  }

  return 0
}

func isSingleSafe(levels []int, j int) (int) {
  var isIncreasing bool
  dir := false
  for i, x := range levels {
    if i == j {
      continue
    }
    if i == 0 {
      continue
    }
    if i == 1 && j == 0 {
      continue
    }
    //fmt.Printf("%d %d \n", j, i)

    pidx := i-1
    if pidx == j {
      pidx--
    }

    p := levels[pidx]

    if abs(x, p) > 3 || abs(x, p) < 1 {
      return 0
    }
    if !dir {
      isIncreasing = x > p
      dir = true
    } else {
      if isIncreasing != (x > p) {
        return 0
      }
    }
  }
  return 1
}

const DEBUG = false

func main() {
  data, _ := os.ReadFile("day02.input")
  if DEBUG {
    fmt.Printf("Read %d bytes\n", len(data))
  }

  lines := strings.Split(string(data), "\n")
  if DEBUG {
    fmt.Printf("Read %d lines\n", len(lines))
  }

  total := 0
  for i, line := range(lines) {
    if i == len(lines) - 1 {
      continue
    }
    numStrs := strings.Split(line, " ")
    var nums []int
    for _, s := range numStrs {
      x, _ := strconv.Atoi(s)
      nums = append(nums, x)

      //fmt.Printf("%d ", x)
    }
    is_safe := isSafe(nums)
    //fmt.Printf("; %d\n", is_safe)
    total += is_safe
  }

  fmt.Printf("\ntotal: %d\n", total)
}
