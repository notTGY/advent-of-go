package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  data, _ := os.ReadFile("day7.test")
  _ := strings.Split(string(data), "\n")


  total := 0

  fmt.Printf("total: %d\n", total)
}
