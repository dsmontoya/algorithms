package sorting

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		a   []byte
		lo  int
		mid int
		hi  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte("EEGMRACERT"), 0, 4, 9}, []byte("ACEEEGMRRT")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.a, tt.args.lo, tt.args.mid, tt.args.hi)
			if got := tt.args.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("array = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}

func Test_topDownMergesort(t *testing.T) {
	type args struct {
		a  []byte
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte("MERGESORTEXAMPLE"), 0, 15}, []byte("AEEEEGLMMOPRRSTX")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topDownMergesort(tt.args.a, tt.args.lo, tt.args.hi)
			if got := tt.args.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("array = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}

func Test_bottomUpMergesort(t *testing.T) {
	type args struct {
		a  []byte
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte("MERGESORTEXAMPLE"), 0, 15}, []byte("AEEEEGLMMOPRRSTX")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bottomUpMergesort(tt.args.a, tt.args.lo, tt.args.hi)
		})
	}
}
