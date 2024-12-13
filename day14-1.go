package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  data, _ := os.ReadFile("day14.test")
  lines := strings.Split(string(data), "\n")

  total := 0

  fmt.Printf("total: %d\n", total)
}
