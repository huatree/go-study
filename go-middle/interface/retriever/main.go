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

  r = &mock.Retriever{Contents: "this is a fake baidu.com"}
  inspect(r)

  r = &real.Retriever{
    UserAgent: "Mozilla/5.0",
    TimeOut:   time.Minute,
  }
  inspect(r)

  // fmt.Println(download(r))

  // Type assertion
  if mockRetriever, ok := r.(mock.Retriever); ok {
    fmt.Println(mockRetriever.Contents)
  } else {
    fmt.Println("not a mock retriever")
  }
}

func inspect(r Retriever) {
  fmt.Printf("%T %v\n", r, r)
  fmt.Println("Type switch:")
  //? r.(type)
  switch v := r.(type) {
  case mock.Retriever:
    fmt.Println("Contents:", v.Contents)
  case *real.Retriever:
    fmt.Println("UserAgent:", v.UserAgent)
  }
}

/*
## 接口变量里有什么

- 实现者的类型
- 实现者的指针

接口变量自带指针

接口变量同样采用值传递，几乎不需要使用接口的指针

## 查看接口类型

- 表示任何类型: interface{}
- Type Assertion
- Type Switch
*/
