package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

var aa = 3
var ss = "kkk"

// bb := true error
var bb = true

var (
  cc = 3
  dd = true
  ff = "def"
)

func variableZeroValue() {
  var a int
  var s string
  fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
  var a, b int = 3, 4
  var s string = "abc"
  fmt.Println(a, b, s)
}

func variableTypeDeduction() {
  var a, b, c, d = 3, 4, true, "def"
  fmt.Println(a, b, c, d)
}

func variableShorter() {
  a, b, c, s := 3, 4, true, "def"
  b = 5
  fmt.Println(a, b, c, s)
}

func euler() {
  a := 3 + 4i
  fmt.Println(cmplx.Abs(a))
}

func euler2() {
  fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
  fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
  var a, b int = 3, 4
  var c int
  c = int(math.Sqrt(float64(a*a + b*b)))
  fmt.Println(c)
}

func consts() {
  const (
    filename = "abc.txt"
    a, b     = 3, 4
  )
  var c int
  c = int(math.Sqrt(a*a + b*b))
  fmt.Println(filename, c)
}

func enums() {
  const (
    cpp = iota
    _
    python
    golang
    javacsript
  )
  const (
    b = 1 << (10 * iota)
    kb
    mb
    gb
    tb
    pb
  )
  fmt.Println(cpp, javacsript, python, golang)
  fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
  // fmt.Println("Hello, World!")
  // variableZeroValue()
  // variableInitialValue()
  // variableTypeDeduction()
  // variableShorter()
  // euler()
  // euler2()
  // triangle()
  // consts()
  enums()
}
