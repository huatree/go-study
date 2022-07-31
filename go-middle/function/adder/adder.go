package main

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(v int) int {
    sum += v
    return sum
  }
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
  return func(v int) (int, iAdder) {
    return base + v, adder2(base + v)
  }
}

func main() {
  // a := adder()
  // for i := 0; i < 10; i++ {
  //   fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
  // }

  b := adder2(0)
  for i := 0; i < 10; i++ {
    var s int
    s, b = b(i)
    fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
  }
}

/*
## 函数式编程

- 函数是一等公民：参数，变量，返回值都可以是函数
- 高阶函数
- 函数 -> 闭包

### "正统"函数式编程

- 不可变性：不能有状态，只有常量和函数
- 函数只能有一个参数
*/
