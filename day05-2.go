package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "slices"
)

func filterRules(
  nums []int, rules []string,
) ([][2]int) {
  var applicable [][2]int

  for _, rule := range rules {
    ruleStrs := strings.Split(rule, "|")
    if len(ruleStrs) != 2 {
      continue
    }
    a, _ := strconv.Atoi(ruleStrs[0])
    b, _ := strconv.Atoi(ruleStrs[1])

    aIndex := slices.Index(nums, a)
    if aIndex == -1 {
      continue
    }

    bIndex := slices.Index(nums, b)
    if bIndex == -1 {
      continue
    }
    applicable = append(applicable, [2]int{a, b})
  }
  return applicable
}

func remove(a [][2]int, rule [2]int) [][2]int {
  i := slices.Index(a, rule)
  a[i] = a[len(a)-1]
  return a[:len(a)-1]
}

func order(update string, rules []string) (int) {
  numStrs := strings.Split(update, ",")
  var nums []int
  for _, numS := range numStrs {
    n, _ := strconv.Atoi(numS)
    nums = append(nums, n)
  }

  inDegree := make(map[int]int)
  for _, n := range nums {
    inDegree[n] = 0
  }

  applicable := filterRules(nums, rules)
  for _, rule := range applicable {
    inDegree[rule[1]]++
  }
  //fmt.Printf("%v\n", inDegree)

  var result []int
  var s []int

  for n, in_deg := range inDegree {
    if in_deg == 0 {
      s = append(s, n)
    }
  }

  for len(s) > 0 {
    n := s[len(s) - 1]
    s = s[:len(s)-1]

    result = append(result, n)

    for _, rule := range applicable {
      m := rule[1]
      //fmt.Printf("%v, %d\n", applicable, i)
      if rule[0] == n {
        applicable = remove(applicable, rule)
      }

      canInsert := true
      for _, newRule := range applicable {
        if newRule[1] == m {
          canInsert = false
        }
      }
      if canInsert {
        s = append(s, m)
      }
    }
  }

  middleIndex := (len(result) - 1)/2
  middle := result[middleIndex]

  return middle
}

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
  data, _ := os.ReadFile("day05.input")
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
    if isCorrect(s, rules) == 0 {
      total += order(s, rules)
    }
  }

  fmt.Printf("total: %d\n", total)
}
