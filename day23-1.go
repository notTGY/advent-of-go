package main

import (
	"fmt"
	"os"
	"strings"
  "strconv"
	"time"
)

func get_input() []string {
	data, _ := os.ReadFile("day23.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
  return lines
}
func get_test() []string {
	data, _ := os.ReadFile("day23.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
  return lines
}

func main() {
	start := time.Now()
	lines := get_test()

	total := 0
  for _, _ = range lines {
    total++
  }

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
