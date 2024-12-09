package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func checksum(id, start, count int) int {
  n := start + count
  one_sum_p := (start)*(start-1) / 2
  one_sum_c := n*(n-1) / 2
  one_sum := one_sum_c - one_sum_p
  /*
  if start == 0 {
    one_sum = one_sum_c
  }
  */
  //fmt.Printf("%d %d\n", id, count)
  return one_sum * id
}

func main() {
  data, _ := os.ReadFile("day09.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]

  line := lines[0]
  sanity_check_lenght := 0
  var digits []int
  for i, c := range line {
    dig, _ := strconv.Atoi(string(c))
    digits = append(digits, dig)
    if i % 2 == 0 {
      sanity_check_lenght += dig
    }
  }

  last_idx := len(line)-1
  last_id := last_idx/2
  last_count := digits[last_idx]

  pos := 0
  total := 0
  max_pos := 0

  for i := 0; i < len(digits); i++ {
    //fmt.Printf("I: %d\n", i)
    id := i / 2
    count := digits[i]
    is_free_space := i % 2 == 1

    if pos != max_pos {
      fmt.Printf("HUH: %d\n", i)
    }
    max_pos += digits[i]

    if !is_free_space {
      if id < last_id {
        total += checksum(id, pos, count)
        pos += count
        continue
      } else {
        total += checksum(id, pos, last_count)
        pos += last_count
        break
      }
    }
    if is_free_space {
      for count >= last_count && count > 0 {
        total += checksum(last_id, pos, last_count)
        pos += last_count

        count -= last_count
        last_id--
        last_count = digits[2*last_id]
        if id >= last_id {
          max_pos -= count
          break
        }
      }
      if id >= last_id {
        break
      }
      if count > 0 {
        total += checksum(last_id, pos, count)
        pos += count
        last_count -= count
      }
      continue
    }
    fmt.Printf("uncaught %d\n", i)
  }
  //fmt.Printf("%d = %d = %d\n", sanity_check_lenght, pos, max_pos)

  fmt.Printf("total: %d\n", total)
}

// 10180210073447 - too high
// 6311986989618 - too high

// 0620201 <- use this to understand why
// 32211
// 01234
// 02434 = 13
