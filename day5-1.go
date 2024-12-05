package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "slices"
)

func checkRule(nums []int, rule string) (bool) {
  ruleStrs := strings.Split(rule, "|")
  if len(ruleStrs) != 2 {
    return false
  }
  a, _ := strconv.Atoi(ruleStrs[0])
  b, _ := strconv.Atoi(ruleStrs[1])

  aIndex := slices.Index(nums, a)
  if aIndex == -1 {
    return true
  }

  bIndex := slices.Index(nums, b)
  if bIndex == -1 {
    return true
  }

  return aIndex < bIndex
}

func isCorrect(s string, rules []string) (int) {
  numStrs := strings.Split(s, ",")
  middleIndex := (len(numStrs)-1)/2
  middle := 0

  var nums []int
  for i, numS := range numStrs {
    n, _ := strconv.Atoi(numS)
    nums = append(nums, n)
    if i == middleIndex {
      middle = n
    }
  }

  for _, rule := range rules {
    if !checkRule(nums, rule) {
      return 0
    }
  }
  //fmt.Printf("Correct: %s; +%d\n", s, middle)
  return middle
}

func main() {
  data, _ := os.ReadFile("day5.input")
  parts := strings.Split(string(data), "\n\n")
  rules := strings.Split(parts[0], "\n")
  updates := strings.Split(parts[1], "\n")

  // extra whitespace in the end of file
  updates = updates[:len(updates)-1]

  //fmt.Printf("%s", rules[len(rules) - 1])
  //fmt.Printf("Len s: %d, Len s[i]: %d\n", len(lines), len(lines[0]))

  total := 0
  for _, s := range updates {
    //fmt.Printf("Line: %s\n", s)
    total += isCorrect(s, rules)
  }

  fmt.Printf("total: %d\n", total)
}
