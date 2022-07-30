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

type Poster interface {
  Post(url string, form map[string]string) string
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
  return r.Get(url)
}

func post(p Poster) {
  p.Post(url, map[string]string{
    "name":  "huatree",
    "study": "golang",
  })
}

type RetrieverPoster interface {
  Retriever
  Poster
  // 当然还可以定义其他方法
}

func session(s RetrieverPoster) string {
  s.Post(url, map[string]string{
    "contents": "another faked baidu.com",
  })
  return s.Get(url)
}

func main() {
  // fmt.Println(download(mock.Retriever{Contents: "this is a fake baidu.com"}))
  var r Retriever

  retriever := mock.Retriever{Contents: "this is a fake baidu.com"}
  r = &retriever
  inspect(r)

  r = &real.Retriever{
    UserAgent: "Mozilla/5.0",
    TimeOut:   time.Minute,
  }
  inspect(r)

  // fmt.Println(download(r))

  // Type assertion
  if mockRetriever, ok := r.(*mock.Retriever); ok {
    fmt.Println(mockRetriever.Contents)
  } else {
    fmt.Println("not a mock retriever")
  }

  fmt.Println("Try a session")
  fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
  fmt.Println("Inspecting", r)
  fmt.Printf("%T %v\n", r, r)
  fmt.Print("Type switch:")
  //? r.(type)
  switch v := r.(type) {
  case *mock.Retriever:
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

## 接口组合

如 io.ReadWriter 组合接口
*/
