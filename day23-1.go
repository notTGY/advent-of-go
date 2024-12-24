package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func get_input() [][]string {
	data, _ := os.ReadFile("day23.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	pairs := [][]string{}
	for _, line := range lines {
		pair := strings.Split(line, "-")
		if pair[0] > pair[1] {
			tmp := pair[0]
			pair[0] = pair[1]
			pair[1] = tmp
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
func get_test() [][]string {
	data, _ := os.ReadFile("day23.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	pairs := [][]string{}
	for _, line := range lines {
		pair := strings.Split(line, "-")
		if pair[0] > pair[1] {
			tmp := pair[0]
			pair[0] = pair[1]
			pair[1] = tmp
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

func main() {
	start := time.Now()
	pairs := get_input()

	i := 0
	computer_ids := make(map[string]int)
	computers := make(map[int]string)
	connections := make(map[int][]int)
	for _, pair := range pairs {
		computer_a := pair[0]
		computer_b := pair[1]
		aid, ok := computer_ids[computer_a]
		if !ok {
			computer_ids[computer_a] = i
			aid = i
			computers[aid] = computer_a
			i++
		}
		bid, ok := computer_ids[computer_b]
		if !ok {
			computer_ids[computer_b] = i
			bid = i
			computers[bid] = computer_b
			i++
		}

		conn_a, ok := connections[aid]
		if !ok {
			conn_a = []int{}
		}
		connections[aid] = append(conn_a, bid)

		conn_b, ok := connections[bid]
		if !ok {
			conn_b = []int{}
		}
		connections[bid] = append(conn_b, aid)
	}

	LANs := [][]int{}
	for i := 0; i < len(computers); i++ {
		conn_i, _ := connections[i]
		if len(conn_i) < 2 {
			continue
		}
		for _, j := range conn_i {
			conn_j, _ := connections[j]
			if len(conn_j) < 2 {
				continue
			}

			for _, k := range conn_j {
				if slices.Index(conn_i, k) == -1 {
					continue
				}

				trie := []int{i, j, k}
				slices.Sort(trie)
				if slices.IndexFunc(LANs, func(x []int) bool {
					return x[0] == trie[0] && x[1] == trie[1] && x[2] == trie[2]
				}) != -1 {
					continue
				}
				LANs = append(LANs, trie)
			}
		}
	}

	total := 0
	for _, LAN := range LANs {
		is_good := false
		for _, computer_id := range LAN {
			computer := computers[computer_id]
			if strings.HasPrefix(computer, "t") {
				is_good = true
			}
		}
		if is_good {
			total++
			for i, computer_id := range LAN {
				computer := computers[computer_id]
				fmt.Printf("%s", computer)
				if i < len(LAN)-1 {
					fmt.Printf(",")
				}
			}
			fmt.Printf("\n")
		}
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
