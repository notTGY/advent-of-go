package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  data, _ := os.ReadFile("day10.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  nrows := len(lines)
  ncols := len(lines[0])


  total := 0

  fmt.Printf("total: %d\n", total)
}

