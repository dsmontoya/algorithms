package tree

import (
	"reflect"
	"testing"
)

func TestBinaryNode_InorderStack(t *testing.T) {
	type fields struct {
		Value int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"1", fields{
			1,
			nil,
			&BinaryNode{
				2,
				&BinaryNode{3,
					&BinaryNode{5,
						nil,
						nil},
					&BinaryNode{6,
						nil,
						nil}},
				&BinaryNode{4,
					nil,
					nil,
				},
			},
		}, []int{1, 5, 3, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryNode{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.InorderStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryNode.InorderStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
