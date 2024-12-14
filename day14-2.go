package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "time"

  "image"
  "image/color"
  "image/png"

//  "golang.org/x/image/font"
//  "golang.org/x/image/font/basicfont"
//  "golang.org/x/image/math/fixed"
)

func parse(
  lines []string, nrows, ncols int,
) ([][2]int, [][2]int) {
  var ps [][2]int
  var vs [][2]int
  for _, line := range lines {
    var p [2]int
    var v [2]int
    for i, pov_i := range strings.Split(line, " ") {
      coordStrings := strings.Split(
        strings.Split(pov_i, "=")[1],
        ",",
      )
      var coords [2]int
      for j, coordString := range coordStrings {
        coord, _ := strconv.Atoi(coordString)
        coords[j] = coord
      }
      if i == 0 {
        p = coords
      } else {
        v = coords
      }
    }
    if v[0] < 0 {
      v[0] += ncols
    }
    if v[1] < 0 {
      v[1] += nrows
    }
    ps = append(ps, p)
    vs = append(vs, v)
  }
  return ps, vs
}

func get_input() (int, int, int, [][2]int, [][2]int) {
  data, _ := os.ReadFile("day14.input")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
  nrows := 103
  ncols := 101
  nsteps := 10000
  ps, vs := parse(lines, nrows, ncols)
  return nrows, ncols, nsteps, ps, vs
}

func get_test() (int, int, int, [][2]int, [][2]int) {
  data, _ := os.ReadFile("day14.test")
  lines := strings.Split(string(data), "\n")
  lines = lines[:len(lines)-1]
  nrows := 7
  ncols := 11
  nsteps := 100
  ps, vs := parse(lines, nrows, ncols)
  return nrows, ncols, nsteps, ps, vs
}

func getCoords(
  nrows, ncols, nsteps int,
  ps, vs [][2]int,
) [][2]int {
  var coords [][2]int
  for i, p := range ps {
    x := (p[0] + nsteps * vs[i][0]) % ncols
    y := (p[1] + nsteps * vs[i][1]) % nrows

    coords = append(coords, [2]int { x, y })
  }
  return coords
}


func main() {
  start := time.Now()
  nrows, ncols, nsteps, ps, vs := get_input()

  grid_width := 100
  width := (ncols+1)*grid_width
  height := (nrows+1)*(nsteps/grid_width)
  upLeft := image.Point{0, 0}
  lowRight := image.Point{width, height}

  img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
  cyan := color.RGBA{100, 200, 200, 0xff}

  for step := 0; step < nsteps; step++ {
    coords := getCoords(nrows, ncols, step, ps, vs)

    offset_x := (ncols+1)*(step%grid_width)
    offset_y := (nrows+1)*(step/grid_width)

    for col := 0; col < ncols; col++ {
      for row := 0; row < nrows; row++ {
        count := 0
        for _, coord := range coords {
          if coord[0] == col && coord[1] == row {
            count++
          }
        }
        var cl color.Color = color.Black
        if count > 0 {
          cl = color.White
        }
        img.Set(offset_x+col, offset_y+row, cl)
      }
    }

/*
    cl := color.RGBA{90,90,90,0x40}
    point := fixed.Point26_6{
      fixed.I(offset_x), fixed.I(offset_y+9),
    }
    d := &font.Drawer{
      Dst: img,
      Src: image.NewUniform(cl),
      Face: basicfont.Face7x13,
      Dot: point,
    }
    d.DrawString(strconv.Itoa(step))
    */

    for x := 0; x < ncols; x++ {
      img.Set(offset_x+x, offset_y+nrows+1, cyan)
    }
    for y := 0; y < nrows; y++ {
      img.Set(offset_x+ncols+1, offset_y+y, cyan)
    }
  }

  f, _ := os.Create("day14.png")
  png.Encode(f, img)

  fmt.Printf("took %s", time.Since(start))
}
