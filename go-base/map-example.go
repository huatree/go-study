// 寻找最长不含有重复字符的子串
package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
  lastOccured := make(map[byte]int)
  start := 0
  maxLength := 0
  for i, ch := range []byte(s) {
    if lastI, ok := lastOccured[ch]; ok && lastI >= start {
      start = lastI + 1
    }
    if i-start+1 > maxLength {
      maxLength = i - start + 1
    }
    lastOccured[ch] = i
  }

  return maxLength
}

func main() {
  fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
  fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
  fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
  fmt.Println(lengthOfNonRepeatingSubStr(""))
  fmt.Println(lengthOfNonRepeatingSubStr("b"))
  fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
  // 因为byte，所以该例子不支持中文字符
  // fmt.Println(lengthOfNonRepeatingSubStr("这里是学习战场"))
  // fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}
