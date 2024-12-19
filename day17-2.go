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

func disassemble(program []int) {
  ops := []string{
    "adv", "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv",
  }
  fmt.Printf("====ASM====\n")
  for i := 0; i < len(program); i += 2 {
    op := ops[program[i]]

    if op == "adv" ||
      op == "bst" ||
      op == "out" ||
      op == "bdv" ||
      op == "cdv" {
      switch program[i+1] {
        case 4:
          fmt.Printf("%2d: %s A\n", i, op)
        case 5:
          fmt.Printf("%2d: %s B\n", i, op)
        case 6:
          fmt.Printf("%2d: %s C\n", i, op)
        default:
          fmt.Printf("%2d: %s %d\n", i, op, program[i+1])
      }
    } else {
      fmt.Printf("%2d: %s %d\n", i, op, program[i+1])
    }
  }
  fmt.Printf("====---====\n")
}

func mangleFromTo(
  program []int, start, end uint64,
) ([]uint64) {
  valid := []uint64{}
  fmt.Printf(
    "len: %d; %b %b\n",
    len(program),
    start,
    end,
  )


  expected_b := []uint8{}
  for _, p := range program {
    expected_b = append(expected_b, (^(uint8(p))) % 8)
  }
  //fmt.Printf("%v\n", expected_b)

  var i uint64

  brute:
  for i = start; i < end; i++ {
    var A uint64 = i
    //fmt.Printf("%v\n", A)
    for _, b := range expected_b {
      var B uint8 = uint8((A % 8) ^ 2)
      var C uint8 = uint8((A / (1 << B)) % 8)
      if B ^ C != b {
        continue brute
      }
      A = A / 8
    }
    fmt.Printf("res: %b\n", i)
    valid = append(valid, i)
  }
  return valid
}

func mangle(program []int) uint64 {
  //program = program[len(program)-1:]

  valid := []uint64{0}
  for l := 1; l <= len(program); l++ {
    to_run := program[len(program)-l:]
    new_valid := []uint64{}
    for _, v := range valid {
      var start uint64 = uint64(v << 3)
      var end uint64 = uint64((v+1) << 3)
      ans := mangleFromTo(to_run, start, end)
      if len(ans) > 0 {
        new_valid = append(new_valid, ans...)
      }
    }
    valid = new_valid
  }

  var min_v uint64 = 0
  for _, v := range valid {
    if min_v == 0 || min_v > v {
      min_v = v
    }
  }

  return min_v
}

func main() {
	_, program := get_input()
  disassemble(program)
  total := mangle(program)
  fmt.Printf("total: %d\n", total)
}

// 23798076909441 - too low
