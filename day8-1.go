package main

import (
  "fmt"
  "os"
  "log"
  "strings"
  "strconv"
)

func main() {
  data, _ := os.ReadFile("day8.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  total := 0

  fmt.Printf("total: %d\n", total)
}
