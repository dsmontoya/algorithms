package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_MaximumDepthTopDown(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want uint
	}{
		{"1", testTree, 3},
		{"2", testTree2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.MaximumDepthTopDown(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.MaximumDepthTopDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNode_MaximumDepthBottomUp(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
		want uint
	}{
		{"1", testTree, 4},
		{"2", testTree2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.root.Value,
				Left:  tt.root.Left,
				Right: tt.root.Right,
			}
			if got := n.MaximumDepthBottomUp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.MaximumDepthBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
