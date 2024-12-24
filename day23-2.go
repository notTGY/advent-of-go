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

func findMaxClique(
	computers []int,
	connections map[int][]int,
	N int,
) [][]int {
	res := [][]int{}

	if N < 2 {
		return res
	}
	if N == 2 {
		for _, a := range computers {
			conns_a, _ := connections[a]
			conns := []int{}
			for _, c := range conns_a {
				if slices.Index(computers, c) != -1 {
					conns = append(conns, c)
				}
			}
			for _, b := range conns {
				cl := []int{a, b}
				slices.Sort(cl)
				if slices.IndexFunc(res, func(x []int) bool {
					for i, a := range cl {
						if a != x[i] {
							return false
						}
					}
					return true
				}) == -1 {
					res = append(res, cl)
				}
			}
		}
		return res
	}

	for _, a := range computers {
		conns_a, _ := connections[a]
		conns := []int{}
		for _, c := range conns_a {
			if slices.Index(computers, c) != -1 {
				conns = append(conns, c)
			}
		}
		if len(conns) < N-1 {
			continue
		}

		cliques := findMaxClique(conns, connections, N-1)
		for _, clique := range cliques {
			cl := append(clique, a)
			slices.Sort(cl)
			if slices.IndexFunc(res, func(x []int) bool {
				for i, a := range cl {
					if a != x[i] {
						return false
					}
				}
				return true
			}) == -1 {
				res = append(res, cl)
			}
		}
	}
	return res
}

func BronKerbosh(
	R, P, X []int, connections map[int][]int,
) []int {
	if len(P) == 0 && len(X) == 0 {
		return R
	}
	max_clique := []int{}
	for _, v := range P {
		if slices.Index(R, v) != -1 {
			new_P := []int{}
			for _, p := range P {
				if p != v {
					new_P = append(new_P, p)
				}
			}
			P = new_P
			X = append(X, v)
			continue
		}
		Nv, _ := connections[v]
		new_P := []int{}
		new_X := []int{}
		for _, x := range Nv {
			if slices.Index(P, x) != -1 {
				new_P = append(new_P, x)
			}
			if slices.Index(X, x) != -1 {
				new_X = append(new_X, x)
			}
		}
		clique := BronKerbosh(append(R, v), new_P, new_X, connections)
		if len(clique) > len(max_clique) {
			max_clique = clique
		}

		new_P = []int{}
		for _, p := range P {
			if p != v {
				new_P = append(new_P, p)
			}
		}
		P = new_P
		X = append(X, v)
	}
	return max_clique
}

func adjacency(
	all_computers []int, connections map[int][]int,
) []int {
	max_set := []int{}
	max_size := 0

	for i, computer := range all_computers {
		fmt.Printf("%d/%d\n", i, len(all_computers))
		new_set, _ := connections[computer]
		set := []int{}

		for len(new_set) != len(set) {
			set = new_set
			if len(set)+1 <= max_size {
				continue
			}
			scores := make(map[int]int)
			scores[i] = len(set) + 1
			score_nums := []int{len(set) + 1}
			for _, neigh := range set {
				score := 2
				conns, _ := connections[neigh]
				for _, v := range conns {
					for _, s := range set {
						if s == v {
							score++
						}
					}
				}
				scores[neigh] = score
				score_nums = append(score_nums, score)
			}
			slices.Sort(score_nums)
			fmt.Printf("scores: %v\n", scores)
			fmt.Printf("score_nums: %v\n", score_nums)
			if score_nums[0] < len(set) {
				to_remove := -1
				for i, s := range scores {
					if s == score_nums[0] {
						to_remove = i
					}
				}

				new_set = []int{}
				for _, i := range set {
					if i != to_remove {
						new_set = append(new_set, i)
					}
				}
			}
		}

		max_set = append(set, computer)
		max_size = len(max_set)
	}
	return max_set
}

func main() {
	start := time.Now()
	pairs := get_test()

	id := 0
	computer_ids := make(map[string]int)
	computers := []string{}
	all_computers := []int{}
	for _, pair := range pairs {
		for _, computer := range pair {
			_, ok := computer_ids[computer]
			if !ok {
				computer_ids[computer] = id
				computers = append(computers, computer)
				all_computers = append(all_computers, id)
				id++
			}
		}
	}

	slices.Sort(computers)
	for i, computer := range computers {
		computer_ids[computer] = i
	}

	connections := make(map[int][]int)
	for _, pair := range pairs {
		computer_a := pair[0]
		computer_b := pair[1]
		aid, _ := computer_ids[computer_a]
		bid, _ := computer_ids[computer_b]

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
	/*
	  for N := 3; N < 5; N++ {
	    new_LANs := findMaxClique(all_computers, connections, N)
	    all_computers = []int{}
	    for _, LAN := range new_LANs {
	      for _, c := range LAN {
	        if slices.Index(all_computers, c) == -1 {
	          all_computers = append(all_computers, c)
	        }
	      }
	    }
	    //fmt.Printf("Did %d; computers: %d\n", N, len(all_computers))
	    if len(new_LANs) > 0 {
	      LANs = new_LANs
	    } else {
	      break
	    }
	  }
	*/
	/*
	  LAN := BronKerbosh(
	    []int{}, all_computers, []int{}, connections,
	  )
	*/
	//fmt.Printf("LANs=%v\n", LANs)

	/*
	  all_computers = []int{}
	  for computer, _ := range computers {
	    all_computers = append(all_computers, computer)
	  }
	*/

	LAN := adjacency(all_computers, connections)
	fmt.Printf("LAN=%v\n", LAN)
	slices.Sort(LAN)
	LANs = [][]int{LAN}

	total := 0
	for _, LAN := range LANs {
		for i, computer_id := range LAN {
			total++
			computer := computers[computer_id]
			fmt.Printf("%s", computer)
			if i < len(LAN)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}

// as,co,do,kh,km,mc,np,nt,un,uq,wc,wz,yo
