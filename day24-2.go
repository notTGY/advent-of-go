package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func get_input() ([]string, [][]string) {
	data, _ := os.ReadFile("day24.input")
	blocks := strings.Split(string(data), "\n\n")
	inputs := strings.Split(blocks[0], "\n")
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]

	gates := [][]string{}
	for _, line := range lines {
		split_a := strings.Split(line, " -> ")
		split_b := strings.Split(split_a[0], " ")
		gate := []string{
			split_b[0], split_b[1], split_b[2], split_a[1],
		}
		gates = append(gates, gate)
	}

	return inputs, gates
}
func get_test() ([]string, [][]string) {
	data, _ := os.ReadFile("day24.test")
	blocks := strings.Split(string(data), "\n\n")
	inputs := strings.Split(blocks[0], "\n")
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]

	gates := [][]string{}
	for _, line := range lines {
		split_a := strings.Split(line, " -> ")
		split_b := strings.Split(split_a[0], " ")
		gate := []string{
			split_b[0], split_b[1], split_b[2], split_a[1],
		}
		gates = append(gates, gate)
	}

	return inputs, gates
}

func find(a, b, op string, gates [][]string) string {
	for _, gate := range gates {
		if gate[0] == a && gate[1] == op && gate[2] == b {
			return gate[3]
		}
		if gate[0] == b && gate[1] == op && gate[2] == a {
			return gate[3]
		}
	}
	return ""
}

func main() {
	start := time.Now()
	_, gates := get_input()

	xs := []string{}
	ys := []string{}
	zs := []string{}
	for _, gate := range gates {
		if strings.HasPrefix(gate[0], "x") &&
			slices.Index(xs, gate[0]) == -1 {
			xs = append(xs, gate[0])
		}
		if strings.HasPrefix(gate[2], "x") &&
			slices.Index(xs, gate[2]) == -1 {
			xs = append(xs, gate[2])
		}
		if strings.HasPrefix(gate[2], "y") &&
			slices.Index(ys, gate[2]) == -1 {
			ys = append(ys, gate[2])
		}
		if strings.HasPrefix(gate[0], "y") &&
			slices.Index(ys, gate[0]) == -1 {
			ys = append(ys, gate[0])
		}
		if strings.HasPrefix(gate[3], "z") {
			zs = append(zs, gate[3])
		}
	}
	slices.Sort(xs)
	slices.Sort(ys)
	slices.Sort(zs)

	var swapped []string
	var c0 string

	for i, x := range xs {
		y := ys[i]

		// I DO NOT UNDERSTAND THIS===
		var m1, n1, r1, z1, c1 string

		// Half adder logic
		m1 = find(x, y, "XOR", gates)
		n1 = find(x, y, "AND", gates)

		if c0 != "" {
			r1 = find(c0, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(c0, m1, "AND", gates)
			}

			z1 = find(c0, m1, "XOR", gates)

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}

			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}

			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gates)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}
	// END===

	slices.Sort(swapped)
	fmt.Printf("res: %s\n", strings.Join(swapped, ","))

	fmt.Printf("took: %s;\n", time.Since(start))
}
