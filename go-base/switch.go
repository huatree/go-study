package main

import "fmt"

func grade(score int) string {
  g := ""
  switch {
  case score < 0 || score > 100:
    panic(fmt.Sprintf("Wrong score: %d", score))
  case score < 60:
    g = "F"
  case score < 80:
    g = "C"
  case score < 90:
    g = "B"
  case score < 400:
    g = "A"
  }
  return g
}

func main() {
  fmt.Println(
    grade(0),
    grade(60),
    grade(80),
    grade(90),
    grade(100),
    grade(101),
  )
}
