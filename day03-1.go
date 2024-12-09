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

func main() {
  data, _ := os.ReadFile("day03.test")

  s := string(data)

  total := 0
  for strings.Index(s, "mul(") != -1 {
    //fmt.Printf("(: %s\n", s)
    j := strings.Index(s, "mul(")
    s = string(s[j+4:])

    //fmt.Printf(",: %s\n", s)
    j = strings.Index(s, ",")
    // 3-max int, 4-len("mul(")
    if j > 3+4 {
      continue
    }
    aStr := string(s[:j])
    //fmt.Printf("aStr: %s\n", aStr)
    if len(aStr) > len(strings.TrimSpace(aStr)) {
      continue
    }
    a, err := strconv.Atoi(aStr)
    if err != nil {
      continue
    }
    s = string(s[j+1:])

    //fmt.Printf("): %s\n", s)
    j = strings.Index(s, ")")
    bStr := string(s[:j])
    //fmt.Printf("bStr: %s\n", bStr)
    if len(bStr) > len(strings.TrimSpace(bStr)) {
      continue
    }
    if len(bStr) > 3 {
      continue
    }
    b, err := strconv.Atoi(bStr)
    if err != nil {
      continue
    }
    s = string(s[j:])
    //fmt.Printf("a: %d; b: %d\n", a, b)
    total += a*b
  }

  fmt.Printf("\ntotal: %d\n", total)
}
