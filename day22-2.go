package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

const M = 1048576

func evolute(
	secret int, times int,
	rewards map[int]int,
) int {
	key := 0

	bananas := secret % 10
	for i := 0; i < times; i++ {
		secret = (secret ^ (secret * 64)) % 16777216
		secret = (secret ^ (secret / 32)) % 16777216
		secret = (secret ^ (secret * 2048)) % 16777216

		new_bananas := secret % 10
		key = ((key * 32) + 9 + new_bananas - bananas) % M
		bananas = new_bananas
		if i >= 3 {
			_, exists := rewards[key]
			if !exists {
				rewards[key] = bananas
			}
		}
	}
	return secret
}

func main() {
	start := time.Now()
	secrets := get_input()

	rewards := []map[int]int{}
	allCombinations := []int{}
	for i, secret := range secrets {
		rewards = append(rewards, make(map[int]int))
		_ = evolute(secret, 2000, rewards[i])
		for k, _ := range rewards[i] {
			if slices.Index(allCombinations, k) == -1 {
				allCombinations = append(allCombinations, k)
			}
		}
	}
	fmt.Printf("Got all combinations in %s\n", time.Since(start))
	total := 0
	for _, k := range allCombinations {
		sum := 0
		for i, _ := range secrets {
			sum += rewards[i][k]
		}
		if sum > total {
			total = sum
		}
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
