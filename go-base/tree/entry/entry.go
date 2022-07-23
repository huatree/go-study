package main

import "go-study/go-base/tree"

func main() {
  var root tree.Node

  root = tree.Node{Value: 3}
  root.Left = &tree.Node{}
  root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
  root.Right.Left = new(tree.Node)
  root.Left.Right = tree.CreateNode(2)
  root.Right.Left.SetValue(4)

  root.Traverse()
}

// 注意
//
// 项目目录暂时需要放在 `go env GOPATH` 所在的目录下，不然`import "go-study/go-base/tree"`会报错。
// 后面会讲依赖管理，到时候就理解了。

// 封装 => 包
//
// 名字一般使用CamelCase
// 首字母大写： public
// 首字母小写：private

// 包
//
// 每一个目录只能有一个包
// main包包含可执行入口
// 为结构定义的方法必须放在同一个包内
// 可以是不同文件
