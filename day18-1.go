package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() ([]string) {
	data, _ := os.ReadFile("day18.input")
  lines := strings.Split(string(data), "\n")
	return lines
}
func get_test() ([]string) {
	data, _ := os.ReadFile("day18.test")
  lines := strings.Split(string(data), "\n")
	return lines
}

func main() {
	lines := get_input()

  total := 0
  fmt.Printf("total: %d\n", total)
}
