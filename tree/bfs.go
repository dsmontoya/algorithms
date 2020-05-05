package tree

import (
	"github.com/golang-collections/collections/queue"
)

func (n *BinaryNode) BFS() (output [][]int) {
	q := queue.New()
	if n == nil {
		return
	}
	q.Enqueue(n)
	for {
		l := q.Len()
		a := make([]int, l)
		for i := 0; i < l; i++ {
			current := q.Dequeue().(*BinaryNode)
			a[i] = current.Value
			if left := current.Left; left != nil {
				q.Enqueue(left)
			}
			if right := current.Right; right != nil {
				q.Enqueue(right)
			}
		}
		output = append(output, a)
		if q.Len() == 0 {
			break
		}
	}
	return
}
