package main

import (
	"fmt"
	"os"
	"strings"
  "time"
)

func get_input() ([]string, []string) {
	data, _ := os.ReadFile("day19.input")
  blocks := strings.Split(string(data), "\n\n")
  towels := strings.Split(blocks[0], ", ")
  lines := strings.Split(blocks[1], "\n")
  lines = lines[:len(lines)-1]
	return towels, lines
}
func get_test() ([]string, []string) {
	data, _ := os.ReadFile("day19.test")
  blocks := strings.Split(string(data), "\n\n")
  towels := strings.Split(blocks[0], ", ")
  lines := strings.Split(blocks[1], "\n")
  lines = lines[:len(lines)-1]
	return towels, lines
}

func isPossible(line string, towels[]string) bool {
  possible := false
  for _, towel := range towels {
    if towel == line {
      return true
    }
    after, found := strings.CutPrefix(line, towel)
    if !found {
      continue
    }

    possible = isPossible(after, towels)
    if possible {
      break
    }
  }
  return possible
}

func main() {
  start := time.Now()
	towels, lines := get_input()

  total := 0
  for _, line := range lines {
    if isPossible(line, towels) {
      total++
    }
  }

  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
