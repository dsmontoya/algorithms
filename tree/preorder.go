package tree

import "github.com/golang-collections/collections/stack"

func (n *BinaryNode) PreorderStack() (output []int) {
	s := stack.New()
	s.Push(n)
	for {
		if s.Len() == 0 {
			break
		}
		current := s.Pop().(*BinaryNode)
		output = append(output, current.Value)
		if right := current.Right; right != nil {
			s.Push(right)
		}
		if left := current.Left; left != nil {
			s.Push(left)
		}
	}
	return
}
