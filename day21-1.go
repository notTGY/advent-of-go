package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func get_input() []string {
	data, _ := os.ReadFile("day21.input")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	return lines
}
func get_test() []string {
	data, _ := os.ReadFile("day21.test")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

var keypad = map[rune][2]int{
	'0': [2]int{1, 3}, // 0
	'1': [2]int{0, 2},
	'2': [2]int{1, 2},
	'3': [2]int{2, 2},
	'4': [2]int{0, 1},
	'5': [2]int{1, 1},
	'6': [2]int{2, 1},
	'7': [2]int{0, 0},
	'8': [2]int{1, 0},
	'9': [2]int{2, 0},
	'A': [2]int{2, 3}, // A
}

var directional = map[rune][2]int{
	'A': [2]int{2, 0},
	'>': [2]int{2, 1},
	'<': [2]int{0, 1},
	'v': [2]int{1, 1},
	'^': [2]int{1, 0},
}

func BFS(from, to rune, m map[rune][2]int) []string {
	pos_from, exists := m[from]
	if !exists {
		fmt.Printf("%v\n\"%s\"", m, string(from))
		panic("Wrong map from")
	}
	pos_to, exists := m[to]
	if !exists {
		panic("Wrong map to")
	}
	dx := pos_to[0] - pos_from[0]
	dy := pos_to[1] - pos_from[1]
	if dx == 0 && dy == 0 {
		return []string{""}
	}

	dict := []string{}
	candidates := make(map[rune][2]int)
	if dx > 0 {
		candidates['>'] = [2]int{pos_from[0] + 1, pos_from[1]}
	}
	if dx < 0 {
		candidates['<'] = [2]int{pos_from[0] - 1, pos_from[1]}
	}
	if dy > 0 {
		candidates['v'] = [2]int{pos_from[0], pos_from[1] + 1}
	}
	if dy < 0 {
		candidates['^'] = [2]int{pos_from[0], pos_from[1] - 1}
	}

	for key, cand := range candidates {
		dig := '.'
		for k, v := range m {
			if v[0] == cand[0] && v[1] == cand[1] {
				dig = k
			}
		}
		if dig != '.' {
			res := BFS(dig, to, m)
			for _, s := range res {
				dict = append(dict, string(key)+s)
			}
		}
	}
	return dict
}

func delve(s string, depth int) string {
	if depth == 0 {
		return s
	}

	total := ""
	pos := 'A'
	for _, sym := range s {
		paths := BFS(pos, sym, directional)

		out := delve(paths[0]+"A", depth-1)
		for i, p := range paths {
			if i == 0 {
				continue
			}
			tmp := delve(p+"A", depth-1)
			if len(tmp) < len(out) {
				out = tmp
			}
		}
		total = total + out

		pos = sym
	}

	return total
}

func getShortest(password string) string {
	depth := 2
	pos := 'A'
	total := ""
	for _, n := range password {
		paths := BFS(pos, n, keypad)
		out := delve(paths[0]+"A", depth)
		for i, p := range paths {
			if i == 0 {
				continue
			}
			tmp := delve(p+"A", depth)
			if len(tmp) < len(out) {
				out = tmp
			}
		}
		total = total + out
		pos = n
	}
	return total
}

var answers = []string{
	"<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A",
	"<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A",
	"<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
	"<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A",
	"<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A",
}

func main() {
	start := time.Now()
	passwords := get_input()

	total := 0

	for _, password := range passwords {
		num, found := strings.CutSuffix(password, "A")
		if !found {
			panic("Invalid code")
		}
		count, _ := strconv.Atoi(num)

		out := getShortest(password)
		l := len(out)
		//fmt.Printf("%d: %d * %d\n", i, l, count)
		//fmt.Printf("%s\n", out)
		//fmt.Printf("%s\n", answers[i])

		total += count * l
	}

	fmt.Printf("took: %s; total: %d\n", time.Since(start), total)
}
