package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  const filename = "abc.txt"
  // contents, err := ioutil.ReadFile(filename) // 返回2个值（Golang支持）
  // if err != nil {
  //   fmt.Println(err)
  // } else {
  //   fmt.Printf("%s\n", contents)
  // }

  if contents, err := ioutil.ReadFile(filename); err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%s\n", contents)
  }
}
