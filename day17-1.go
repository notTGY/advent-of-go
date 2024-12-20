package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() ([]int, []int) {
	data, _ := os.ReadFile("day17.input")
	blocks := strings.Split(string(data), "\n\n")
	registerStrs := strings.Split(blocks[0], "\n")
	registers := []int{}
	for _, s := range registerStrs {
		splitted := strings.Split(s, ": ")
		x, _ := strconv.Atoi(splitted[1])
		registers = append(registers, x)
	}

	progStr := strings.Split(blocks[1], "\n")[0]
	progList := strings.Split(progStr, ": ")[1]
	program := []int{}
	for _, ps := range strings.Split(progList, ",") {
		x, _ := strconv.Atoi(ps)
		program = append(program, x)
	}
	return registers, program
}

func get_test() ([]int, []int) {
	data, _ := os.ReadFile("day17.test")
	blocks := strings.Split(string(data), "\n\n")
	registerStrs := strings.Split(blocks[0], "\n")
	registers := []int{}
	for _, s := range registerStrs {
		splitted := strings.Split(s, ": ")
		x, _ := strconv.Atoi(splitted[1])
		registers = append(registers, x)
	}

	progStr := strings.Split(blocks[1], "\n")[0]
	progList := strings.Split(progStr, ": ")[1]
	program := []int{}
	for _, ps := range strings.Split(progList, ",") {
		x, _ := strconv.Atoi(ps)
		program = append(program, x)
	}
	return registers, program
}

func adv(registers []int, operand int) int {
	num := registers[0]
	denom := 1 << operand

	return num / denom
}

func bxl(registers []int, operand int) int {
	return registers[1] ^ operand
}

func bst(registers []int, operand int) int {
	return operand % 8
}

func bxc(registers []int) int {
	return registers[1] ^ registers[2]
}

func combo(registers []int, operand int) int {
	if operand < 4 {
		return operand
	}
	return registers[operand-4]
}

func checkQuine(out []int, program []int) bool {
	if len(out) != len(program) {
		return false
	}
	for i, o := range out {
		if program[i] != o {
			return false
		}
	}
	return true
}

func eval(registers []int, program []int) []int {
	out := []int{}

	for ip := 0; ip < len(program); {
		switch program[ip] {
		case 0:
			registers[0] = adv(
				registers,
				combo(registers, program[ip+1]),
			)
		case 1:
			registers[1] = bxl(
				registers,
				program[ip+1],
			)
		case 2:
			registers[1] = bst(
				registers,
				combo(registers, program[ip+1]),
			)
		case 3:
			if registers[0] != 0 {
				ip = program[ip+1]
				continue
			}
		case 4:
			registers[1] = bxc(
				registers,
			)
		case 5:
			out = append(
				out,
				combo(registers, program[ip+1])%8,
			)
		case 6:
			registers[1] = adv(
				registers,
				combo(registers, program[ip+1]),
			)
		case 7:
			registers[2] = adv(
				registers,
				combo(registers, program[ip+1]),
			)
		}

		ip += 2
	}
	return out
}

func main() {
	registers, program := get_input()

	out := eval(registers, program)

	for i, reg := range registers {
		fmt.Printf("reg %d: %d\n", i, reg)
	}

	if len(out) == 0 {
		fmt.Printf("empty")
	}
	for i, o := range out {
		fmt.Printf("%s", strconv.Itoa(o))
		if i < len(out)-1 {
			fmt.Printf(",")
		}
	}

}
