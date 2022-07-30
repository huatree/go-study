package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"
)

func printFile(filname string) {
  file, err := os.Open(filname)
  if err != nil {
    panic(err)
  }
  printFileContents(file)
}

func printFileContents(reader io.Reader) {
  scanner := bufio.NewScanner(reader)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}

func main() {
  printFile("interface/read/abc.txt")

  s := `abc"d"
  hhh
  123

  py
  `
  printFileContents(strings.NewReader(s)) // 当然，还有如bytes.NewReader等
}
