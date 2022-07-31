package main

import (
  "fmt"
  "go-middle/function/tree"
)

type myTreeNode struct {
  *tree.Node
}

func (myNode *myTreeNode) postOrder() {
  if myNode == nil || myNode.Node == nil {
    return
  }
  left := myTreeNode{myNode.Left}
  right := myTreeNode{myNode.Right}
  left.postOrder()
  right.postOrder()
  myNode.Print()
}

func main() {
  root := myTreeNode{&tree.Node{Value: 3}}
  root.Left = &tree.Node{}
  root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
  root.Right.Left = new(tree.Node)
  root.Left.Right = tree.CreateNode(2)
  root.Right.Left.SetValue(4)

  root.Traverse()

  nodeCount := 0
  root.TraverseFunc(func(node *tree.Node) {
    nodeCount++
  })
  fmt.Println("Node count:", nodeCount)
}

/*
## 闭包的作用

- 更为自然，不需要修饰如何访问自由变量
- 没有Lambda表达式，但是有匿名函数
*/
