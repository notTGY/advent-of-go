package main

import (
  "fmt"
  "os"
  "log"
  "strings"
  "strconv"
)

func canAssembe(r int, nums []int) bool {
  last := nums[len(nums)-1]
  if len(nums) == 1 {
    return last == r
  }
  rest := nums[:len(nums)-1]

  if r % last == 0 &&
    canAssembe(r / last, rest) {
      return true
  }

  diff := r - last
  same := true
  for l := len(strconv.Itoa(last)); l > 0 && same; l-- {
    if diff % 10 != 0 {
      same = false
    }
    diff = diff / 10
  }
  if same && canAssembe(diff, rest) {
    return true
  }

  return canAssembe(r - last, rest)
}

func main() {
  data, _ := os.ReadFile("day7.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  total := 0
  for _, line := range lines {
    splitByColon := strings.Split(line, ": ")
    if len(splitByColon) != 2 {
      log.Fatal("HUH")
    }
    r, _ := strconv.Atoi(splitByColon[0])
    splitBySpace := strings.Split(
      string(splitByColon[1]), " ",
    )

    var nums []int
    for _, s := range splitBySpace {
      num, _ := strconv.Atoi(s)
      nums = append(nums, num)
    }

    if canAssembe(r, nums) {
      total += r
    }
  }

  fmt.Printf("total: %d\n", total)
}
