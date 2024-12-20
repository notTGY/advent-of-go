package main

import (
	"fmt"
	"os"
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
		if edge[0]+1 > ncols {
			ncols = edge[0] + 1
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
	}
	for _, box := range boxes {
		m[box[1]][box[0]] = 'O'
	}
	m[robot[1]][robot[0]] = '@'

	for row := 0; row < nrows; row++ {
		for col := 0; col < ncols; col++ {
			fmt.Printf("%s", string(m[row][col]))
		}
		fmt.Printf("\n")
	}
}

func move(
	robot [2]int,
	boxes [][2]int,
	edges [][2]int,
	inc int,
	axis int,
) ([2]int, [][2]int) {
	// X=col=[0]
	along := 0
	ortogonal := 1

	// Y=row=[1]
	if axis == 1 {
		along = 1
		ortogonal = 0
	}

	edge_along_val := -1
	for _, edge := range edges {
		along_val := edge[along]
		if edge[ortogonal] == robot[ortogonal] {
			if along_val*inc > inc*robot[along] {
				if edge_along_val == -1 ||
					edge_along_val*inc > inc*along_val {
					edge_along_val = along_val
				}
			}
		}
	}

	//fmt.Printf("edge: %v\n", edges[edgeIdx])

	var relevant []int
	var new_boxes [][2]int
	for i, box := range boxes {
		if box[ortogonal] == robot[ortogonal] &&
			box[along]*inc > inc*robot[along] &&
			box[along]*inc < inc*edge_along_val {
			relevant = append(relevant, i)
		} else {
			new_boxes = append(
				new_boxes,
				[2]int{box[0], box[1]},
			)
		}
	}

	//fmt.Printf("relevant: %d\n", len(relevant))

	spaceCoord := -1
	for coord := robot[along] + inc; coord*inc < inc*edge_along_val; coord += inc {
		//fmt.Printf("coord: %d\n", coord)
		filled := false
		for _, relIdx := range relevant {
			filled = filled || (boxes[relIdx][along] == coord)
		}
		if !filled {
			spaceCoord = coord
			break
		}
	}
	//fmt.Printf("space: %d\n", spaceCoord)

	if spaceCoord != -1 {
		robot[along] += inc
	}
	for _, relIdx := range relevant {
		old_box := boxes[relIdx]
		old_along := old_box[along]
		if spaceCoord != -1 &&
			old_along*inc < inc*spaceCoord {
			old_along += inc
		}

		new_box := [2]int{old_box[0], old_box[1]}
		new_box[along] = old_along

		new_boxes = append(new_boxes, new_box)
	}

	return robot, new_boxes
}

func main() {
	data, _ := os.ReadFile("day15.input")
	blocks := strings.Split(string(data), "\n\n")
	room_map := strings.Split(blocks[0], "\n")
	motions := strings.Split(blocks[1], "\n")
	motions = motions[:len(motions)-1]

	var robot [2]int
	var boxes [][2]int
	var edges [][2]int

	for row, line := range room_map {
		for col, cell := range line {
			coord := [2]int{col, row}
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

	for _, motion_group := range motions {
		for _, movement := range motion_group {
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
			//fmt.Printf("Move %s:\n", string(movement))
			//printMap(robot, boxes, edges)
			//fmt.Printf("\n")
		}
	}

	total := 0
	for _, box := range boxes {
		total += 100*box[1] + box[0]
	}

	fmt.Printf("total: %d\n", total)
}
