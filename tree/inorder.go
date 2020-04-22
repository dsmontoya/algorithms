package tree

import (
	"github.com/golang-collections/collections/stack"
)

func (n *BinaryNode) InorderStack() (result []int) {
	s := stack.New()
	current := n
	for {
		if current == nil && s.Len() == 0 {
			break
		}
		for {
			if current == nil {
				break
			}
			s.Push(current)
			current = current.Left
		}
		current = s.Pop().(*BinaryNode)
		result = append(result, current.Value)
		current = current.Right
	}
	return result
}
