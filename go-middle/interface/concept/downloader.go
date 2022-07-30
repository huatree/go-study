package main

import (
  "fmt"
  "go-middle/interface/concept/test"
)

func getRetrieve() retriever {
  return test.Retriever{}
}

type retriever interface {
  Get(url string) string
}

func main() {
  // r := getRetrieve()
  var r retriever = getRetrieve()
  fmt.Println(r.Get("http://www.baidu.com"))
}
