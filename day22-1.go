package main

import (
	"fmt"
	"os"
	"strings"
  "strconv"
	"time"
)

func get_input() []int {
	data, _ := os.ReadFile("day22.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
  numbers := []int{}
  for _, line := range lines {
    x, _ := strconv.Atoi(line)
    numbers = append(numbers, x)
  }
	return numbers
}
func get_test() []int {
	data, _ := os.ReadFile("day22.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
  numbers := []int{}
  for _, line := range lines {
    x, _ := strconv.Atoi(line)
    numbers = append(numbers, x)
  }
	return numbers
}


func evolute(secret int, times int) int {
  for i := 0; i < times; i++ {
    secret = (secret ^ (secret * 64)) % 16777216
    secret = (secret ^ (secret / 32)) % 16777216
    secret = (secret ^ (secret * 2048)) % 16777216
  }
  return secret
}

func main() {
	start := time.Now()
	secrets := get_input()

	total := 0
  for _, secret := range secrets {
    evoluted := evolute(secret, 2000)
    total += evoluted
  }

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
