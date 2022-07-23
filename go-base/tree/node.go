package tree

import "fmt"

type Node struct {
  Value       int
  Left, Right *Node
}

/**
 * @description: 自定义工厂函数实现构造函数（Go语言没有构造函数）
 * @param {int} value
 * @return {*}
 */
func CreateNode(value int) *Node {
  return &Node{Value: value} // 返回局部变量的地址
}

/**
 * @description: 获取Node.value值
 * @param {Node} node
 * @return {*}
 */
func (node Node) Print() { // or func Print(node Node) {
  fmt.Print(node.Value, " ")
}

// 使用指针作为方法接收者
// 只有使用指针才可以改变结构内容
// nil指针也可以调用方法
/**
 * @description: 设置Node.value值
 * @param {int} value
 * @return {*}
 */
func (node *Node) SetValue(value int) {
  if node == nil {
    fmt.Println("Setting value to nil" +
      "node.Ingored")
    return
  }
  node.Value = value
}
