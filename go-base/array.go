package main

import "fmt"

// 数组是值类型
func printArray(arr [5]int) {
  arr[0] = 100
  for i, v := range arr {
    fmt.Println(i, v)
  }
}

func printArrayRef(arr *[5]int) {
  arr[0] = 100
  for i, v := range arr {
    fmt.Println(i, v)
  }
}

func main() {
  var arr1 [5]int
  arr2 := [3]int{1, 3, 5}
  arr3 := [...]int{2, 4, 6, 8, 10}
  var grid [4][5]int
  fmt.Println(arr1, arr2, arr3)
  fmt.Println(grid)

  for i := 0; i < len(arr3); i++ {
    fmt.Println(arr3[i])
  }

  for i := range arr3 {
    fmt.Println(arr3[i])
  }

  for i, v := range arr3 {
    fmt.Println(i, v)
  }

  for _, v := range arr3 { // 可通过下划线省略变量
    fmt.Println(v)
  }
  fmt.Println("printArray(arr1)")
  printArray(arr1)
  fmt.Println("printArray(arr3)")
  printArray(arr3)
  fmt.Println("arr1 and arr3")
  fmt.Println(arr1, arr3)
  fmt.Println("---------")
  fmt.Println("printArrayRef(arr1)")
  printArrayRef(&arr1)
  fmt.Println("printArrayRef(arr3)")
  printArrayRef(&arr3)
  fmt.Println("arr1 and arr3")
  fmt.Println(arr1, arr3)
}
