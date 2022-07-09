package main

import "fmt"

var a = 3

func compare(a int) int {
  return a - 1
}

func swap(a, b *int) {
  *b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
  return b, a
}

func main() {
  fmt.Println(a, compare(a))

  a, b := 3, 4
  swap(&a, &b)
  fmt.Println(a, b)

  c, d := 3, 4
  fmt.Println(swap2(c, d))
}
