package sorting

import "testing"

func Test_less(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"A-Z", args{'A', 'Z'}, true,
		},
		{
			"A-B", args{'A', 'B'}, true,
		},
		{
			"Z-A", args{'Z', 'A'}, false,
		},
		{
			"B-A", args{'B', 'A'}, false,
		},
		{
			"A-A", args{'A', 'A'}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := less(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("less() = %v, want %v", got, tt.want)
			}
		})
	}
}
