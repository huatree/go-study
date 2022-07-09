package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func converToBin(n int) string {
  reslut := ""
  for ; n > 0; n /= 2 {
    lsb := n % 2
    reslut = strconv.Itoa(lsb) + reslut
  }
  return reslut
}

func readFile(filname string) {
  file, err := os.Open(filname)
  if err != nil {
    panic(err)
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}

func main() {
  // fmt.Println(
  //   converToBin(5),
  //   converToBin(13))

  readFile("abc.txt")
}
