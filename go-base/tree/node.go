package main

import "fmt"

type treeNode struct {
  value       int
  left, right *treeNode
}

/**
 * @description: 自定义工厂函数实现构造函数（Go语言没有构造函数）
 * @param {int} value
 * @return {*}
 */
func createNode(value int) *treeNode {
  return &treeNode{value: value} // 返回局部变量的地址
}

/**
 * @description: 获取treeNode.value值
 * @param {treeNode} node
 * @return {*}
 */
func (node treeNode) print() { // or func print(node treeNode) {
  fmt.Print(node.value, " ")
}

// 使用指针作为方法接收者
// 只有使用指针才可以改变结构内容
// nil指针也可以调用方法
/**
 * @description: 设置treeNode.value值
 * @param {int} value
 * @return {*}
 */
func (node *treeNode) setValue(value int) {
  if node == nil {
    fmt.Println("Setting value to nil" +
      "node.Ingored")
    return
  }
  node.value = value
}

func (node *treeNode) traverse() {
  if node == nil {
    return
  }
  node.left.traverse()
  node.print()
  node.right.traverse()
}

func main() {
  var root treeNode

  // 结构的创建
  // 不论地址还是结构本身，一律使用.来访问成员
  root = treeNode{value: 3}
  root.left = &treeNode{}
  root.right = &treeNode{5, nil, nil}
  root.right.left = new(treeNode) // 0
  root.left.right = createNode(2)

  nodes := []treeNode{
    {value: 3},
    {},
    {6, nil, &root},
  }
  fmt.Println(nodes)

  // --为结构定义方法--
  root.print() // 3
  fmt.Println()

  root.right.left.setValue(4)
  root.right.left.print() // 4
  fmt.Println()

  root.setValue(100)
  root.print() // 100
  fmt.Println()

  pRoot := &root // pRoot是一个地址
  pRoot.print()
  fmt.Println()
  pRoot.setValue(200)
  pRoot.print()
  fmt.Println()

  var cpRoot *treeNode
  fmt.Println("cpRoot", cpRoot) // cpRoot <nil>
  cpRoot.setValue(400)
  cpRoot = &root
  cpRoot.setValue(600)
  cpRoot.print() // 600
  fmt.Println()

  root.traverse() // 0260045, 纵向从左到右一次遍历
}

// 值接收者 vs 指针接收者
//
// 自身
// 要改变内容必须使用指针接收者
// 结构过大也考虑使用指针接收者（考虑性能问题）
// 一致性：如果指针接收者，最好都是指针接收者
//
// 其他
// 值接收者 是go语言特有
// 值/指针接收者均可以接收值/指针
