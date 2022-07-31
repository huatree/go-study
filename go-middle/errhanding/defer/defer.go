package main

import (
  "bufio"
  "fmt"
  "go-middle/functional/fib"
  "os"
)

func tryDefer() {
  defer fmt.Println(1)
  defer fmt.Println(2)
  fmt.Println(3)
  // panic("error occurred")
  fmt.Println(4)
}

func tryDefer2() {
  for i := 0; i < 100; i++ {
    defer fmt.Println(i)
    if i == 30 {
      panic("printed too many")
    }
  }
}

func writeFile(filename string) {
  file, err := os.Create(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  writer := bufio.NewWriter(file)
  defer writer.Flush()

  f := fib.Fibonacci()
  for i := 0; i < 20; i++ {
    fmt.Fprintln(writer, f())
  }
}

func main() {
  // tryDefer() // 3 2 1
  // writeFile("fib.txt")
  tryDefer2()
}

/*
## defer

defer 栈，先进后出

### 何时使用defer调用

- Open/Close
- Lock/Unlock
- PrintHeader/PrintFooter
*/
