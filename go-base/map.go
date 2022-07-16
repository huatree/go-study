package main

import "fmt"

func main() {
  // 创建
  m := map[string]string{
    "name":    "ccmouse",
    "course":  "golang",
    "site":    "imooc",
    "quality": "notbad",
  }

  // len 获取元素个数
  fmt.Println(len(m))

  m2 := make(map[string]int) // m2 == empty map

  var m3 map[string]int // m3 == nil

  fmt.Println(m, m2, m3)

  fmt.Println("Traversing map:")
  for k, v := range m { // 遍历，无序
    fmt.Println(k, v)
  }

  fmt.Println("Getting values:")
  courseName, ok := m["course"] // 获取元素
  fmt.Println(courseName, ok)
  if courseName, ok := m["coure"]; ok { // key不存在时，获得value类型的初始值
    fmt.Println(courseName, ok)
  } else {
    fmt.Println("this key is not exist")
  }
  fmt.Println("Deleting values:")
  name, ok := m["name"]
  fmt.Println(name, ok)

  delete(m, "name")
  name, ok = m["name"]
  fmt.Println(name, ok)
}

// map的key
// map使用哈希表，必须可以比较相等
// 除了slice，map，function的内建类型都可以作为key
// struct类型不包含上述字段，也可作为key
