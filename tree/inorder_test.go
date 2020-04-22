package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_InorderStack(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want []int
	}{
		{"1", testTree, []int{1, 5, 3, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.InorderStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.InorderStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
