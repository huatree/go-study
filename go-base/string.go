package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  s := "Yes我的世界"
  fmt.Println(s, []byte(s), len(s))
  for _, b := range []byte(s) {
    fmt.Printf("%X ", b)
  }
  fmt.Println()
  for i, ch := range s { // ch is rune, pos对
    fmt.Printf("(%d, %X)", i, ch)
  }
  fmt.Println()
  fmt.Println("Rune count:", utf8.RuneCountInString((s)))
  bytes := []byte(s)
  for len(bytes) > 0 {
    ch, size := utf8.DecodeRune(bytes)
    bytes = bytes[size:]
    fmt.Printf("%c ", ch)
  }
  fmt.Println()

  for i, ch := range []rune(s) { // rune对
    fmt.Printf("(%d, %c) ", i, ch)
  }
  fmt.Println()
}

// 使用range遍历pos, rune对
// 使用utf8.RuneCountInString()获得字符数量
// 使用len获得字节长度
// 使用[]byte获得字节
