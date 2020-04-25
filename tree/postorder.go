package tree

import (
	"github.com/golang-collections/collections/stack"
)

func (n *BinaryNode) PostorderStack() (output []int) {
	s := stack.New()
	s.Push(n)
	for {
		if s.Len() == 0 {
			break
		}
		current := s.Pop().(*BinaryNode)
		output = append([]int{current.Value}, output...)
		if left := current.Left; left != nil {
			s.Push(left)
		}
		if right := current.Right; right != nil {
			s.Push(right)
		}
	}
	return output
}
