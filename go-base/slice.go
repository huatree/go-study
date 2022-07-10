package main

import "fmt"

func updateSlice(s []int) {
  s[0] = 100
}

func main() {
  arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  fmt.Println("arr[2:6] = ", arr[2:6])
  fmt.Println("arr[:6] = ", arr[:6])
  s1 := arr[2:]
  fmt.Println("s1", s1)
  s2 := arr[:]
  fmt.Println("s2", s2)

  fmt.Println("after updateSlice(s1)")
  updateSlice(s1)
  fmt.Println(s1)
  fmt.Println(arr)

  fmt.Println("after updateSlice(s2)")
  updateSlice(s2)
  fmt.Println(s2)
  fmt.Println(arr)

  fmt.Println("reslice")
  fmt.Println(s2)
  s2 = s2[:5]
  fmt.Println(s2)
  s2 = s2[2:]
  fmt.Println(s2)

  fmt.Println("extending slice")
  arr[0], arr[2] = 0, 2
  fmt.Println("arr=", arr)
  s1 = arr[2:6]
  s2 = s1[3:5]
  fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1= [2 3 4 5]
  fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) // s2= [5, 6]
  // fmt.Println(s1[3:7]) // error: slice bounds out of range [:7] with capacity 6
  fmt.Println(s1[3:6]) //  s1= [2 3 4 5] => [5 6 7]
  // slice可以向后扩展，不可以向前扩展
  // s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

  // 向slice添加元素 s = append(s, val)
  s3 := append(s2, 10)
  s4 := append(s3, 11)
  s5 := append(s4, 12)
  fmt.Println("s3, s4, s5=", s3, s4, s5) // s3, s4, s5= [5 6 10] [5 6 10 11] [5 6 10 11 12]
  // s4 and s5 no longer view arr.
  // 添加元素时如果超越cap, 系统会重新分配更大的底层数组
  // 由于值传递的关系，必须接收append的返回值
  fmt.Println("arr=", arr)
}
