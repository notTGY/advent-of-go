package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func printMap(
	robot [2]int,
	boxes [][2]int,
	edges [][2]int,
) {
	nrows := 0
	ncols := 0
	for _, edge := range edges {
		if edge[0]+2 > ncols {
			ncols = edge[0] + 2
		}
		if edge[1]+1 > nrows {
			nrows = edge[1] + 1
		}
	}

	var m [][]rune
	for row := 0; row < nrows; row++ {
		m = append(m, []rune{})
		for col := 0; col < ncols; col++ {
			m[row] = append(m[row], '.')
		}
	}

	for _, edge := range edges {
		m[edge[1]][edge[0]] = '#'
		m[edge[1]][edge[0]+1] = '#'
	}
	for _, box := range boxes {
		m[box[1]][box[0]] = '['
		m[box[1]][box[0]+1] = ']'
	}
	m[robot[1]][robot[0]] = '@'

	for row := 0; row < nrows; row++ {
		for col := 0; col < ncols; col++ {
			fmt.Printf("%s", string(m[row][col]))
		}
		fmt.Printf("\n")
	}
}

func collider(
	obj [2]int,
	box [2]int,
	obj_size int,
) bool {
	if obj[1] != box[1] {
		return false
	}
	// obj - []
	//       xy

	// box - []
	//       zw

	xz := obj[0] == box[0]
	xw := obj[0] == box[0]+1
	yz := obj[0]+1 == box[0]
	yw := obj[0]+1 == box[0]+1

	if obj_size == 1 {
		return xz || xw
	}

	return xz || xw || yz || yw
}

func canMove(
	obj [2]int,
	boxes [][2]int,
	edges [][2]int,
	inc int,
	axis int,
	obj_size int,
	ignore []int,
) (bool, []int) {
	relevant := []int{}

	new_obj := [2]int{obj[0], obj[1]}
	new_obj[axis] += inc

	for _, edge := range edges {
		collides := collider(
			new_obj,
			edge,
			obj_size,
		)
		if collides {
			return false, relevant
		}
	}

	colliding := []int{}
	for i, box := range boxes {
		if slices.Index(ignore, i) != -1 {
			continue
		}

		collides := collider(
			new_obj,
			box,
			obj_size,
		)
		if collides {
			colliding = append(colliding, i)
			relevant = append(relevant, i)
		}
	}
	if len(colliding) == 0 {
		return true, relevant
	}

	for _, i := range colliding {
		box := boxes[i]

		can, rel := canMove(
			box,
			boxes,
			edges,
			inc,
			axis,
			2,
			append(ignore, i),
		)
		if !can {
			return false, relevant
		}
		for _, i := range rel {
			if slices.Index(relevant, i) == -1 {
				relevant = append(relevant, i)
			}
		}
	}

	return true, relevant
}

func move(
	robot [2]int,
	boxes [][2]int,
	edges [][2]int,
	inc int,
	axis int,
) ([2]int, [][2]int) {
	can, relevant := canMove(
		robot,
		boxes,
		edges,
		inc,
		axis,
		1,
		[]int{},
	)
	if !can {
		return robot, boxes
	}
	//fmt.Printf("relevant: %d\n", len(relevant))

	var new_boxes [][2]int
	for _, box := range boxes {
		new_box := [2]int{box[0], box[1]}
		new_boxes = append(new_boxes, new_box)
	}
	for _, i := range relevant {
		new_boxes[i][axis] += inc
	}
	new_robot := [2]int{robot[0], robot[1]}
	new_robot[axis] += inc
	return new_robot, new_boxes
}

func main() {
	data, _ := os.ReadFile("day15.test")
	blocks := strings.Split(string(data), "\n\n")
	room_map := strings.Split(blocks[0], "\n")
	motions := strings.Split(blocks[1], "\n")
	motions = motions[:len(motions)-1]

	var robot [2]int
	var boxes [][2]int
	var edges [][2]int

	for row, line := range room_map {
		for col, cell := range line {
			coord := [2]int{((col + 1) * 2) - 2, row}
			switch cell {
			case '#':
				edges = append(edges, coord)
			case 'O':
				boxes = append(boxes, coord)
			case '@':
				robot = coord
			}
		}
	}

	DEBUG := true
	if DEBUG {
		fmt.Printf("Initial state:\n")
		printMap(robot, boxes, edges)
		fmt.Printf("\n")
	}

motion_group_loop:
	for _, motion_group := range motions {
		for i, movement := range motion_group {
			if i > 2 {
				break motion_group_loop
			}
			switch movement {
			case '>':
				robot, boxes = move(
					robot, boxes, edges, 1, 0,
				)
			case '<':
				robot, boxes = move(
					robot, boxes, edges, -1, 0,
				)
			case '^':
				robot, boxes = move(
					robot, boxes, edges, -1, 1,
				)
			case 'v':
				robot, boxes = move(
					robot, boxes, edges, 1, 1,
				)
			}
			if DEBUG {
				fmt.Printf("Move %s:\n", string(movement))
				printMap(robot, boxes, edges)
				fmt.Printf("\n")
			}
		}
	}

	total := 0
	for _, box := range boxes {
		total += 100*box[1] + box[0]
	}

	fmt.Printf("total: %d\n", total)
}

// 1451191 - too high
