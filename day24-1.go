package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func get_input() ([]string, []string) {
	data, _ := os.ReadFile("day24.input")
	blocks := strings.Split(string(data), "\n\n")
	inputs := strings.Split(blocks[0], "\n")
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]
	return inputs, lines
}
func get_test() ([]string, []string) {
	data, _ := os.ReadFile("day24.test")
	blocks := strings.Split(string(data), "\n\n")
	inputs := strings.Split(blocks[0], "\n")
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]
	return inputs, lines
}

func main() {
	start := time.Now()
	inputs, lines := get_input()

	wire_values := make(map[string]bool)
	for _, input := range inputs {
		splitted := strings.Split(input, ": ")
		val := true
		if splitted[1] == "0" {
			val = false
		}
		wire_values[splitted[0]] = val
	}

	zs := []string{}
	gates := [][]string{}
	for _, line := range lines {
		split_a := strings.Split(line, " -> ")
		split_b := strings.Split(split_a[0], " ")
		gate := []string{
			split_b[0], split_b[1], split_b[2], split_a[1],
		}
		gates = append(gates, gate)
		if strings.HasPrefix(split_a[1], "z") {
			zs = append(zs, split_a[1])
		}
	}
	slices.Sort(zs)

	for {
		gottem := true
		for _, z := range zs {
			_, ok := wire_values[z]
			if !ok {
				gottem = false
			}
		}
		if gottem {
			break
		}

		for _, gate := range gates {
			a := gate[0]
			b := gate[2]
			c := gate[3]
			val_a, ok := wire_values[a]
			if !ok {
				continue
			}
			val_b, ok := wire_values[b]
			if !ok {
				continue
			}
			val := true
			switch gate[1] {
			case "XOR":
				val = (val_a || val_b) && !(val_a && val_b)
			case "OR":
				val = val_a || val_b
			case "AND":
				val = val_a && val_b
			}
			wire_values[c] = val
		}
	}

	total := 0
	pow := 1
	for _, z := range zs {
		vb, _ := wire_values[z]
		v := 0
		if vb {
			v = 1
		}
		total = total + v*pow
		pow *= 2
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
