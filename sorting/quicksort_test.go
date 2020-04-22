package sorting

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
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
		{"1", args{[]byte("QUICKSORTEXAMPLE"), 0, 15}, []byte("ACEEIKLMOPQRSTUX-")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.a, tt.args.lo, tt.args.hi)
			if got := tt.args.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("array = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		a  []byte
		lo int
		hi int
	}
	tests := []struct {
		name  string
		args  args
		wantJ int
		wantA []byte
	}{
		{"1", args{[]byte("KRATELEPUIMQCXOS"), 0, 15}, 5, []byte("ECAIEKLPUTMQRXOS")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.a, tt.args.lo, tt.args.hi); got != tt.wantJ {
				t.Errorf("partition() = %v, wantJ %v", got, tt.wantJ)
			}
			if got := tt.args.a; !reflect.DeepEqual(got, tt.wantA) {
				t.Errorf("array = %v, wantA %v", tt.args.a, tt.wantA)
			}
		})
	}
}
