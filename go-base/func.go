package main

import (
  "fmt"
  "math"
  "reflect"
  "runtime"
)

func eval(a, b int, op string) (int, error) {
  switch op {
  case "+":
    return a + b, nil
  case "-":
    return a - b, nil
  case "*":
    return a * b, nil
  case "/":
    q, _ := div(a, b)
    return q, nil
  default:
    return 0, fmt.Errorf("unsupported operation: %s", op)
  }
}

// 返回多个值
func div(a, b int) (q, r int) {
  return a / b, a % b
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
  p := reflect.ValueOf(op).Pointer()
  opName := runtime.FuncForPC(p).Name()
  fmt.Printf("Calling function %s with args"+
    "(%d, %d)\n", opName, a, b)
  return op(a, b)
}

func pow(a, b int) int {
  return int(math.Pow(float64(a), float64(b)))
}

// 可变参数和
// 没有默认参数，可选参数；只有可变参数列表
func sum(values ...int) int {
  s := 0
  for i := range values {
    s += values[i]
  }
  return s
}

func main() {
  if result, err := eval(2, 4, "*"); err != nil {
    fmt.Println("Error: ", err)
  } else {
    fmt.Println(result)
  }

  q, r := div(13, 3)
  fmt.Println(q, r)

  fmt.Println(apply(pow, 3, 4))
  fmt.Println(apply(
    func(a, b int) int {
      return int(math.Pow(float64(a), float64(b)))
    }, 3, 4))

  fmt.Println(sum(1, 3, 4, 6))
}
