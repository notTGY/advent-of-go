package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
  "time"
)

func get_input() ([]string) {
	data, _ := os.ReadFile("day19.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
	return lines
}
func get_test() ([]string) {
	data, _ := os.ReadFile("day19.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
	return lines
}

func main() {
  start := time.Now()
	lines := get_input()

  total := 0

  fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
