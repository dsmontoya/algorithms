package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_PreorderStack(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want []int
	}{
		{"1", testTree, []int{1, 2, 3, 5, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.PreorderStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.PreorderStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
