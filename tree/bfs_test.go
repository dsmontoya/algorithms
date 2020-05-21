package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_BFS(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want [][]int
	}{
		{"1", testTree, [][]int{{1}, {2}, {3, 4}, {5, 6}}},
		{"2", testTree2, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.BFS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
