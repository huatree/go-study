package main

import (
  "fmt"
  "go-middle/interface/retriever/mock"
  "go-middle/interface/retriever/real"
  "time"
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
  fmt.Printf("%T %v\n", r, r)
  r = real.Retriever{
    UserAgent: "Mozilla/5.0",
    TimeOut:   time.Minute,
  }
  fmt.Printf("%T %v\n", r, r)
  // fmt.Println(download(r))
}
