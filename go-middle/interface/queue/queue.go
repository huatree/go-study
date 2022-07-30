package queue

// 定义别名
// 这里的interface{}表示为任意类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
  *q = append(*q, v)
  // 指定类型，执行 q.Push("hello")，即运行时报错
  // *q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
  head := (*q)[0]
  *q = (*q)[1:]
  return head
}

func (q *Queue) IsEmpty() bool {
  return len(*q) == 0
}
