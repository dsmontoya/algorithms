package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_PostorderStack(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want []int
	}{
		{"1", testTree, []int{5, 6, 3, 4, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.PostorderStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.PostorderStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
