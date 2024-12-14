package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

const DEBUG = false

func main() {
	data, _ := os.ReadFile("day01.input")
	if DEBUG {
		fmt.Printf("Read %d bytes\n", len(data))
	}

	lines := strings.Split(string(data), "\n")
	if DEBUG {
		fmt.Printf("Read %d lines\n", len(lines))
	}

	var arr1, arr2 []int
	for i, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) < 2 {
			continue
		}

		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		if i < 5 && DEBUG {
			fmt.Printf("%s;;%s\n", nums[0], nums[1])
			fmt.Printf("%d;;%d\n\n", a, b)
		}

		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}

	if DEBUG {
		fmt.Printf("arr1: %d; arr2: %d\n", len(arr1), len(arr2))
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	total := 0
	for i, a := range arr1 {
		b := arr2[i]
		d := abs(a, b)
		total += d
		fmt.Printf("%d %d; %d\n", a, b, d)
	}

	fmt.Printf("\ntotal: %d\n", total)
}
