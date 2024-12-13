package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "math"
)

func gcd(m, n int) int {
  if m == 0 {
    return n
  }
  if n == 0 {
    return m
  }
  if m == n {
    return m
  }

  if m == 1 && n > 1 {
    return 1
  }
  if n == 1 && m > 1 {
    return 1
  }

  if m % 2 == 0 && n % 2 == 0 {
    return 2*gcd(m/2, n/2)
  }

  if m % 2 == 0 && n % 2 != 0 {
    return gcd(m/2, n)
  }
  if m % 2 != 0 && n % 2 == 0 {
    return gcd(m, n/2)
  }

  if n > m {
    return gcd(m, (n-m)/2)
  }
  return gcd((m-n)/2, n)
}

func lcm(a, b int) int {
  return a * b / gcd(a, b)
}

func findA(x, y, x_req, y_req int) int {
  if x_req < 0 || y_req < 0 {
    return -1
  }
  if x_req % x != 0 || y_req % y != 0 {
    return -1
  }
  Bx := x_req / x
  By := y_req / y
  if Bx == By {
    return Bx
  }
  return -1
}

func solveF(ax, bx, ay, by, tx, ty float64) (
  float64, float64,
) {
  det := ax*by - bx*ay
  if det == 0 {
    return -1, -1
  }

  detX := tx*by - bx*ty
  detY := ax*ty - tx*ay

  return detX / det, detY / det
}

func solve(ax, bx, ay, by, tx, ty int) (int, int) {
  fa, fb := solveF(
    float64(ax),
    float64(bx),
    float64(ay),
    float64(by),
    float64(tx),
    float64(ty),
  )
  if fa < 0 || fb < 0 {
    return -1, -1
  }

  ca := math.Ceil(fa)
  cb := math.Ceil(fb)

  if ca != fa || cb != fb {
    return -1, -1
  }

  ia, _ := strconv.Atoi(fmt.Sprintf("%.0f", ca))
  ib, _ := strconv.Atoi(fmt.Sprintf("%.0f", cb))

  return ia, ib
}

func getMinTokens(x1, y1, x2, y2, x3, y3 int) int {
  delta := 10000000000000
  //delta := 0
  x3 += delta
  y3 += delta

  /*
  fmt.Printf("X: %d; Y: %d\n", x1, y1)
  fmt.Printf("X: %d; Y: %d\n", x2, y2)
  fmt.Printf("X: %d; Y: %d\n", x3, y3)
  */

  tokens := 0

  /*
  x_cycle := lcm(x1, x2)
  y_cycle := lcm(y1, y2)

  a_x_cycle := x_cycle / x1
  a_y_cycle := y_cycle / y1

  a_full_cycle := lcm(a_x_cycle, a_y_cycle)

  b_x_cycle := x_cycle / x2
  b_y_cycle := y_cycle / y2

  b_full_cycle := lcm(b_x_cycle, b_y_cycle)

  remainder_x := x3 % (b_full_cycle * x2)
  remainder_y := y3 % (b_full_cycle * y2)
  B_min := -1
  */

  A_min, B_min := solve(
    x1, x2, y1, y2, x3, y3,
  )

  if B_min == -1 {
    //fmt.Printf("\n")
    return 0
  }
  tokens += A_min * 3 + B_min

  /*
  B_full_cycle := (x3 / (b_full_cycle * x2)) * b_full_cycle
  A_full_cycle := (x3 / (a_full_cycle * x1)) * a_full_cycle * 3
  if B_full_cycle < A_full_cycle {
    tokens += B_full_cycle
  } else {
    tokens += A_full_cycle
  }
  */
  //fmt.Printf("tokens: %d\n\n", tokens)
  return tokens
}

func main() {
  data, _ := os.ReadFile("day13.input")
  blocks := strings.Split(string(data), "\n\n")
  //blocks = blocks[:len(blocks)-1]
  blocks[len(blocks)-1] = blocks[
    len(blocks)-1,
  ][:len(blocks[len(blocks)-1]) - 1]

  total := 0
  for _, block := range blocks {
    lines := strings.Split(block, "\n")

    line1 := lines[0]
    xStr1 := line1[
      strings.Index(line1, "X+") + 2 : strings.Index(line1, ",")]
    x1, _ := strconv.Atoi(string(xStr1))
    yStr1 := line1[
      strings.Index(line1, "Y+") + 2 :]
    y1, _ := strconv.Atoi(string(yStr1))

    line2 := lines[1]
    xStr2 := line2[
      strings.Index(line2, "X+") + 2 : strings.Index(line2, ",")]
    x2, _ := strconv.Atoi(string(xStr2))
    yStr2 := line2[
      strings.Index(line2, "Y+") + 2 :]
    y2, _ := strconv.Atoi(string(yStr2))


    line3 := lines[2]
    xStr3 := line3[
      strings.Index(line3, "X=") + 2 : strings.Index(line3, ",")]
    x3, _ := strconv.Atoi(string(xStr3))
    yStr3 := line3[
      strings.Index(line3, "Y=") + 2 :]
    y3, _ := strconv.Atoi(string(yStr3))

    total += getMinTokens(x1, y1, x2, y2, x3, y3)
  }

  fmt.Printf("total: %d\n", total)
}
