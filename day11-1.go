package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "log"
)

func evoluteStone(s string) (res []string) {
  if s == "0" {
    res = append(res, "1")
    return res
  }

  pow := len(s)
  if pow % 2 == 0 {
    a := s[:pow/2]
    res = append(res, a)

    b := strings.TrimLeft(s[pow/2:], "0")
    if len(b) == 0 {
      res = append(res, "0")
    } else {
      res = append(res, b)
    }
    return res
  }

  n, err := strconv.Atoi(s)
  if err != nil {
    log.Fatal(err)
  }
  new_s := fmt.Sprintf("%d", n * 2024)

  new_n, err := strconv.Atoi(new_s)
  if err != nil {
    log.Fatal(err)
  }
  if new_n / 2024 != n {
    log.Fatal("Numbers mismatch")
  }

  res = append(res, new_s)
  return res
}

func evolute(
  stones []string, cache map[string][]string,
) ([]string, map[string][]string) {
  var res []string
  cache_hits := 0
  for _, s := range stones {

    r, new_cache, hit := withCache(s, cache)
    if !hit {
      cache = new_cache
    } else {
      cache_hits++
    }

    res = append(res, r...)
  }
  fmt.Printf(
    "Cache hits: %.1f%%; cache size: %d\n",
    100 * float32(cache_hits) / float32(len(stones)),
    len(cache),
  )

  return res, cache
}

func withCache(
  s string, cache map[string][]string,
) ([]string, map[string][]string, bool) {
  res, ok := cache[s]
  if ok {
    return res, cache, true
  }

  res = evoluteStone(s)
  cache[s] = res

  return res, cache, false
}

func get_key(s string, step int) string {
  return fmt.Sprintf("%s;%d", s, step)
}

func getRes(
  stones []string, n int, cache map[string]int,
) int {
  total := 0
  for _, s := range stones {
    key := get_key(s, n)
    i, ok := cache[key]
    if ok {
      total += i
      continue
    }

    val := 0
    if n == 1 {
      val = len(evoluteStone(s))
      cache[key] = val
    } else {
      val = getRes(evoluteStone(s), n - 1, cache)
      cache[key] = val
    }
    total += val
  }

  return total
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

  data, _ := os.ReadFile("day11.input")
  lines := strings.Split(string(data), "\n")

  stones := strings.Split(lines[0], " ")

  cache := make(map[string]int)
  total := getRes(stones, 25, cache)

  fmt.Printf("total: %d\n", total)
}

// 185452 - too low
