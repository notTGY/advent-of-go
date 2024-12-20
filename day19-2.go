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

func countPossible(
	line string,
	towels []string,
	cache *map[string]int,
) int {
	v, ok := (*cache)[line]
	if ok {
		return v
	}

	possible := 0
	for _, towel := range towels {
		if towel == line {
			possible += 1
			continue
		}
		after, found := strings.CutPrefix(line, towel)
		if found {
			possible += countPossible(after, towels, cache)
		}
	}
	(*cache)[line] = possible

	return possible
}

func main() {
	start := time.Now()
	towels, lines := get_input()

	total := 0
	cache := make(map[string]int)
	for _, line := range lines {
		total += countPossible(line, towels, &cache)
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
