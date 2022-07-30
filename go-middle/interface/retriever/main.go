package main

import (
  "fmt"
  "go-middle/interface/retriever/mock"
  "go-middle/interface/retriever/real"
)

type Retriever interface {
  Get(url string) string
}

func download(r Retriever) string {
  return r.Get("http://www.baidu.com")
}

func main() {
  // fmt.Println(download(mock.Retriever{Contents: "this is a fake baidu.com"}))
  var r Retriever
  r = mock.Retriever{Contents: "this is a fake baidu.com"}
  r = real.Retriever{}
  fmt.Println(download(r))
}
