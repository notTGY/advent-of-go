package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  data, _ := os.ReadFile("day10.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  total := 0

  fmt.Printf("total: %d\n", total)
}
