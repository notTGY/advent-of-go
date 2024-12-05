package main

import (
  "fmt"
  "os"
  "strings"
)


func main() {
  data, _ := os.ReadFile("day6.test")
  parts := strings.Split(string(data), "\n\n")
  rules := strings.Split(parts[0], "\n")
  updates := strings.Split(parts[1], "\n")

  // extra whitespace in the end of file
  updates = updates[:len(updates)-1]

  total := 0
  for _, s := range updates {
    total += isCorrect(s, rules)
  }

  fmt.Printf("total: %d\n", total)
}
